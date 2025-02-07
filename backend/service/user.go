package service

import (
	"context"
	"domain-manager/ent"
	"domain-manager/ent/user"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{client: client}
}

// CreateUser 创建新用户
func (s *UserService) CreateUser(ctx context.Context, username, password string) (*ent.User, error) {
	// 检查用户名是否已存在
	exists, err := s.client.User.Query().Where(user.UsernameEQ(username)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("用户名已被注册")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	now := time.Now()
	return s.client.User.Create().
		SetUsername(username).
		SetPassword(string(hashedPassword)).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
}

// VerifyUser 验证用户登录
func (s *UserService) VerifyUser(ctx context.Context, username, password string) (*ent.User, error) {
	// 查找用户
	user, err := s.client.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

// ChangePassword 修改用户密码
func (s *UserService) ChangePassword(ctx context.Context, userID int, oldPassword, newPassword string) error {
	// 查找用户
	user, err := s.client.User.Get(ctx, userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	return s.client.User.UpdateOne(user).
		SetPassword(string(hashedPassword)).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
}
