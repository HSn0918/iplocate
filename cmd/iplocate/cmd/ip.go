package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/hsn0918/iplocate/pkg/api"
	"github.com/hsn0918/iplocate/pkg/utils"
)

var (
	// IP命令标志
	showRawResponse bool
)

// ipCmd 表示IP查询命令
var ipCmd = &cobra.Command{
	Use:   "ip [IP地址...]",
	Short: "通过IP地址查询位置信息",
	Long:  `通过IP地址查询位置信息，包括国家、省份、城市等基本地理位置信息。支持多个IP地址，可以通过 -a 标志或直接作为参数提供，使用空格分隔。`,
	Run: func(cmd *cobra.Command, args []string) {
		ipAddrs, _ := cmd.Flags().GetString("addr")
		var ipList []string

		// 如果提供了 -a 标志，解析其中的 IP 地址
		if ipAddrs != "" {
			ipList = append(ipList, strings.Fields(ipAddrs)...)
		}

		// 添加位置参数中的 IP 地址
		ipList = append(ipList, args...)

		if len(ipList) == 0 {
			utils.Log.Error("错误: 必须提供至少一个IP地址")
			cmd.Help()
			return
		}

		utils.Log.Infof("开始查询 %d 个IP地址的位置信息", len(ipList))

		locationService := api.NewLocationService()
		// 设置调试模式
		locationService.SetDebug(debugMode)

		// 遍历每个IP地址进行查询
		for _, ipAddr := range ipList {
			utils.Log.Infof("查询IP: %s 的位置信息", ipAddr)

			locationData, err := locationService.GetLocationByIP(ipAddr)
			if err != nil {
				utils.Log.Errorf("获取IP位置信息失败: %v", err)
				continue
			}

			utils.Log.Infof("成功获取IP: %s 的位置信息", ipAddr)
			utils.PrintIPLocationInfo(locationData)

			// 如果指定了显示原始响应，则打印原始响应信息
			if showRawResponse {
				utils.PrintIPLocationRawResponse(locationData)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
	ipCmd.Flags().StringP("addr", "a", "", "要查询的IP地址，多个IP使用空格分隔")
	ipCmd.Flags().BoolVarP(&showRawResponse, "raw", "r", false, "显示原始响应信息")
}
