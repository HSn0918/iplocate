package cmd

import (
	"github.com/hsn0918/iplocate/pkg/api"
	"github.com/hsn0918/iplocate/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	// 经纬度命令标志
	latlngShowRawResponse bool
)

// latlngCmd 表示经纬度查询命令
var latlngCmd = &cobra.Command{
	Use:   "latlng",
	Short: "通过经纬度查询详细位置信息",
	Long:  `通过经纬度查询详细位置信息，包括国家、省份、城市、区县等详细地理位置信息。`,
	Run: func(cmd *cobra.Command, args []string) {
		lat, _ := cmd.Flags().GetFloat64("lat")
		lng, _ := cmd.Flags().GetFloat64("lng")

		if lat == 0 || lng == 0 {
			utils.Log.Error("错误: 必须提供有效的经纬度")
			cmd.Help()
			return
		}

		utils.Log.Infof("开始查询经纬度: [%f, %f] 的位置信息", lat, lng)

		locationService := api.NewLocationService()
		// 设置调试模式
		locationService.SetDebug(debugMode)

		detailDatas, err := locationService.GetDetailByLatLngWithTags(lat, lng)
		if err != nil {
			utils.Log.Errorf("获取经纬度详细信息失败: %v", err)
			return
		}

		utils.Log.Infof("成功获取经纬度: [%f, %f] 的位置信息", lat, lng)
		for _, detailData := range detailDatas {
			utils.PrintLatLngDetailInfo(&detailData)
			// 如果指定了显示原始响应，则打印原始响应信息
			if latlngShowRawResponse {
				utils.PrintLatLngDetailRawResponse(&detailData.Detail)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(latlngCmd)
	latlngCmd.Flags().Float64P("lat", "t", 0, "纬度值 (必需)")
	latlngCmd.Flags().Float64P("lng", "g", 0, "经度值 (必需)")
	latlngCmd.Flags().BoolVarP(&latlngShowRawResponse, "raw", "r", false, "显示原始响应信息")
	latlngCmd.MarkFlagRequired("lat")
	latlngCmd.MarkFlagRequired("lng")
}
