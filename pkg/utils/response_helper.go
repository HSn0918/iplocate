package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/hsn0918/iplocate/pkg/models"
)

// GetRawResponseBody 获取原始响应体
func GetRawResponseBody(resp *resty.Response) string {
	if resp == nil {
		return ""
	}
	return string(resp.Body())
}

// GetRawResponseHeaders 获取原始响应头
func GetRawResponseHeaders(resp *resty.Response) string {
	if resp == nil {
		return ""
	}

	var headers strings.Builder
	for k, v := range resp.Header() {
		headers.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ", ")))
	}
	return headers.String()
}

// GetRawResponseInfo 获取原始响应信息
func GetRawResponseInfo(resp *resty.Response) string {
	if resp == nil {
		return ""
	}

	var info strings.Builder
	info.WriteString(fmt.Sprintf("状态码: %d\n", resp.StatusCode()))
	info.WriteString(fmt.Sprintf("请求耗时: %v\n", resp.Time()))
	info.WriteString(fmt.Sprintf("响应大小: %d bytes\n", len(resp.Body())))
	info.WriteString(fmt.Sprintf("请求URL: %s\n", resp.Request.URL))
	info.WriteString(fmt.Sprintf("请求方法: %s\n", resp.Request.Method))
	return info.String()
}

// PrintIPLocationRawResponse 打印IP位置原始响应
func PrintIPLocationRawResponse(data *models.IPLocationData) {
	if data == nil || data.RawResponse == nil {
		fmt.Println("无原始响应数据")
		return
	}

	fmt.Println("\n======== 原始响应信息 ========")
	fmt.Println(GetRawResponseInfo(data.RawResponse))

	fmt.Println("\n------ 响应头 ------")
	fmt.Println(GetRawResponseHeaders(data.RawResponse))

	fmt.Println("\n------ 响应体 ------")
	// 格式化JSON输出
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data.RawResponse.Body(), "", "  ")
	if err != nil {
		fmt.Println(GetRawResponseBody(data.RawResponse))
	} else {
		fmt.Println(prettyJSON.String())
	}
}

// PrintLatLngDetailRawResponse 打印经纬度详细位置原始响应
func PrintLatLngDetailRawResponse(data *models.LatLngDetail) {
	if data == nil || data.RawResponse == nil {
		fmt.Println("无原始响应数据")
		return
	}

	fmt.Println("\n======== 原始响应信息 ========")
	fmt.Println(GetRawResponseInfo(data.RawResponse))

	fmt.Println("\n------ 响应头 ------")
	fmt.Println(GetRawResponseHeaders(data.RawResponse))

	fmt.Println("\n------ 响应体 ------")
	// 格式化JSON输出
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data.RawResponse.Body(), "", "  ")
	if err != nil {
		fmt.Println(GetRawResponseBody(data.RawResponse))
	} else {
		fmt.Println(prettyJSON.String())
	}
}
