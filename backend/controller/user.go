package controller

import (
	"domain-manager/config"
	"domain-manager/response"
	"domain-manager/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	user, err := c.userService.CreateUser(ctx, req.Username, req.Password)
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, "创建用户失败")
		return
	}

	response.Success(ctx, gin.H{"id": user.ID, "username": user.Username})
}

// Login 用户登录
func (c *UserController) Login(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	user, err := c.userService.VerifyUser(ctx, req.Username, req.Password)
	if err != nil {
		response.Error(ctx, response.InvalidPassword)
		return
	}

	// 生成JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * time.Duration(config.GlobalConfig.JWT.Expire)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GlobalConfig.JWT.Secret))
	if err != nil {
		response.ErrorMessage(ctx, response.InternalError, "生成token失败")
		return
	}

	response.Success(ctx, gin.H{"token": tokenString})
}

// ChangePassword 修改密码
func (c *UserController) ChangePassword(ctx *gin.Context) {
	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, response.InvalidParams)
		return
	}

	userID := ctx.GetInt("user_id")
	if err := c.userService.ChangePassword(ctx, userID, req.OldPassword, req.NewPassword); err != nil {
		response.ErrorMessage(ctx, response.InternalError, "修改密码失败")
		return
	}

	response.Success(ctx, nil)
}
