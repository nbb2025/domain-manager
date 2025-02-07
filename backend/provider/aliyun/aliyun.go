package aliyun

import (
	"domain-manager/provider"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type AliyunProvider struct {
	client *alidns.Client
}

// NewAliyunProvider 创建阿里云DNS服务提供商实例
func NewAliyunProvider(accessKey, secretKey string) (*AliyunProvider, error) {
	client, err := alidns.NewClientWithAccessKey(
		"cn-hangzhou",
		accessKey,
		secretKey,
	)
	if err != nil {
		return nil, fmt.Errorf("创建阿里云DNS客户端失败: %v", err)
	}
	return &AliyunProvider{client: client}, nil
}

// GetDomains 获取域名列表
func (p *AliyunProvider) GetDomains() ([]provider.Domain, error) {
	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"

	response, err := p.client.DescribeDomains(request)
	if err != nil {
		return nil, fmt.Errorf("获取域名列表失败: %v", err)
	}

	domains := make([]provider.Domain, 0)
	for _, d := range response.Domains.Domain {
		domains = append(domains, provider.Domain{
			Name:       d.DomainName,
			Punycode:   d.PunyCode,
			RecordLine: []string{}, // 阿里云不支持线路设置，返回空数组
		})
	}

	return domains, nil
}

// GetRecords 获取域名解析记录
func (p *AliyunProvider) GetRecords(domain string) ([]provider.Record, error) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.DomainName = domain

	response, err := p.client.DescribeDomainRecords(request)
	if err != nil {
		return nil, fmt.Errorf("获取解析记录失败: %v", err)
	}

	records := make([]provider.Record, 0)
	for _, r := range response.DomainRecords.Record {
		records = append(records, provider.Record{
			Id:        r.RecordId,
			Type:      r.Type,
			Name:      r.RR,
			Value:     r.Value,
			Line:      r.Line,
			TTL:       int(r.TTL),
			Enabled:   r.Status == "ENABLE",
			UpdatedAt: "",
		})
	}

	return records, nil
}

// AddRecord 添加解析记录
func (p *AliyunProvider) AddRecord(domain string, record provider.Record) error {
	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"
	request.DomainName = domain
	request.RR = record.Name
	request.Type = record.Type
	request.Value = record.Value
	request.Line = record.Line
	request.TTL = requests.NewInteger(record.TTL)

	_, err := p.client.AddDomainRecord(request)
	if err != nil {
		return fmt.Errorf("添加解析记录失败: %v", err)
	}

	return nil
}

// UpdateRecord 更新解析记录
func (p *AliyunProvider) UpdateRecord(domain string, record provider.Record) error {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"
	request.RecordId = record.Id
	request.RR = record.Name
	request.Type = record.Type
	request.Value = record.Value
	request.Line = record.Line
	request.TTL = requests.NewInteger(record.TTL)

	_, err := p.client.UpdateDomainRecord(request)
	if err != nil {
		return fmt.Errorf("更新解析记录失败: %v", err)
	}

	return nil
}

// DeleteRecord 删除解析记录
func (p *AliyunProvider) DeleteRecord(domain string, recordId string) error {
	request := alidns.CreateDeleteDomainRecordRequest()
	request.Scheme = "https"
	request.RecordId = recordId

	_, err := p.client.DeleteDomainRecord(request)
	if err != nil {
		return fmt.Errorf("删除解析记录失败: %v", err)
	}

	return nil
}
