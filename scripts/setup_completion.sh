#!/bin/bash

# IPLocate 命令行补全设置脚本

set -e

# 检测shell类型
SHELL_TYPE=$(basename "$SHELL")
CONFIG_FILE=""
COMPLETION_CMD=""

case "$SHELL_TYPE" in
    bash)
        CONFIG_FILE="$HOME/.bashrc"
        COMPLETION_CMD="source <(iplocate completion bash)"
        ;;
    zsh)
        CONFIG_FILE="$HOME/.zshrc"
        COMPLETION_CMD="source <(iplocate completion zsh)"
        ;;
    fish)
        CONFIG_FILE="$HOME/.config/fish/config.fish"
        echo "Fish shell 需要手动设置补全:"
        echo "iplocate completion fish > ~/.config/fish/completions/iplocate.fish"
        exit 0
        ;;
    *)
        echo "不支持的shell类型: $SHELL_TYPE"
        echo "请手动设置命令行补全:"
        echo "Bash: source <(iplocate completion bash)"
        echo "Zsh: source <(iplocate completion zsh)"
        echo "Fish: iplocate completion fish > ~/.config/fish/completions/iplocate.fish"
        echo "PowerShell: iplocate completion powershell > iplocate.ps1 && . ./iplocate.ps1"
        exit 1
        ;;
esac

# 检查配置文件是否存在
if [ ! -f "$CONFIG_FILE" ]; then
    echo "配置文件不存在: $CONFIG_FILE"
    echo "创建新文件..."
    touch "$CONFIG_FILE"
fi

# 检查是否已经添加了补全命令
if grep -q "iplocate completion" "$CONFIG_FILE"; then
    echo "命令行补全已经设置在 $CONFIG_FILE 中"
else
    echo "添加命令行补全到 $CONFIG_FILE..."
    echo "" >> "$CONFIG_FILE"
    echo "# IPLocate 命令行补全" >> "$CONFIG_FILE"
    echo "$COMPLETION_CMD" >> "$CONFIG_FILE"
    echo "命令行补全设置成功!"
    echo "请重新加载配置文件或重启终端以应用更改:"
    echo "source $CONFIG_FILE"
fi

echo ""
echo "您也可以手动添加以下命令到您的shell配置文件中:"
echo "$COMPLETION_CMD"