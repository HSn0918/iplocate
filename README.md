# IPLocate API Location Query Tool

A command-line tool based on API for querying location information through IP addresses or latitude and longitude coordinates.

<p align="right">
  <a href="#iplocate-api-位置查询工具">中文</a> | <b>English</b>
</p>

## Features

- Query basic location information by IP address
- Query detailed location information by latitude and longitude coordinates
- Support for multiple IP addresses (space-separated)
- MCP (Model Context Protocol) support for integration with LLM applications

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

### MCP Service

IPLocate now supports the Model Context Protocol (MCP), which allows it to be used as a tool by LLM applications:

```bash
# Start the IPLocate MCP service
./iplocate mcp

# Run with debug mode
./iplocate mcp -d

# Save logs to a file
./iplocate mcp -l mcp_service.log
```

#### Using the NPM Package

For easier integration with LLM applications, you can use our npm package:

```bash
# Install globally
npm install -g @hsn0918/iplocate-mcp

# Or use directly with npx
npx @hsn0918/iplocate-mcp
```

The npm package includes pre-compiled binaries for:
- Windows (x64)
- macOS (Intel/Apple Silicon)
- Linux (x64/ARM64)

#### MCP Tools

The MCP service provides the following tools:

1. **ip_location** - Query location by IP address
   - Parameter: `ip` (string) - The IP address to query

2. **latlng_location** - Query location by latitude and longitude
   - Parameter: `lat` (number) - Latitude value
   - Parameter: `lng` (number) - Longitude value

#### Integrating with LLM Applications

To integrate IPLocate into LLM applications, add it to your MCP configuration:

```json
{
  "mcpServers": {
    "iplocate-mcp": {
      "isActive": true,
      "name": "iplocate-mcp",
      "type": "stdio",
      "description": "IP地址和经纬度位置查询服务，提供全球IP地址位置信息查询和经纬度详细地址查询功能",
      "command": "npx",
      "args": [
        "-y",
        "@hsn0918/iplocate-mcp@1.0.1"
      ],
      "env": {}
    }
  }
}
```

This configuration uses the npm package, which makes it easy to deploy in any environment without needing to install Go or compile the source code.

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

mcp:
  (Uses global options)
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
│   ├── mcp_config.json # MCP configuration example
├── npm-package/        # NPM package for MCP service
│   ├── index.js        # NPM package entry point
│   ├── package.json    # NPM package definition
│   └── bin/            # Precompiled binaries
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

# Test MCP functionality
npx @hsn0918/iplocate-mcp
```

## Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - Command-line interface library
- [github.com/go-resty/resty](https://github.com/go-resty/resty) - Simple HTTP and REST client library
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Structured logging library
- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) - Model Context Protocol implementation for Go

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
- 支持 MCP（Model Context Protocol）协议，可集成到 LLM 应用中

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

### MCP 服务

IPLocate 现在支持模型上下文协议（MCP），可以作为 LLM 应用的工具使用：

```bash
# 启动 IPLocate MCP 服务
./iplocate mcp

# 使用调试模式运行
./iplocate mcp -d

# 将日志保存到文件
./iplocate mcp -l mcp_service.log
```

#### 使用 NPM 包

为了更容易与 LLM 应用集成，您可以使用我们的 npm 包：

```bash
# 全局安装
npm install -g @hsn0918/iplocate-mcp

# 或者直接使用 npx 运行
npx @hsn0918/iplocate-mcp
```

npm 包包含以下预编译的二进制文件：
- Windows (x64)
- macOS (Intel/Apple Silicon)
- Linux (x64/ARM64)

#### MCP 工具

MCP 服务提供以下工具：

1. **ip_location** - 通过 IP 地址查询位置
   - 参数：`ip`（字符串）- 要查询的 IP 地址

2. **latlng_location** - 通过经纬度查询位置
   - 参数：`lat`（数字）- 纬度值
   - 参数：`lng`（数字）- 经度值

#### 与 LLM 应用集成

要将 IPLocate 集成到 LLM 应用中，请将其添加到您的 MCP 配置中：

```json
{
  "mcpServers": {
    "iplocate-mcp": {
      "isActive": true,
      "name": "iplocate-mcp",
      "type": "stdio",
      "description": "IP地址和经纬度位置查询服务，提供全球IP地址位置信息查询和经纬度详细地址查询功能",
      "command": "npx",
      "args": [
        "-y",
        "@hsn0918/iplocate-mcp@1.0.1"
      ],
      "env": {}
    }
  }
}
```

这个配置使用 npm 包，使得在任何环境中部署都很容易，无需安装 Go 或编译源代码。

### 自动补全

该工具支持为 Bash、Zsh、Fish 和 PowerShell 生成自动补全脚本：

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

# Test MCP functionality
npx @hsn0918/iplocate-mcp
```

## 依赖库

- [github.com/spf13/cobra](https://github.com/spf13/cobra) - 命令行界面库
- [github.com/go-resty/resty](https://github.com/go-resty/resty) - 简单的 HTTP 和 REST 客户端库
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - 结构化日志库
- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) - Model Context Protocol implementation for Go

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
