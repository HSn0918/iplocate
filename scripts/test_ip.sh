#!/bin/bash

# 测试 IP 查询功能的脚本

# 确保脚本在错误时退出
set -e

# 编译程序
echo "编译程序..."
go build -o meitu cmd/meitu/main.go

# 测试 IP 查询
echo "测试 IP 查询功能..."
./meitu ip -a 60.191.18.194

# 测试 IP 查询（调试模式）
echo -e "\n测试 IP 查询功能（调试模式）..."
./meitu ip -a 60.191.18.194 -d

# 测试 IP 查询（显示原始响应）
echo -e "\n测试 IP 查询功能（显示原始响应）..."
./meitu ip -a 60.191.18.194 -r

# 测试完整查询
echo -e "\n测试完整查询功能..."
./meitu full -a 60.191.18.194

# 测试完整查询（显示原始响应）
echo -e "\n测试完整查询功能（显示原始响应）..."
./meitu full -a 60.191.18.194 -r

echo -e "\n测试完成！"