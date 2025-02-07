package controller

import (
	"domain-manager/response"
	"domain-manager/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProviderKeyController struct {
	domainService *service.DomainService
}

func NewProviderKeyController(domainService *service.DomainService) *ProviderKeyController {
	return &ProviderKeyController{domainService: domainService}
}

// GetProviderKeys 获取服务商密钥列表
func (c *ProviderKeyController) GetProviderKeys(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	keys, err := c.domainService.GetProviderKeys(ctx, userID)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	// 对密钥进行加密处理
	maskedKeys := make([]map[string]interface{}, len(keys))
	for i, key := range keys {
		maskedKey := make(map[string]interface{})
		// 处理 access_key
		if str := key.AccessKey; str != "" {
			if len(str) > 6 {
				maskedKey["access_key"] = str[:3] + "****" + str[len(str)-3:]
			} else {
				maskedKey["access_key"] = "****"
			}
		}

		// 处理 secret_key
		if str := key.SecretKey; str != "" {
			if len(str) > 6 {
				maskedKey["secret_key"] = str[:3] + "****" + str[len(str)-3:]
			} else {
				maskedKey["secret_key"] = "****"
			}
		}

		// 添加其他字段
		maskedKey["name"] = key.Name
		maskedKey["provider"] = key.Provider
		maskedKey["id"] = key.ID
		maskedKey["created_at"] = key.CreatedAt
		maskedKeys[i] = maskedKey
	}

	response.Success(ctx, maskedKeys)
}

// AddProviderKey 添加服务商密钥
func (c *ProviderKeyController) AddProviderKey(ctx *gin.Context) {
	var req struct {
		Name      string `json:"name" binding:"required"`
		Provider  string `json:"provider" binding:"required"`
		AccessKey string `json:"access_key" binding:"required"`
		SecretKey string `json:"secret_key"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	userID := ctx.GetInt("user_id")
	key, err := c.domainService.AddProviderKey(ctx, userID, req.Name, req.Provider, req.AccessKey, req.SecretKey)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, key)
}

// UpdateProviderKey 更新服务商密钥
func (c *ProviderKeyController) UpdateProviderKey(ctx *gin.Context) {
	var req struct {
		Provider  string `json:"provider" binding:"required"`
		AccessKey string `json:"access_key" binding:"required"`
		SecretKey string `json:"secret_key"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	userID := ctx.GetInt("user_id")
	key, err := c.domainService.UpdateProviderKey(ctx, userID, req.Provider, req.AccessKey, req.SecretKey)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, key)
}

// UpdateProviderKeyName 更新服务商密钥名称
func (c *ProviderKeyController) UpdateProviderKeyName(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	var req struct {
		NewName string `json:"new_name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	// 从 JWT 中获取用户 ID
	userID := ctx.GetInt("user_id")
	key, err := c.domainService.UpdateProviderKeyName(ctx, userID, id, req.NewName)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, key)
}

// DeleteProviderKey 删除服务商密钥
func (c *ProviderKeyController) DeleteProviderKey(ctx *gin.Context) {
	var req struct {
		Provider string `json:"provider" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	userID := ctx.GetInt("user_id")
	if err := c.domainService.DeleteProviderKey(ctx, userID, req.Provider); err != nil {
		response.ErrorMessage(ctx, response.InternalError, err.Error())
		return
	}

	response.Success(ctx, nil)
}
