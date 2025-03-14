package cmd

import (
	"fmt"

	"github.com/hsn0918/iplocate/pkg/utils"
	"github.com/spf13/cobra"
)

// 版本信息
var (
	Version   = "0.0.1"
	BuildDate = "2024-03-14"
	GitCommit = "unknown"
)

// versionCmd 表示版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  `显示IPLocate位置查询工具的版本信息，包括版本号、构建日期和Git提交哈希。`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Log.Info("显示版本信息")
		fmt.Printf("IPLocate位置查询工具 v%s\n", Version)
		fmt.Printf("构建日期: %s\n", BuildDate)
		fmt.Printf("Git提交: %s\n", GitCommit)
		fmt.Println("依赖库:")
		fmt.Println("  - github.com/spf13/cobra")
		fmt.Println("  - github.com/go-resty/resty/v2")
		fmt.Println("  - github.com/sirupsen/logrus")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
