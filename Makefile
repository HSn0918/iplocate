.PHONY: all build clean test test-unit test-integration run help install install-local uninstall uninstall-local vendor docker-build docker-push docker-buildx buildctl-build

# 项目名称
PROJECT_NAME := IPLocate
# 主程序入口
MAIN_FILE := cmd/iplocate/main.go
# 输出二进制文件
BINARY_NAME := iplocate
# 版本号
VERSION := 1.0.0
# 构建日期
BUILD_DATE := $(shell date +%Y-%m-%d)
# Git 提交哈希
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
# Go 编译标志
GO_FLAGS := -ldflags "-X github.com/hsn0918/iplocate/cmd/iplocate/cmd.Version=$(VERSION) -X github.com/hsn0918/iplocate/cmd/iplocate/cmd.BuildDate=$(BUILD_DATE) -X github.com/hsn0918/iplocate/cmd/iplocate/cmd.GitCommit=$(GIT_COMMIT)"
# 安装目录
INSTALL_DIR := /usr/local/bin
# 用户本地安装目录
LOCAL_INSTALL_DIR := $(HOME)/bin

# 默认目标
all: build

# 编译项目
build:
	@echo "编译项目..."
	@go build $(GO_FLAGS) -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "编译完成: $(BINARY_NAME)"

# 运行项目
run: build
	@echo "运行项目..."
	@./$(BINARY_NAME)

# 清理生成的文件
clean:
	@echo "清理项目..."
	@rm -f $(BINARY_NAME)
	@go clean
	@echo "清理完成"

# 运行单元测试
test-unit:
	@echo "运行单元测试..."
	@go test -v ./pkg/... -count=1
	@echo "单元测试完成"

# 运行集成测试
test-integration:
	@echo "运行集成测试..."
	@go test -v -tags=integration ./pkg/... -count=1
	@echo "集成测试完成"

# 运行所有测试
test: test-unit test-integration

# 测试特定 IP 地址
test-ip:
	@echo "测试 IP 地址: 60.191.18.194"
	@./$(BINARY_NAME) ip -a 60.191.18.194

# 测试特定 IP 地址（显示原始响应）
test-ip-raw:
	@echo "测试 IP 地址（显示原始响应）: 60.191.18.194"
	@./$(BINARY_NAME) ip -a 60.191.18.194 -r

# 测试完整查询
test-full:
	@echo "测试完整查询: 60.191.18.194"
	@./$(BINARY_NAME) full -a 60.191.18.194

# 测试完整查询（显示原始响应）
test-full-raw:
	@echo "测试完整查询（显示原始响应）: 60.191.18.194"
	@./$(BINARY_NAME) full -a 60.191.18.194 -r

# 测试经纬度查询
test-latlng:
	@echo "测试经纬度查询: [30.3095, 120.1536]"
	@./$(BINARY_NAME) latlng -t 30.3095 -g 120.1536

# 测试经纬度查询（显示原始响应）
test-latlng-raw:
	@echo "测试经纬度查询（显示原始响应）: [30.3095, 120.1536]"
	@./$(BINARY_NAME) latlng -t 30.3095 -g 120.1536 -r

# 生成 Bash 自动补全脚本
completion-bash:
	@echo "生成 Bash 自动补全脚本..."
	@./$(BINARY_NAME) completion bash > iplocate_completion.bash
	@echo "生成完成: iplocate_completion.bash"

# 生成 Zsh 自动补全脚本
completion-zsh:
	@echo "生成 Zsh 自动补全脚本..."
	@./$(BINARY_NAME) completion zsh > iplocate_completion.zsh
	@echo "生成完成: iplocate_completion.zsh"

# 安装到系统（需要管理员权限）
install: build
	@echo "安装到系统 ($(INSTALL_DIR))..."
	@echo "注意: 此操作可能需要管理员权限"
	@if [ -d $(INSTALL_DIR) ]; then \
		cp $(BINARY_NAME) $(INSTALL_DIR)/ && echo "✅ 安装成功: $(INSTALL_DIR)/$(BINARY_NAME)" || \
		{ echo "❌ 错误: 安装失败，可能需要管理员权限"; \
		  echo "  请尝试使用: sudo make install"; \
		  echo "  或使用无需管理员权限的命令: make install-local"; exit 1; }; \
	else \
		echo "❌ 错误: 目录 $(INSTALL_DIR) 不存在"; \
		echo "  请尝试使用: make install-local"; \
		exit 1; \
	fi

# 安装到用户本地目录（不需要管理员权限）
install-local: build
	@echo "安装到用户本地目录 ($(LOCAL_INSTALL_DIR))..."
	@mkdir -p $(LOCAL_INSTALL_DIR)
	@cp $(BINARY_NAME) $(LOCAL_INSTALL_DIR)/
	@echo "✅ 安装成功: $(LOCAL_INSTALL_DIR)/$(BINARY_NAME)"
	@if ! echo $$PATH | grep -q "$(LOCAL_INSTALL_DIR)"; then \
		echo "⚠️  警告: $(LOCAL_INSTALL_DIR) 不在您的 PATH 环境变量中"; \
		echo "  请添加以下行到您的 shell 配置文件 (.bashrc, .zshrc 等):"; \
		echo "    export PATH=\$$PATH:$(LOCAL_INSTALL_DIR)"; \
		echo "  然后执行:"; \
		echo "    source ~/.$(shell basename $$SHELL)rc"; \
	else \
		echo "✅ $(LOCAL_INSTALL_DIR) 已在您的 PATH 环境变量中"; \
	fi

# 从系统卸载（需要管理员权限）
uninstall:
	@echo "从系统卸载 ($(INSTALL_DIR)/$(BINARY_NAME))..."
	@echo "注意: 此操作可能需要管理员权限"
	@if [ -f $(INSTALL_DIR)/$(BINARY_NAME) ]; then \
		rm -f $(INSTALL_DIR)/$(BINARY_NAME) && echo "✅ 卸载成功: $(INSTALL_DIR)/$(BINARY_NAME)" || \
		{ echo "❌ 错误: 卸载失败，可能需要管理员权限"; echo "  请尝试使用: sudo make uninstall"; exit 1; }; \
	else \
		echo "⚠️  警告: $(INSTALL_DIR)/$(BINARY_NAME) 不存在，无需卸载"; \
	fi

# 从用户本地目录卸载
uninstall-local:
	@echo "从用户本地目录卸载 ($(LOCAL_INSTALL_DIR)/$(BINARY_NAME))..."
	@if [ -f $(LOCAL_INSTALL_DIR)/$(BINARY_NAME) ]; then \
		rm -f $(LOCAL_INSTALL_DIR)/$(BINARY_NAME) && echo "✅ 卸载成功: $(LOCAL_INSTALL_DIR)/$(BINARY_NAME)"; \
	else \
		echo "⚠️  警告: $(LOCAL_INSTALL_DIR)/$(BINARY_NAME) 不存在，无需卸载"; \
	fi
	@if [ -d $(LOCAL_INSTALL_DIR) ] && [ -z "$(ls -A $(LOCAL_INSTALL_DIR))" ]; then \
		echo "📁 $(LOCAL_INSTALL_DIR) 目录为空，是否删除? [y/N] " && read ans && \
		if [ "$${ans:-N}" = "y" ] || [ "$${ans:-N}" = "Y" ]; then \
			rmdir $(LOCAL_INSTALL_DIR) && echo "✅ 已删除空目录: $(LOCAL_INSTALL_DIR)"; \
		fi; \
	fi

# 创建vendor目录
vendor:
	@echo "创建vendor目录..."
	@go mod vendor
	@echo "✅ vendor目录创建完成"

# 构建Docker镜像
docker-build: vendor
	@echo "构建Docker镜像 $(PROJECT_NAME):$(VERSION)..."
	@docker build -t $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):$(VERSION) .
	@docker tag $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):$(VERSION) $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):latest
	@echo "✅ Docker镜像构建完成: $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):$(VERSION)"

# 推送Docker镜像到仓库
docker-push: docker-build
	@echo "推送Docker镜像到仓库..."
	@echo "请确保已登录到Docker仓库"
	@docker push $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):$(VERSION)
	@docker push $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):latest
	@echo "✅ Docker镜像推送完成"

# 使用Docker Buildx构建多平台镜像
docker-buildx: vendor
	@echo "使用Docker Buildx构建多平台镜像..."
	@if ! docker buildx ls | grep -q multiplatform-builder; then \
		echo "创建多平台构建器..."; \
		docker buildx create --name multiplatform-builder --use; \
	else \
		docker buildx use multiplatform-builder; \
	fi
	@echo "构建并推送多平台镜像 $(PROJECT_NAME):$(VERSION)..."
	@docker buildx build --platform linux/amd64,linux/arm64 \
		-t $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):$(VERSION) \
		-t $(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):latest \
		--push .
	@echo "✅ 多平台镜像构建并推送完成"

# 使用buildctl构建镜像
buildctl-build: vendor
	@echo "使用buildctl构建镜像..."
	@if ! command -v buildctl &> /dev/null; then \
		echo "❌ 错误: buildctl未安装，请先安装buildkit"; \
		echo "  可以通过以下命令安装:"; \
		echo "  brew install buildkit (macOS)"; \
		echo "  或访问 https://github.com/moby/buildkit 获取安装指南"; \
		exit 1; \
	fi
	@buildctl build \
		--frontend dockerfile.v0 \
		--local context=. \
		--local dockerfile=. \
		--output type=image,name=$(shell echo $(PROJECT_NAME) | tr '[:upper:]' '[:lower:]'):$(VERSION),push=true
	@echo "✅ buildctl镜像构建完成"

# 显示帮助信息
help:
	@echo "IPLocate API 位置查询工具 - Makefile 帮助"
	@echo ""
	@echo "可用命令:"
	@echo "  make build          - 编译项目"
	@echo "  make run            - 运行项目"
	@echo "  make clean          - 清理生成的文件"
	@echo "  make test           - 运行所有测试"
	@echo "  make test-unit      - 运行单元测试"
	@echo "  make test-integration - 运行集成测试"
	@echo "  make test-ip        - 测试 IP 地址查询"
	@echo "  make test-ip-raw    - 测试 IP 地址查询（显示原始响应）"
	@echo "  make test-latlng    - 测试经纬度查询"
	@echo "  make test-latlng-raw - 测试经纬度查询（显示原始响应）"
	@echo "  make test-full      - 测试完整查询"
	@echo "  make test-full-raw  - 测试完整查询（显示原始响应）"
	@echo "  make completion-bash - 生成 Bash 自动补全脚本"
	@echo "  make completion-zsh  - 生成 Zsh 自动补全脚本"
	@echo "  make install        - 安装到系统 (需要管理员权限)"
	@echo "  make install-local  - 安装到用户本地目录 (不需要管理员权限)"
	@echo "  make uninstall      - 从系统卸载 (需要管理员权限)"
	@echo "  make uninstall-local - 从用户本地目录卸载"
	@echo "  make vendor         - 创建vendor目录"
	@echo "  make docker-build   - 构建Docker镜像"
	@echo "  make docker-push    - 推送Docker镜像到仓库"
	@echo "  make docker-buildx  - 使用Docker Buildx构建并推送多平台镜像"
	@echo "  make buildctl-build - 使用buildctl构建镜像"
	@echo "  make help           - 显示此帮助信息"
	@echo ""
	@echo "全局选项:"
	@echo "  -d, --debug          启用调试模式"
	@echo "  -l, --log string     日志文件路径 (默认输出到控制台)"
	@echo "  -o, --output-level   输出级别 (0=基本, 1=正常, 2=详细，默认为0)"
	@echo "  -c, --config string  配置文件路径 (默认为 $HOME/.iplocate.yaml)"