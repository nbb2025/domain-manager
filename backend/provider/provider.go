package provider

// Provider 定义了DNS服务提供商需要实现的接口
type Provider interface {
	// GetDomains 获取域名列表
	GetDomains() ([]Domain, error)

	// GetRecords 获取域名解析记录
	GetRecords(domain string) ([]Record, error)

	// AddRecord 添加解析记录
	AddRecord(domain string, record Record) error

	// UpdateRecord 更新解析记录
	UpdateRecord(domain string, record Record) error

	// DeleteRecord 删除解析记录
	DeleteRecord(domain string, recordId string) error
}

// Domain 表示域名信息
type Domain struct {
	Name       string
	Punycode   string
	RecordLine []string
}

// Record 表示解析记录
type Record struct {
	Id        string
	Type      string
	Name      string
	Value     string
	Line      string
	TTL       int
	Enabled   bool // 解析记录是否启用
	Proxied   bool // Cloudflare代理状态（小云朵）
	UpdatedAt string
}
