#!/bin/bash

# 自动发布新版本的脚本
# 使用方法: ./scripts/release.sh [版本号]
# 例如: ./scripts/release.sh 1.0.1

set -e

# 检查是否提供了版本号
if [ -z "$1" ]; then
  echo "错误: 必须提供版本号"
  echo "使用方法: $0 [版本号]"
  echo "例如: $0 1.0.1"
  exit 1
fi

VERSION=$1
VERSION_TAG="v$VERSION"

# 检查当前目录是否是项目根目录
if [ ! -f "go.mod" ] || [ ! -d ".git" ]; then
  echo "错误: 请在项目根目录运行此脚本"
  exit 1
fi

# 检查工作目录是否干净
if [ -n "$(git status --porcelain)" ]; then
  echo "错误: 工作目录不干净，请先提交或暂存所有更改"
  git status
  exit 1
fi

# 更新版本号
echo "更新版本号为 $VERSION..."
sed -i '' "s/VERSION := .*/VERSION := $VERSION/" Makefile

# 提交更改
echo "提交版本更新..."
git add Makefile
git commit -m "chore: 更新版本号为 $VERSION"

# 创建标签
echo "创建标签 $VERSION_TAG..."
git tag -a "$VERSION_TAG" -m "版本 $VERSION"

# 推送到远程仓库
echo "推送更改和标签到远程仓库..."
git push origin master
git push origin "$VERSION_TAG"

echo "✅ 发布流程已启动!"
echo "GitHub Actions 将自动构建并发布 release。"
echo "请访问 GitHub 仓库的 Actions 页面查看进度。"