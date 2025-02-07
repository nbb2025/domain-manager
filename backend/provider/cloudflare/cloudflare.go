package cloudflare

import (
	"context"
	"domain-manager/provider"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
)

type CloudflareProvider struct {
	api *cloudflare.API
}

// NewCloudflareProvider 创建Cloudflare DNS服务提供商实例
func NewCloudflareProvider(apiToken string) (*CloudflareProvider, error) {
	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		return nil, fmt.Errorf("创建Cloudflare DNS客户端失败: %v", err)
	}
	return &CloudflareProvider{api: api}, nil
}

// GetDomains 获取域名列表
func (p *CloudflareProvider) GetDomains() ([]provider.Domain, error) {
	zones, err := p.api.ListZones(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取域名列表失败: %v", err)
	}

	domains := make([]provider.Domain, 0)
	for _, zone := range zones {
		domains = append(domains, provider.Domain{
			Name:     zone.Name,
			Punycode: zone.Name, // Cloudflare API不直接提供Punycode
		})
	}

	return domains, nil
}

// GetRecords 获取域名解析记录
func (p *CloudflareProvider) GetRecords(domain string) ([]provider.Record, error) {
	zoneID, err := p.api.ZoneIDByName(domain)
	if err != nil {
		return nil, fmt.Errorf("获取域名ID失败: %v", err)
	}

	records, _, err := p.api.ListDNSRecords(context.Background(), &cloudflare.ResourceContainer{Identifier: zoneID}, cloudflare.ListDNSRecordsParams{})
	if err != nil {
		return nil, fmt.Errorf("获取解析记录失败: %v", err)
	}

	result := make([]provider.Record, 0)
	for _, r := range records {
		result = append(result, provider.Record{
			Id:        r.ID,
			Type:      r.Type,
			Name:      r.Name,
			Value:     r.Content,
			Line:      "default", // Cloudflare不支持线路
			TTL:       r.TTL,
			Enabled:   true, // Cloudflare的记录默认都是启用的
			Proxied:   r.Proxied != nil && *r.Proxied,
			UpdatedAt: r.ModifiedOn.Format("2006-01-02 15:04:05"),
		})
	}

	return result, nil
}

// AddRecord 添加解析记录
func (p *CloudflareProvider) AddRecord(domain string, record provider.Record) error {
	zoneID, err := p.api.ZoneIDByName(domain)
	if err != nil {
		return fmt.Errorf("获取域名ID失败: %v", err)
	}

	_, err = p.api.CreateDNSRecord(context.Background(), &cloudflare.ResourceContainer{Identifier: zoneID}, cloudflare.CreateDNSRecordParams{
		Type:    record.Type,
		Name:    record.Name,
		Content: record.Value,
		TTL:     record.TTL,
		Proxied: &record.Proxied,
	})

	if err != nil {
		return fmt.Errorf("添加解析记录失败: %v", err)
	}

	return nil
}

// UpdateRecord 更新解析记录
func (p *CloudflareProvider) UpdateRecord(domain string, record provider.Record) error {
	zoneID, err := p.api.ZoneIDByName(domain)
	if err != nil {
		return fmt.Errorf("获取域名ID失败: %v", err)
	}

	params := cloudflare.UpdateDNSRecordParams{
		Type:    record.Type,
		Name:    record.Name,
		Content: record.Value,
		TTL:     record.TTL,
		Proxied: &record.Proxied,
	}

	_, err = p.api.UpdateDNSRecord(context.Background(), &cloudflare.ResourceContainer{Identifier: zoneID}, params)
	if err != nil {
		return fmt.Errorf("更新解析记录失败: %v", err)
	}

	return nil
}

// DeleteRecord 删除解析记录
func (p *CloudflareProvider) DeleteRecord(domain string, recordId string) error {
	zoneID, err := p.api.ZoneIDByName(domain)
	if err != nil {
		return fmt.Errorf("获取域名ID失败: %v", err)
	}

	err = p.api.DeleteDNSRecord(context.Background(), &cloudflare.ResourceContainer{Identifier: zoneID}, recordId)
	if err != nil {
		return fmt.Errorf("删除解析记录失败: %v", err)
	}

	return nil
}
