package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/hsn0918/iplocate/pkg/api"
	"github.com/hsn0918/iplocate/pkg/utils"
)

var (
	// 完整查询命令标志
	fullShowRawResponse bool
)

// fullCmd 表示完整查询命令
var fullCmd = &cobra.Command{
	Use:   "full [IP地址...]",
	Short: "先通过IP查询位置，再通过经纬度查询详细信息",
	Long:  `先通过IP地址查询基本位置信息，然后使用获取到的经纬度查询更详细的位置信息。支持多个IP地址，可以通过 -a 标志或直接作为参数提供，使用空格分隔。`,
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

		utils.Log.Infof("开始完整查询流程，共 %d 个IP地址", len(ipList))

		locationService := api.NewLocationService()
		// 设置调试模式
		locationService.SetDebug(debugMode)

		// 遍历每个IP地址进行查询
		for _, ipAddr := range ipList {
			utils.Log.Infof("开始处理IP: %s", ipAddr)

			// 第一步: 获取IP的位置信息
			utils.Log.Infof("步骤1: 查询IP %s 的基本位置信息", ipAddr)
			locationData, err := locationService.GetLocationByIP(ipAddr)
			if err != nil {
				utils.Log.Errorf("获取IP位置信息失败: %v", err)
				continue
			}
			utils.Log.Infof("成功获取IP: %s 的基本位置信息", ipAddr)
			utils.PrintIPLocationInfo(locationData)

			// 如果指定了显示原始响应，则打印原始响应信息
			if fullShowRawResponse {
				utils.PrintIPLocationRawResponse(locationData)
			}

			// 第二步: 使用经纬度获取详细信息
			utils.Log.Infof("步骤2: 使用经纬度 [%f, %f] 查询详细位置信息", locationData.Lat, locationData.Lng)

			detailDatas, err := locationService.GetDetailByLatLngWithTags(locationData.Lat, locationData.Lng)
			if err != nil {
				utils.Log.Errorf("获取经纬度详细信息失败: %v", err)
				continue
			}
			utils.Log.Info("成功获取详细位置信息")
			for _, detailData := range detailDatas {
				utils.PrintLatLngDetailInfo(&detailData)
				// 如果指定了显示原始响应，则打印原始响应信息
				if fullShowRawResponse {
					utils.PrintLatLngDetailRawResponse(&detailData.Detail)
				}
			}

			utils.Log.Infof("IP %s 的完整查询流程完成", ipAddr)
		}

		utils.Log.Info("所有IP地址的完整查询流程完成")
	},
}

func init() {
	rootCmd.AddCommand(fullCmd)
	fullCmd.Flags().StringP("addr", "a", "", "要查询的IP地址，多个IP使用空格分隔")
	fullCmd.Flags().BoolVarP(&fullShowRawResponse, "raw", "r", false, "显示原始响应信息")
}
