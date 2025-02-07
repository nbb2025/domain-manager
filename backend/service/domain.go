package service

import (
	"context"
	"domain-manager/ent"
	"domain-manager/ent/domain"
	"domain-manager/ent/providerkey"
	"domain-manager/ent/user"
	"domain-manager/provider"
	"domain-manager/provider/aliyun"
	"domain-manager/provider/cloudflare"
	"domain-manager/provider/tencent"
	"fmt"
	"time"
)

type DomainService struct {
	client *ent.Client
}

func NewDomainService(client *ent.Client) *DomainService {
	return &DomainService{client: client}
}

// GetProvider 根据提供商类型和凭证获取DNS服务提供商实例
func (s *DomainService) GetProvider(ctx context.Context, userId int, providerType string) (provider.Provider, error) {
	// 获取用户的服务提供商凭证
	key, err := s.client.ProviderKey.Query().
		Where(
			providerkey.HasUserWith(user.IDEQ(userId)),
			providerkey.Provider(providerType),
		).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("获取服务提供商凭证失败: %v", err)
	}

	// 根据提供商类型创建相应的Provider实例
	switch providerType {
	case "aliyun":
		return aliyun.NewAliyunProvider(key.AccessKey, key.SecretKey)
	case "cloudflare":
		return cloudflare.NewCloudflareProvider(key.AccessKey)
	case "tencent":
		return tencent.NewTencentProvider(key.AccessKey, key.SecretKey)
	default:
		return nil, fmt.Errorf("不支持的服务提供商类型: %s", providerType)
	}
}

// AddDomain 添加域名
func (s *DomainService) AddDomain(ctx context.Context, userId int, name, providerType string) (*ent.Domain, error) {
	// 验证域名是否已存在
	exists, err := s.client.Domain.Query().
		Where(
			domain.Name(name),
			domain.HasUserWith(user.IDEQ(userId)),
		).Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("检查域名是否存在失败: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("域名已存在")
	}

	// 创建域名记录
	d, err := s.client.Domain.Create().
		SetName(name).
		SetProvider(providerType).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetUserID(userId).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("创建域名失败: %v", err)
	}

	return d, nil
}

// GetDomains 获取用户的域名列表
func (s *DomainService) GetDomains(ctx context.Context, userId int) ([]*ent.Domain, error) {
	domains, err := s.client.Domain.Query().
		Where(domain.HasUserWith(user.IDEQ(userId))).
		Order(ent.Desc(domain.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("获取域名列表失败: %v", err)
	}

	return domains, nil
}

// GetAvailableDomains 获取可用域名列表
func (s *DomainService) GetAvailableDomains(ctx context.Context, userId int, providerType string) ([]string, error) {
	// 获取服务提供商实例
	p, err := s.GetProvider(ctx, userId, providerType)
	if err != nil {
		return nil, fmt.Errorf("获取服务提供商实例失败: %v", err)
	}

	// 获取域名列表
	domains, err := p.GetDomains()
	if err != nil {
		return nil, fmt.Errorf("获取域名列表失败: %v", err)
	}

	// 只返回域名字符串
	domainNames := make([]string, 0, len(domains))
	for _, d := range domains {
		domainNames = append(domainNames, d.Name)
	}

	return domainNames, nil
}

// GetDomain 获取单个域名信息
func (s *DomainService) GetDomain(ctx context.Context, userId int, domainId int) (*ent.Domain, error) {
	d, err := s.client.Domain.Query().
		Where(
			domain.IDEQ(domainId),
			domain.HasUserWith(user.IDEQ(userId)),
		).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("获取域名信息失败: %v", err)
	}

	return d, nil
}

// GetDomainRecords 获取域名的解析记录
func (s *DomainService) GetDomainRecords(ctx context.Context, userId int, domainId int) ([]provider.Record, error) {
	// 获取域名信息
	d, err := s.client.Domain.Query().
		Where(
			domain.ID(domainId),
			domain.HasUserWith(user.IDEQ(userId)),
		).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("获取域名信息失败: %v", err)
	}

	// 获取DNS服务提供商实例
	p, err := s.GetProvider(ctx, userId, d.Provider)
	if err != nil {
		return nil, err
	}

	// 获取解析记录
	records, err := p.GetRecords(d.Name)
	if err != nil {
		return nil, fmt.Errorf("获取解析记录失败: %v", err)
	}

	return records, nil
}

// AddDomainRecord 添加域名解析记录
func (s *DomainService) AddDomainRecord(ctx context.Context, userId int, domainId int, record provider.Record) error {
	// 获取域名信息
	d, err := s.client.Domain.Query().
		Where(
			domain.ID(domainId),
			domain.HasUserWith(user.IDEQ(userId)),
		).Only(ctx)

	if err != nil {
		return fmt.Errorf("获取域名信息失败: %v", err)
	}

	// 获取DNS服务提供商实例
	p, err := s.GetProvider(ctx, userId, d.Provider)
	if err != nil {
		return err
	}

	// 添加解析记录
	if err := p.AddRecord(d.Name, record); err != nil {
		return fmt.Errorf("添加解析记录失败: %v", err)
	}

	return nil
}

// UpdateDomainRecord 更新域名解析记录
func (s *DomainService) UpdateDomainRecord(ctx context.Context, userId int, domainId int, record provider.Record) error {
	// 获取域名信息
	d, err := s.client.Domain.Query().
		Where(
			domain.ID(domainId),
			domain.HasUserWith(user.IDEQ(userId)),
		).Only(ctx)

	if err != nil {
		return fmt.Errorf("获取域名信息失败: %v", err)
	}

	// 获取DNS服务提供商实例
	p, err := s.GetProvider(ctx, userId, d.Provider)
	if err != nil {
		return err
	}

	// 直接调用DNS服务商API更新解析记录
	if err := p.UpdateRecord(d.Name, record); err != nil {
		return fmt.Errorf("更新解析记录失败: %v", err)
	}

	return nil
}

// DeleteDomainRecord 删除域名解析记录
func (s *DomainService) DeleteDomainRecord(ctx context.Context, userId int, domainId int, recordId string) error {
	// 获取域名信息
	d, err := s.client.Domain.Query().
		Where(
			domain.ID(domainId),
			domain.HasUserWith(user.IDEQ(userId)),
		).Only(ctx)

	if err != nil {
		return fmt.Errorf("获取域名信息失败: %v", err)
	}

	// 获取DNS服务提供商实例
	p, err := s.GetProvider(ctx, userId, d.Provider)
	if err != nil {
		return err
	}

	// 删除解析记录
	if err := p.DeleteRecord(d.Name, recordId); err != nil {
		return fmt.Errorf("删除解析记录失败: %v", err)
	}

	return nil
}

// GetProviderKeys 获取用户的服务商密钥列表
func (s *DomainService) GetProviderKeys(ctx context.Context, userId int) ([]*ent.ProviderKey, error) {
	keys, err := s.client.ProviderKey.Query().
		Where(providerkey.HasUserWith(user.IDEQ(userId))).
		Order(ent.Desc(providerkey.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("获取服务商密钥列表失败: %v", err)
	}

	return keys, nil
}

// AddProviderKey 添加服务商密钥
func (s *DomainService) AddProviderKey(ctx context.Context, userId int, name, provider, accessKey, secretKey string) (*ent.ProviderKey, error) {
	// 验证是否已存在该服务商的密钥
	exists, err := s.client.ProviderKey.Query().
		Where(
			providerkey.HasUserWith(user.IDEQ(userId)),
			providerkey.Provider(provider),
		).Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("检查服务商密钥是否存在失败: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("该服务商的密钥已存在")
	}

	// 创建服务商密钥
	key, err := s.client.ProviderKey.Create().
		SetName(name).
		SetProvider(provider).
		SetAccessKey(accessKey).
		SetSecretKey(secretKey).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetUserID(userId).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("创建服务商密钥失败: %v", err)
	}

	return key, nil
}

// UpdateProviderKey 更新服务商密钥
func (s *DomainService) UpdateProviderKey(ctx context.Context, userId int, provider, accessKey, secretKey string) (*ent.ProviderKey, error) {
	// 获取现有密钥
	key, err := s.client.ProviderKey.Query().
		Where(
			providerkey.HasUserWith(user.IDEQ(userId)),
			providerkey.Provider(provider),
		).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("获取服务商密钥失败: %v", err)
	}

	// 更新密钥信息
	key, err = key.Update().
		SetAccessKey(accessKey).
		SetSecretKey(secretKey).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("更新服务商密钥失败: %v", err)
	}

	return key, nil
}

// DeleteProviderKey 删除服务商密钥
func (s *DomainService) DeleteProviderKey(ctx context.Context, userId int, provider string) error {
	// 获取要删除的密钥
	key, err := s.client.ProviderKey.Query().
		Where(
			providerkey.HasUserWith(user.IDEQ(userId)),
			providerkey.Provider(provider),
		).Only(ctx)

	if err != nil {
		return fmt.Errorf("获取服务商密钥失败: %v", err)
	}

	// 删除密钥
	if err := s.client.ProviderKey.DeleteOne(key).Exec(ctx); err != nil {
		return fmt.Errorf("删除服务商密钥失败: %v", err)
	}

	return nil
}

// UpdateProviderKeyName 更新服务商密钥名称
func (s *DomainService) UpdateProviderKeyName(ctx context.Context, userId int, id int, newName string) (*ent.ProviderKey, error) {
	// 验证密钥是否属于当前用户
	key, err := s.client.ProviderKey.Query().
		Where(
			providerkey.HasUserWith(user.IDEQ(userId)),
			providerkey.ID(id),
		).Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("密钥不存在或无权访问")
		}
		return nil, fmt.Errorf("获取服务商密钥失败: %v", err)
	}

	// 验证新名称是否已存在
	exists, err := s.client.ProviderKey.Query().
		Where(
			providerkey.HasUserWith(user.IDEQ(userId)),
			providerkey.Name(newName),
			providerkey.IDNEQ(id),
		).Exist(ctx)

	if err != nil {
		return nil, fmt.Errorf("检查密钥名称是否存在失败: %v", err)
	}

	if exists {
		return nil, fmt.Errorf("密钥名称已存在")
	}

	// 更新密钥名称
	key, err = key.Update().
		SetName(newName).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("更新密钥名称失败: %v", err)
	}

	return key, nil
}
