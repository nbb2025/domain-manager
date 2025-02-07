package controller

import (
	"domain-manager/provider"
	"domain-manager/response"
	"domain-manager/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DomainController struct {
	domainService *service.DomainService
}

func NewDomainController(domainService *service.DomainService) *DomainController {
	return &DomainController{domainService: domainService}
}

// GetDomains 获取域名列表
func (c *DomainController) GetDomains(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	domains, err := c.domainService.GetDomains(ctx, userID)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, "获取域名列表失败")
		return
	}

	response.Success(ctx, domains)
}

// AddDomain 添加域名
func (c *DomainController) AddDomain(ctx *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Provider string `json:"provider" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	userID := ctx.GetInt("user_id")
	domain, err := c.domainService.AddDomain(ctx, userID, req.Name, req.Provider)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, domain)
}

// GetDomain 获取单个域名信息
func (c *DomainController) GetDomain(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	domainID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	domain, err := c.domainService.GetDomain(ctx, userID, domainID)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, domain)
}

// GetAvailableDomains 获取可用域名列表
func (c *DomainController) GetAvailableDomains(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	provider := ctx.Query("provider")
	if provider == "" {
		response.Error(ctx, response.InvalidParams)
		return
	}

	domains, err := c.domainService.GetAvailableDomains(ctx, userID, provider)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, domains)
}

// GetDomainRecords 获取域名解析记录
func (c *DomainController) GetDomainRecords(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	domainID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	records, err := c.domainService.GetDomainRecords(ctx, userID, domainID)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, records)
}

// AddDomainRecord 添加域名解析记录
func (c *DomainController) AddDomainRecord(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	domainID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	var req struct {
		Type     string `json:"type" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Value    string `json:"value" binding:"required"`
		Priority int    `json:"priority"`
		TTL      int    `json:"ttl"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	record := provider.Record{
		Type:  req.Type,
		Name:  req.Name,
		Value: req.Value,
		TTL:   req.TTL,
	}

	if err := c.domainService.AddDomainRecord(ctx, userID, domainID, record); err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// UpdateDomainRecord 更新域名解析记录
func (c *DomainController) UpdateDomainRecord(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	domainID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	var req struct {
		ID       string `json:"id" binding:"required"`
		Type     string `json:"type" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Value    string `json:"value" binding:"required"`
		Priority int    `json:"priority"`
		TTL      int    `json:"ttl"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	record := provider.Record{
		Id:    req.ID,
		Type:  req.Type,
		Name:  req.Name,
		Value: req.Value,
		TTL:   req.TTL,
	}

	if err := c.domainService.UpdateDomainRecord(ctx, userID, domainID, record); err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, nil)
}

// DeleteDomainRecord 删除域名解析记录
func (c *DomainController) DeleteDomainRecord(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	domainID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	recordID := ctx.Param("record_id")
	if recordID == "" {
		response.Error(ctx, response.InvalidParams)
		return
	}

	if err := c.domainService.DeleteDomainRecord(ctx, userID, domainID, recordID); err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, nil)
}
