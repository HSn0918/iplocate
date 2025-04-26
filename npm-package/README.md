# iplocate-mcp

IP地址和经纬度位置查询 MCP (Model Context Protocol) 服务

## 简介

这个包提供了一个 MCP 服务，可以让大语言模型 (LLM) 通过 IP 地址或经纬度查询位置信息。它是 [iplocate](https://github.com/hsn0918/iplocate) 工具的 MCP 封装版本。

## 功能

- 通过 IP 地址查询基本位置信息
- 通过经纬度查询详细位置信息
- 支持标准的 MCP 协议，易于与 LLM 应用集成

## 安装

```bash
# 全局安装
npm install -g @hsn0918/iplocate-mcp

# 或者使用 npx 运行而不安装
npx @hsn0918/iplocate-mcp
```

## 前提条件

- Node.js 14.0 或更高版本
- 支持的操作系统:
  - Windows (x64)
  - macOS (Intel/Apple Silicon)
  - Linux (x64/ARM64)

本包包含预编译的 iplocate 二进制文件，无需额外依赖 Go 环境。

## 使用方法

### 直接运行

```bash
# 如果全局安装了包
iplocate-mcp

# 或者使用 npx
npx @hsn0918/iplocate-mcp
```

### 与 LLM 应用集成

在您的 MCP 配置中添加：

```json
{
  "iplocate-mcp": {
    "isActive": true,
    "name": "iplocate-mcp",
    "type": "stdio",
    "description": "IP地址和经纬度位置查询服务，提供全球IP地址位置信息查询和经纬度详细地址查询功能",
    "command": "npx",
    "args": [
      "-y",
      "@hsn0918/iplocate-mcp"
    ],
    "env": {}
  }
}
```

## 可用工具

这个 MCP 服务提供两个工具：

1. **ip_location** - 通过 IP 地址查询位置
   - 参数：`ip`（字符串）- 要查询的 IP 地址

2. **latlng_location** - 通过经纬度查询位置
   - 参数：`lat`（数字）- 纬度值
   - 参数：`lng`（数字）- 经度值

## 示例

### IP 位置查询

```json
{
  "id": "1",
  "type": "callTool",
  "params": {
    "name": "ip_location",
    "arguments": {
      "ip": "8.8.8.8"
    }
  }
}
```

### 经纬度位置查询

```json
{
  "id": "2",
  "type": "callTool",
  "params": {
    "name": "latlng_location",
    "arguments": {
      "lat": 40.7128,
      "lng": -74.006
    }
  }
}
```

## 故障排除

- 如果遇到权限问题，尝试使用 `sudo` 运行
- 确保您的操作系统和架构是受支持的
- 如果遇到"找不到可执行文件"错误，请检查 bin 目录是否存在且包含正确的二进制文件

## 许可证

MIT