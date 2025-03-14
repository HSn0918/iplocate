package cmd

import (
	"fmt"
	"os"

	"github.com/hsn0918/iplocate/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	// 全局标志
	debugMode   bool
	logFile     string
	showVersion bool
	outputLevel int // 输出级别
)

// rootCmd 表示没有子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:   "iplocate",
	Short: "IP位置查询工具",
	Long: `IPLocate是一个命令行工具，可以通过IP地址或经纬度查询位置信息。
可以查询基本的地理位置信息，如国家、省份、城市等，也可以查询更详细的位置信息。`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在所有命令执行前初始化日志系统
		utils.SetupLogger(debugMode, logFile)

		// 设置输出级别
		utils.SetOutputLevel(utils.OutputLevel(outputLevel))
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 如果指定了版本标志，则显示版本信息
		if showVersion {
			versionCmd.Run(cmd, args)
			return
		}

		// 否则显示帮助信息
		cmd.Help()
	},
}

// Execute 添加所有子命令到根命令并设置标志。
// 这由 main.main() 调用。只需要对 rootCmd 调用一次。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// 在这里，您可以定义根命令的标志和配置设置
	// Cobra 支持持久性标志，如果在这里定义，则对该应用程序可见
	rootCmd.PersistentFlags().StringP("config", "c", "", "配置文件路径 (默认为 $HOME/.iplocate.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "启用调试模式")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log", "l", "", "日志文件路径 (默认输出到控制台)")
	rootCmd.PersistentFlags().IntVarP(&outputLevel, "output-level", "o", 0, "输出级别 (0=基本, 1=正常, 2=详细)")

	// Cobra 还支持本地标志，仅在直接调用此操作时运行
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "显示版本信息")
}
