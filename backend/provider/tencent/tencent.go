package tencent

import (
	"domain-manager/provider"
	"fmt"
	"strconv"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

type TencentProvider struct {
	client *dnspod.Client
}

// NewTencentProvider 创建腾讯云DNS服务提供商实例
func NewTencentProvider(secretId, secretKey string) (*TencentProvider, error) {
	credential := common.NewCredential(secretId, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"

	client, err := dnspod.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		return nil, fmt.Errorf("创建腾讯云DNS客户端失败: %v", err)
	}
	return &TencentProvider{client: client}, nil
}

// GetDomains 获取域名列表
func (p *TencentProvider) GetDomains() ([]provider.Domain, error) {
	request := dnspod.NewDescribeDomainListRequest()

	response, err := p.client.DescribeDomainList(request)
	if err != nil {
		return nil, fmt.Errorf("获取域名列表失败: %v", err)
	}

	domains := make([]provider.Domain, 0)
	for _, d := range response.Response.DomainList {
		domains = append(domains, provider.Domain{
			Name:       *d.Name,
			Punycode:   *d.Punycode,
			RecordLine: []string{}, // 腾讯云不支持线路设置，返回空数组
		})
	}

	return domains, nil
}

// GetRecords 获取域名解析记录
func (p *TencentProvider) GetRecords(domain string) ([]provider.Record, error) {
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = &domain

	response, err := p.client.DescribeRecordList(request)
	if err != nil {
		return nil, fmt.Errorf("获取解析记录失败: %v", err)
	}

	records := make([]provider.Record, 0)
	for _, r := range response.Response.RecordList {
		records = append(records, provider.Record{
			Id:        strconv.FormatUint(*r.RecordId, 10),
			Type:      *r.Type,
			Name:      *r.Name,
			Value:     *r.Value,
			Line:      *r.Line,
			TTL:       int(*r.TTL),
			Enabled:   *r.Status == "ENABLE",
			UpdatedAt: *r.UpdatedOn,
		})
	}

	return records, nil
}

// AddRecord 添加解析记录
func (p *TencentProvider) AddRecord(domain string, record provider.Record) error {
	request := dnspod.NewCreateRecordRequest()
	request.Domain = &domain
	request.SubDomain = &record.Name
	request.RecordType = &record.Type
	request.RecordLine = &record.Line
	request.Value = &record.Value
	ttl := uint64(record.TTL)
	request.TTL = &ttl

	_, err := p.client.CreateRecord(request)
	if err != nil {
		return fmt.Errorf("添加解析记录失败: %v", err)
	}

	return nil
}

// UpdateRecord 更新解析记录
func (p *TencentProvider) UpdateRecord(domain string, record provider.Record) error {
	request := dnspod.NewModifyRecordRequest()
	request.Domain = &domain
	recordId, _ := strconv.ParseUint(record.Id, 10, 64)
	request.RecordId = &recordId
	request.SubDomain = &record.Name
	request.RecordType = &record.Type
	request.RecordLine = &record.Line
	request.Value = &record.Value
	ttl := uint64(record.TTL)
	request.TTL = &ttl

	_, err := p.client.ModifyRecord(request)
	if err != nil {
		return fmt.Errorf("更新解析记录失败: %v", err)
	}

	return nil
}

// DeleteRecord 删除解析记录
func (p *TencentProvider) DeleteRecord(domain string, recordId string) error {
	request := dnspod.NewDeleteRecordRequest()
	request.Domain = &domain
	id, _ := strconv.ParseUint(recordId, 10, 64)
	request.RecordId = &id

	_, err := p.client.DeleteRecord(request)
	if err != nil {
		return fmt.Errorf("删除解析记录失败: %v", err)
	}

	return nil
}
