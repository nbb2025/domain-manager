# Domain Manager

一个基于 Vue3 和 Go 的 DNS 域名管理系统，支持多个 DNS 服务商的域名管理和解析记录维护。

## 项目说明

本项目是一个跨平台的 DNS 域名管理系统，旨在提供一个统一的界面来管理多个 DNS 服务商的域名。目前支持以下 DNS 服务商：

- 阿里云 DNS
- 腾讯云 DNSPod
- Cloudflare

## 主要功能

- 多 DNS 服务商支持
- 统一的域名管理界面
- DNS 解析记录的增删改查
- 支持多用户管理
- 服务商密钥管理

## 技术栈

### 前端

- Vue 3
- Vue Router
- Tailwind CSS
- Vite

### 后端

- Go
- Gin Web Framework
- Ent ORM
- SQLite

## 开发状态

⚠️ **重要提示：本项目目前处于开发阶段，存在以下已知问题：**

- 部分功能尚未完全实现
- 可能存在未知的 bug
- 缺少完整的测试覆盖
- 安全性尚未经过全面审查
- 性能优化有待改进

**因此，强烈建议不要在生产环境中使用本项目。**

## 待实现功能

- [ ] 批量导入导出功能
- [ ] 域名分组管理
- [ ] 操作日志记录
- [ ] 更多 DNS 服务商支持
- [ ] 完整的单元测试
- [ ] 性能优化
- [ ] 安全性增强

## 安装和使用

### 前端

```bash
cd frontend
npm install
npm run dev
```

### 后端

```bash
cd backend
go mod tidy
go run main.go
```

## 贡献

由于项目仍在开发中，如果您发现任何问题或有改进建议，欢迎提出 Issue 或提交 Pull Request。

## 许可证

MIT License