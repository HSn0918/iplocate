# 第一阶段：构建环境
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的工具
RUN apk add --no-cache git make

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 复制vendor目录
COPY vendor/ ./vendor/

# 复制源代码
COPY . .

# 使用Makefile构建应用（使用vendor）
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o iplocate cmd/iplocate/main.go

# 第二阶段：运行环境
FROM alpine:latest

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 设置时区为亚洲/上海
ENV TZ=Asia/Shanghai

# 创建非root用户
RUN adduser -D -h /app appuser

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/iplocate /app/

# 切换到非root用户
USER appuser

# 设置入口点
ENTRYPOINT ["/app/iplocate"]

# 默认命令
CMD ["--help"]