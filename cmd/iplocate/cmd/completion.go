package cmd

import (
	"os"

	"github.com/hsn0918/iplocate/pkg/utils"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "生成自动补全脚本",
	Long: `生成IPLocate位置查询工具的自动补全脚本。

支持以下 shell 环境:
  - bash
  - zsh
  - fish
  - powershell

使用示例:
  # Bash
  $ source <(iplocate completion bash)
  # 或者添加到 .bashrc 文件中
  $ iplocate completion bash > ~/.bash_completion

  # Zsh
  $ source <(iplocate completion zsh)
  # 或者添加到 .zshrc 文件中
  $ iplocate completion zsh > "${fpath[1]}/_iplocate"

  # Fish
  $ iplocate completion fish > ~/.config/fish/completions/iplocate.fish

  # PowerShell
  $ iplocate completion powershell > iplocate.ps1
  $ . ./iplocate.ps1
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}

		if err != nil {
			utils.Log.Errorf("生成自动补全脚本失败: %v", err)
		} else {
			utils.Log.Info("自动补全脚本生成成功")
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
