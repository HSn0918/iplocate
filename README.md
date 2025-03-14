# IPLocate API 位置查询工具

这是一个基于 API 的位置查询命令行工具，可以通过 IP 地址或经纬度查询位置信息。

## 功能特点

- 通过 IP 地址查询基本位置信息
- 通过经纬度查询详细位置信息
- 支持完整查询流程（先 IP 查询，再经纬度查询）
- 支持调试模式和日志记录
- 使用 Resty 库进行高效的 HTTP 请求
- 使用 Logrus 库进行结构化日志记录
- 支持命令行自动补全功能
- 支持显示原始 HTTP 响应数据
- 提供 Makefile 简化常用操作

## 安装

```bash
# 克隆仓库
git clone https://github.com/hsn0918/iplocate.git
cd iplocate

# 使用 Makefile 编译
make build

# 安装到系统目录（需要管理员权限）
sudo make install

# 或者安装到用户本地目录（不需要管理员权限）
make install-local

# 卸载（如果安装在系统目录，需要管理员权限）
sudo make uninstall

# 从用户本地目录卸载
make uninstall-local

# 或者使用 Go 命令安装
go install github.com/hsn0918/iplocate/cmd/iplocate@latest
```

## 使用方法

### 使用 Makefile

项目提供了 Makefile 来简化常用操作：

```bash
# 编译项目
make build

# 运行单元测试
make test-unit

# 运行集成测试
make test-integration

# 运行所有测试
make test

# 测试 IP 查询功能
make test-ip

# 测试 IP 查询功能（显示原始响应）
make test-ip-raw

# 测试经纬度查询功能
make test-latlng

# 测试经纬度查询功能（显示原始响应）
make test-latlng-raw

# 测试完整查询功能
make test-full

# 测试完整查询功能（显示原始响应）
make test-full-raw

# 生成 Bash 自动补全脚本
make completion-bash

# 生成 Zsh 自动补全脚本
make completion-zsh

# 安装到系统目录（需要管理员权限）
sudo make install

# 安装到用户本地目录（不需要管理员权限）
make install-local

# 从系统目录卸载（需要管理员权限）
sudo make uninstall

# 从用户本地目录卸载
make uninstall-local

# 清理生成的文件
make clean

# 显示帮助信息
make help
```

### IP 查询

```bash
# 查询指定 IP 的位置信息
./iplocate ip -a 114.114.114.114

# 启用调试模式
./iplocate ip -a 114.114.114.114 -d

# 显示原始响应数据
./iplocate ip -a 114.114.114.114 -r

# 将日志输出到文件
./iplocate ip -a 114.114.114.114 -l ip_query.log
```

### 经纬度查询

```bash
# 查询指定经纬度的位置信息
./iplocate latlng -t 39.9042 -g 116.4074

# 启用调试模式
./iplocate latlng -t 39.9042 -g 116.4074 -d

# 显示原始响应数据
./iplocate latlng -t 39.9042 -g 116.4074 -r
```

### 完整查询

```bash
# 先通过 IP 查询位置，再通过经纬度查询详细信息
./iplocate full -a 114.114.114.114

# 启用调试模式
./iplocate full -a 114.114.114.114 -d

# 显示原始响应数据
./iplocate full -a 114.114.114.114 -r
```

### 自动补全

工具支持为 Bash、Zsh、Fish 和 PowerShell 生成自动补全脚本：

```bash
# 生成 Bash 自动补全脚本并应用
source <(./iplocate completion bash)

# 永久添加到 Bash 配置
./iplocate completion bash > ~/.bash_completion

# 生成 Zsh 自动补全脚本并应用
source <(./iplocate completion zsh)

# 永久添加到 Zsh 配置
./iplocate completion zsh > "${fpath[1]}/_iplocate"

# 生成 Fish 自动补全脚本
./iplocate completion fish > ~/.config/fish/completions/iplocate.fish

# 生成 PowerShell 自动补全脚本
./iplocate completion powershell > iplocate.ps1
. ./iplocate.ps1
```

### 全局选项

```
-d, --debug          启用调试模式
-l, --log string     日志文件路径 (默认输出到控制台)
-o, --output-level   输出级别 (0=基本, 1=正常, 2=详细，默认为0)
-c, --config string  配置文件路径 (默认为 $HOME/.iplocate.yaml)
-h, --help           查看帮助信息
```

### 命令特定选项

```
ip:
  -a, --addr string   要查询的IP地址 (必需)
  -r, --raw           显示原始响应信息

latlng:
  -t, --lat float     纬度值 (必需)
  -g, --lng float     经度值 (必需)
  -r, --raw           显示原始响应信息

full:
  -a, --addr string   要查询的IP地址 (必需)
  -r, --raw           显示原始响应信息
```

## 开发

### 项目结构

```
iplocate/
├── cmd/                # 命令行应用
│   └── iplocate/       # 主程序
│       ├── cmd/        # 命令定义
│       └── main.go     # 程序入口
├── pkg/                # 包目录
│   ├── api/            # API 服务
│   ├── models/         # 数据模型
│   └── utils/          # 工具函数
├── scripts/            # 脚本文件
├── go.mod              # Go 模块定义
├── go.sum              # Go 模块校验和
├── Makefile            # 构建脚本
└── README.md           # 项目说明
```

### 测试

```bash
# 运行单元测试
go test ./pkg/...

# 运行集成测试
go test -tags=integration ./pkg/...

# 或者使用 Makefile
make test
```

## 依赖库

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - 命令行界面库
- [github.com/go-resty/resty](https://github.com/go-resty/resty) - 简单的 HTTP 和 REST 客户端库
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - 结构化日志库

## 许可证

MIT