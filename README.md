# IPLocate API Location Query Tool

A command-line tool based on API for querying location information through IP addresses or latitude and longitude coordinates.

<p align="right">
  <a href="#iplocate-api-位置查询工具">中文</a> | <b>English</b>
</p>

## Features

- Query basic location information by IP address
- Query detailed location information by latitude and longitude coordinates
- Support for multiple IP addresses (space-separated)

## Installation

```bash
# Clone the repository
git clone https://github.com/hsn0918/iplocate.git
cd iplocate

# Compile using Makefile
make build

# Install to system directory (requires administrator privileges)
sudo make install

# Or install to user local directory (no administrator privileges required)
make install-local

# Uninstall (if installed in system directory, requires administrator privileges)
sudo make uninstall

# Uninstall from user local directory
make uninstall-local

# Or install using Go command
go install github.com/hsn0918/iplocate/cmd/iplocate@latest
```

## Usage

### IP Query

```bash
# Query location information for a specific IP
./iplocate ip -a 114.114.114.114

# Query multiple IP addresses
./iplocate ip -a "114.114.114.114 8.8.8.8"

# Or specify IPs directly as arguments
./iplocate ip 114.114.114.114 8.8.8.8

# Enable debug mode
./iplocate ip -a 114.114.114.114 -d

# Show raw response data
./iplocate ip -a 114.114.114.114 -r

# Output logs to a file
./iplocate ip -a 114.114.114.114 -l ip_query.log
```

### Latitude and Longitude Query

```bash
# Query location information for specific coordinates
./iplocate latlng -t 39.9042 -g 116.4074

# Enable debug mode
./iplocate latlng -t 39.9042 -g 116.4074 -d

# Show raw response data
./iplocate latlng -t 39.9042 -g 116.4074 -r
```

### Full Query

```bash
# First query location by IP, then query detailed information by coordinates
./iplocate full -a 114.114.114.114

# Query multiple IP addresses
./iplocate full -a "114.114.114.114 8.8.8.8"

# Or specify IPs directly as arguments
./iplocate full 114.114.114.114 8.8.8.8

# Enable debug mode
./iplocate full -a 114.114.114.114 -d

# Show raw response data
./iplocate full -a 114.114.114.114 -r
```

### Auto-completion

The tool supports generating auto-completion scripts for Bash, Zsh, Fish, and PowerShell:

```bash
# Generate and apply Bash auto-completion script
source <(iplocate completion bash)

# Permanently add to Bash configuration
iplocate completion bash > ~/.bash_completion

# Generate and apply Zsh auto-completion script
source <(iplocate completion zsh)

# Permanently add to Zsh configuration
iplocate completion zsh > "${fpath[1]}/_iplocate"

# Generate Fish auto-completion script
iplocate completion fish > ~/.config/fish/completions/iplocate.fish

# Generate PowerShell auto-completion script
iplocate completion powershell > iplocate.ps1
. iplocate.ps1
```

### Global Options

```
-d, --debug          Enable debug mode
-l, --log string     Log file path (default output to console)
-o, --output-level   Output level (0=basic, 1=normal, 2=detailed, default is 0)
-c, --config string  Configuration file path (default is $HOME/.iplocate.yaml)
-h, --help           View help information
```

### Command-specific Options

```
ip:
  -a, --addr string   IP address to query (required)
  -r, --raw           Show raw response information

latlng:
  -t, --lat float     Latitude value (required)
  -g, --lng float     Longitude value (required)
  -r, --raw           Show raw response information

full:
  -a, --addr string   IP address to query (required)
  -r, --raw           Show raw response information
```

## Development

### Project Structure

```
iplocate/
├── cmd/                # Command-line application
│   └── iplocate/       # Main program
│       ├── cmd/        # Command definitions
│       └── main.go     # Program entry
├── pkg/                # Package directory
│   ├── api/            # API services
│   ├── models/         # Data models
│   └── utils/          # Utility functions
├── scripts/            # Script files
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Makefile            # Build scripts
└── README.md           # Project description
```

### Testing

```bash
# Run unit tests
go test ./pkg/...

# Run integration tests
go test -tags=integration ./pkg/...

# Or use Makefile
make test
```

## Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - Command-line interface library
- [github.com/go-resty/resty](https://github.com/go-resty/resty) - Simple HTTP and REST client library
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Structured logging library

## Disclaimer and Usage Restrictions

### Disclaimer

This tool is for learning and research purposes only. The developer is not responsible for any consequences arising from the use of this tool. By using this tool, you agree to assume all risks of use.

### Usage Restrictions

1. **No Abuse**: It is strictly prohibited to use this tool for any illegal activities or actions that infringe on the rights of others.
2. **API Usage Restrictions**: Please comply with the terms of use and restrictions of relevant API providers.
3. **Data Protection**: Data obtained using this tool should comply with relevant data protection regulations.
4. **Commercial Use**: For commercial use, please first obtain authorization from the relevant API providers.

### Legal Notice

The APIs and services used in this tool are owned by their respective owners. This tool has no affiliation with these service providers. All trademarks and service marks are the property of their respective owners.

By using this tool, you acknowledge that you have read and agree to the above disclaimer and usage restrictions. If you do not agree, please stop using this tool immediately.

## License

MIT

---

# IPLocate API 位置查询工具

<p align="right">
  <b>中文</b> | <a href="#iplocate-api-location-query-tool">English</a>
</p>

这是一个基于 API 的位置查询命令行工具，可以通过 IP 地址或经纬度查询位置信息。

## 功能特点

- 通过 IP 地址查询基本位置信息
- 通过经纬度查询详细位置信息
- 支持多个IP地址查询（空格分隔）

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

### IP 查询

```bash
# 查询指定 IP 的位置信息
./iplocate ip -a 114.114.114.114

# 查询多个IP地址
./iplocate ip -a "114.114.114.114 8.8.8.8"

# 或者直接在命令后面指定IP
./iplocate ip 114.114.114.114 8.8.8.8

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

# 查询多个IP地址
./iplocate full -a "114.114.114.114 8.8.8.8"

# 或者直接在命令后面指定IP
./iplocate full 114.114.114.114 8.8.8.8

# 启用调试模式
./iplocate full -a 114.114.114.114 -d

# 显示原始响应数据
./iplocate full -a 114.114.114.114 -r
```

### 自动补全

工具支持为 Bash、Zsh、Fish 和 PowerShell 生成自动补全脚本：

```bash
# 生成 Bash 自动补全脚本并应用
source <(iplocate completion bash)

# 永久添加到 Bash 配置
iplocate completion bash > ~/.bash_completion

# 生成 Zsh 自动补全脚本并应用
source <(iplocate completion zsh)

# 永久添加到 Zsh 配置
iplocate completion zsh > "${fpath[1]}/_iplocate"

# 生成 Fish 自动补全脚本
iplocate completion fish > ~/.config/fish/completions/iplocate.fish

# 生成 PowerShell 自动补全脚本
iplocate completion powershell > iplocate.ps1
. iplocate.ps1
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

## 免责声明与使用限制

### 免责声明

本工具仅供学习和研究目的使用。开发者不对使用本工具产生的任何后果负责。使用本工具即表示您同意自行承担使用风险。

### 使用限制

1. **禁止滥用**：严禁将本工具用于任何非法活动或侵犯他人权益的行为。
2. **API使用限制**：请遵守相关API提供商的使用条款和限制。
3. **数据保护**：使用本工具获取的数据应当遵守相关数据保护法规。
4. **商业使用**：如需商业使用，请先获得相关API提供商的授权。

### 法律声明

本工具中使用的API和服务归其各自所有者所有。本工具与这些服务提供商没有任何关联。所有商标和服务标记均为其各自所有者的财产。

使用本工具即表示您已阅读并同意上述免责声明和使用限制。如不同意，请立即停止使用本工具。

## 许可证

MIT
