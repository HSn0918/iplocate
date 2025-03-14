package cmd

import (
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
	Use:   "full",
	Short: "先通过IP查询位置，再通过经纬度查询详细信息",
	Long:  `先通过IP地址查询基本位置信息，然后使用获取到的经纬度查询更详细的位置信息。`,
	Run: func(cmd *cobra.Command, args []string) {
		ipAddr, _ := cmd.Flags().GetString("addr")
		if ipAddr == "" {
			utils.Log.Error("错误: 必须提供IP地址")
			cmd.Help()
			return
		}

		utils.Log.Infof("开始完整查询流程，IP: %s", ipAddr)

		locationService := api.NewLocationService()
		// 设置调试模式
		locationService.SetDebug(debugMode)

		// 第一步: 获取IP的位置信息
		utils.Log.Info("步骤1: 查询IP基本位置信息")
		locationData, err := locationService.GetLocationByIP(ipAddr)
		if err != nil {
			utils.Log.Errorf("获取IP位置信息失败: %v", err)
			return
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
			return
		}
		utils.Log.Info("成功获取详细位置信息")
		for _, detailData := range detailDatas {
			utils.PrintLatLngDetailInfo(&detailData)
			// 如果指定了显示原始响应，则打印原始响应信息
			if fullShowRawResponse {
				utils.PrintLatLngDetailRawResponse(&detailData.Detail)
			}
		}

		utils.Log.Info("完整查询流程完成")
	},
}

func init() {
	rootCmd.AddCommand(fullCmd)
	fullCmd.Flags().StringP("addr", "a", "", "要查询的IP地址 (必需)")
	fullCmd.Flags().BoolVarP(&fullShowRawResponse, "raw", "r", false, "显示原始响应信息")
	fullCmd.MarkFlagRequired("addr")
}
