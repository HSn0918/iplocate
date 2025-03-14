# IPLocate API 位置查询工具

这是一个基于 API 的位置查询命令行工具，可以通过 IP 地址或经纬度查询位置信息。

## 功能特点

- 通过 IP 地址查询基本位置信息
- 通过经纬度查询详细位置信息

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
