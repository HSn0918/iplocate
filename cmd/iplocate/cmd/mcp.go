package cmd

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"

	"github.com/hsn0918/iplocate/pkg/api"
	"github.com/hsn0918/iplocate/pkg/utils"
)

// mcpCmd 代表MCP服务命令
var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "启动IP位置查询MCP服务",
	Long:  `启动IP位置查询MCP服务，通过标准输入输出与LLM应用进行通信，提供IP地址查询和经纬度位置查询功能。`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Log.Info("启动 IP 位置查询 MCP 服务...")

		// 创建MCP服务器
		s := server.NewMCPServer(
			"IPLocate MCP Service",
			"1.0.0",
		)

		// 创建并添加IP位置查询工具
		ipTool := mcp.NewTool("ip_location",
			mcp.WithDescription("通过IP地址查询位置信息"),
			mcp.WithString("ip",
				mcp.Required(),
				mcp.Description("要查询的IP地址"),
			),
		)

		// 添加IP位置查询工具处理函数
		s.AddTool(ipTool, ipLocationHandler)

		// 创建并添加经纬度查询工具
		latlngTool := mcp.NewTool("latlng_location",
			mcp.WithDescription("通过经纬度查询位置信息"),
			mcp.WithNumber("lat",
				mcp.Required(),
				mcp.Description("纬度"),
			),
			mcp.WithNumber("lng",
				mcp.Required(),
				mcp.Description("经度"),
			),
		)

		// 添加经纬度查询工具处理函数
		s.AddTool(latlngTool, latLngLocationHandler)

		// 启动标准输入输出服务
		utils.Log.Info("MCP服务已启动，等待输入...")
		if err := server.ServeStdio(s); err != nil {
			utils.Log.Errorf("MCP服务错误: %v", err)
		}
	},
}

// IP位置查询工具处理函数
func ipLocationHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 从参数中获取IP地址
	ip, ok := request.Params.Arguments["ip"].(string)
	if !ok {
		return mcp.NewToolResultError("IP地址必须是字符串"), nil
	}

	utils.Log.Infof("正在查询IP: %s 的位置信息", ip)

	// 创建位置服务
	locationService := api.NewLocationService()
	locationService.SetDebug(debugMode)

	// 获取位置信息
	locationData, err := locationService.GetLocationByIP(ip)
	if err != nil {
		errMsg := fmt.Sprintf("获取IP位置信息失败: %v", err)
		utils.Log.Error(errMsg)
		return mcp.NewToolResultError(errMsg), nil
	}

	utils.Log.Infof("成功获取IP: %s 的位置信息", ip)

	// 构建位置信息文本
	result := fmt.Sprintf("IP: %s\n经度: %f\n纬度: %f\n国家: %s\n省份: %s\n城市: %s\n区县: %s",
		locationData.IP,
		locationData.Lng,
		locationData.Lat,
		locationData.Rgeo.Country,
		locationData.Rgeo.Province,
		locationData.Rgeo.City,
		locationData.Rgeo.District,
	)

	return mcp.NewToolResultText(result), nil
}

// 经纬度位置查询工具处理函数
func latLngLocationHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 从参数中获取经纬度
	lat, ok := request.Params.Arguments["lat"].(float64)
	if !ok {
		return mcp.NewToolResultError("纬度必须是数字"), nil
	}

	lng, ok := request.Params.Arguments["lng"].(float64)
	if !ok {
		return mcp.NewToolResultError("经度必须是数字"), nil
	}

	utils.Log.Infof("正在查询经度: %f, 纬度: %f 的位置信息", lng, lat)

	// 创建位置服务
	locationService := api.NewLocationService()
	locationService.SetDebug(debugMode)

	// 获取位置信息
	taggedDetails, err := locationService.GetDetailByLatLngWithTags(lat, lng)
	if err != nil {
		errMsg := fmt.Sprintf("获取经纬度位置信息失败: %v", err)
		utils.Log.Error(errMsg)
		return mcp.NewToolResultError(errMsg), nil
	}

	utils.Log.Infof("成功获取经度: %f, 纬度: %f 的位置信息", lng, lat)

	// 构建位置信息文本
	var result string
	if len(taggedDetails) > 0 {
		// 使用第一个结果作为基本信息
		detail := taggedDetails[0].Detail
		result = fmt.Sprintf("经度: %f\n纬度: %f\n国家: %s\n省份: %s\n城市: %s\n区县: %s\n详细地址: %s",
			detail.Lng,
			detail.Lat,
			detail.Country,
			detail.Province,
			detail.City,
			detail.District,
			detail.Detail,
		)

		// 如果有更多的tag结果，附加到后面
		if len(taggedDetails) > 1 {
			result += "\n\n其他详细信息:"
			for i, tagged := range taggedDetails[1:] {
				if i < 3 { // 只显示前三个额外的tag结果
					result += fmt.Sprintf("\n[Tag %d] %s", tagged.Tag, tagged.Detail.Detail)
				} else {
					break
				}
			}
		}
	} else {
		result = "未能获取到有效的位置信息"
	}

	return mcp.NewToolResultText(result), nil
}

func init() {
	rootCmd.AddCommand(mcpCmd)
}
