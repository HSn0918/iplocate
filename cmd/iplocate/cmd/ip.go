package cmd

import (
	"github.com/hsn0918/iplocate/pkg/api"
	"github.com/hsn0918/iplocate/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	// IP命令标志
	showRawResponse bool
)

// ipCmd 表示IP查询命令
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "通过IP地址查询位置信息",
	Long:  `通过IP地址查询位置信息，包括国家、省份、城市等基本地理位置信息。`,
	Run: func(cmd *cobra.Command, args []string) {
		ipAddr, _ := cmd.Flags().GetString("addr")
		if ipAddr == "" {
			utils.Log.Error("错误: 必须提供IP地址")
			cmd.Help()
			return
		}

		utils.Log.Infof("开始查询IP: %s 的位置信息", ipAddr)

		locationService := api.NewLocationService()
		// 设置调试模式
		locationService.SetDebug(debugMode)

		locationData, err := locationService.GetLocationByIP(ipAddr)
		if err != nil {
			utils.Log.Errorf("获取IP位置信息失败: %v", err)
			return
		}

		utils.Log.Infof("成功获取IP: %s 的位置信息", ipAddr)
		utils.PrintIPLocationInfo(locationData)

		// 如果指定了显示原始响应，则打印原始响应信息
		if showRawResponse {
			utils.PrintIPLocationRawResponse(locationData)
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
	ipCmd.Flags().StringP("addr", "a", "", "要查询的IP地址 (必需)")
	ipCmd.Flags().BoolVarP(&showRawResponse, "raw", "r", false, "显示原始响应信息")
	ipCmd.MarkFlagRequired("addr")
}
