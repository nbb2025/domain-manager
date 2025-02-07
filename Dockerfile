# 前端构建阶段
FROM node:18-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install && npm install terser
COPY frontend/ ./
RUN npm run build

# 后端构建阶段
FROM golang:1.20-alpine AS backend-builder
WORKDIR /app/backend
COPY backend/go.* ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o domain-manager .

# 最终运行阶段
FROM alpine:3.18
WORKDIR /app

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 创建数据目录
RUN mkdir -p /app/data

# 复制构建产物
COPY --from=frontend-builder /app/frontend/dist /app/dist
COPY --from=backend-builder /app/backend/main/domain-manager /app
COPY --from=backend-builder /app/backend/config /app/config

# 设置环境变量
ENV TZ=Asia/Shanghai

# 声明数据卷
VOLUME ["/app/data"]

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["/app/domain-manager"]