#!/bin/bash

# 编译脚本，用于生成不同平台的二进制文件

echo "开始编译不同平台的 iplocate 二进制文件..."

# 确保 bin 目录存在
mkdir -p bin

# 清理旧的二进制文件
rm -f bin/iplocate-*
rm -f bin/iplocate.exe

# 编译 MacOS (Intel & Apple Silicon)
echo "编译 macOS 版本..."
GOOS=darwin GOARCH=arm64 go build -o bin/iplocate-darwin-arm64 ../cmd/iplocate/main.go
GOOS=darwin GOARCH=amd64 go build -o bin/iplocate-darwin-amd64 ../cmd/iplocate/main.go

# 如果在 macOS ARM64 上，将 ARM64 版本复制为默认的 macOS 版本
if [[ "$(uname)" == "Darwin" && "$(uname -m)" == "arm64" ]]; then
  cp bin/iplocate-darwin-arm64 bin/iplocate-darwin
elif [[ "$(uname)" == "Darwin" && "$(uname -m)" == "x86_64" ]]; then
  cp bin/iplocate-darwin-amd64 bin/iplocate-darwin
else
  # 如果不在 macOS 上，默认使用 ARM64 作为 macOS 版本
  cp bin/iplocate-darwin-arm64 bin/iplocate-darwin
fi

# 编译 Linux (x64 & ARM64)
echo "编译 Linux 版本..."
GOOS=linux GOARCH=amd64 go build -o bin/iplocate-linux ../cmd/iplocate/main.go
GOOS=linux GOARCH=arm64 go build -o bin/iplocate-linux-arm64 ../cmd/iplocate/main.go

# 编译 Windows
echo "编译 Windows 版本..."
GOOS=windows GOARCH=amd64 go build -o bin/iplocate.exe ../cmd/iplocate/main.go

# 设置权限
chmod +x bin/iplocate-*

echo "编译完成。以下是生成的二进制文件:"
ls -la bin/

echo "你现在可以发布 npm 包了。"