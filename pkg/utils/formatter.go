package utils

import (
	"fmt"
	"strings"

	"github.com/hsn0918/iplocate/pkg/models"
)

// OutputLevel 定义输出级别
type OutputLevel int

const (
	// OutputLevelBasic 基本输出级别，只显示关键信息
	OutputLevelBasic OutputLevel = iota
	// OutputLevelNormal 正常输出级别，显示大部分信息
	OutputLevelNormal
	// OutputLevelVerbose 详细输出级别，显示所有信息
	OutputLevelVerbose
)

// 默认输出级别
var currentOutputLevel = OutputLevelNormal

// SetOutputLevel 设置输出级别
func SetOutputLevel(level OutputLevel) {
	currentOutputLevel = level
}

// GetOutputLevel 获取当前输出级别
func GetOutputLevel() OutputLevel {
	return currentOutputLevel
}

// PrintIPLocationInfo 打印IP位置信息
func PrintIPLocationInfo(data *models.IPLocationData) {
	fmt.Println("🌐 ======== IP地址信息 ======== 🌐")
	fmt.Printf("🔍 IP: %s\n", data.IP)
	fmt.Printf("🧭 经度: %.6f\n", data.Lng)
	fmt.Printf("🧭 纬度: %.6f\n", data.Lat)

	if currentOutputLevel >= OutputLevelNormal {
		fmt.Printf("📊 数据来源: %s\n", data.FromWhere)
	}

	fmt.Println("\n📍 ------ 基础地理位置信息 ------ 📍")
	fmt.Printf("🏳️ 国家: %s\n", data.Rgeo.Country)
	fmt.Printf("🏞️ 省份: %s\n", data.Rgeo.Province)
	fmt.Printf("🏙️ 城市: %s\n", data.Rgeo.City)
	fmt.Printf("🏡 区县: %s\n", data.Rgeo.District)

	if currentOutputLevel >= OutputLevelNormal {
		fmt.Printf("🔢 行政区划代码: %s\n", data.Rgeo.Adcode)
	}

	// 打印完整地址
	fullAddress := strings.TrimSpace(fmt.Sprintf("%s %s %s %s",
		data.Rgeo.Country,
		data.Rgeo.Province,
		data.Rgeo.City,
		data.Rgeo.District))
	fmt.Printf("\n📮 基础完整地址: %s\n", fullAddress)
}

// PrintLatLngDetailInfo 打印经纬度详细位置信息
func PrintLatLngDetailInfo(data *models.TaggedLatLngDetail) {

	fmt.Println()
	fmt.Printf("\n🗺️ ======== [%d]详细位置信息 ======== 🗺️\n", data.Tag)
	fmt.Printf("🧭 经度: %.6f\n", data.Detail.Lng)
	fmt.Printf("🧭 纬度: %.6f\n", data.Detail.Lat)
	fmt.Printf("🏳️ 国家: %s\n", data.Detail.Country)
	fmt.Printf("🏞️ 省份: %s\n", data.Detail.Province)
	fmt.Printf("🏙️ 城市: %s\n", data.Detail.City)

	if currentOutputLevel >= OutputLevelNormal {
		fmt.Printf("🔤 城市拼音: %s\n", data.Detail.CityPinyin)
	}

	fmt.Printf("🏡 区县: %s\n", data.Detail.District)
	fmt.Printf("📍 区域名称: %s\n", data.Detail.AreaName)
	fmt.Printf("📝 详细地址: %s\n", data.Detail.Detail)

	if currentOutputLevel >= OutputLevelVerbose {
		fmt.Printf("🔢 区域ID: %d\n", data.Detail.Area)
		fmt.Printf("🔢 父区域ID: %d\n", data.Detail.ParentArea)
		fmt.Printf("🔢 地区ID: %d\n", data.Detail.Id)
		fmt.Printf("🔢 点评城市ID: %d\n", data.Detail.DpCityId)
		fmt.Printf("🔢 原始城市ID: %d\n", data.Detail.OriginCityID)
		fmt.Printf("🏙️ 开放城市名称: %s\n", data.Detail.OpenCityName)
		fmt.Printf("✅ 是否开放: %t\n", data.Detail.IsOpen)
		fmt.Printf("🌏 是否国外: %t\n", data.Detail.IsForeign)
	}

	// 打印完整详细地址
	fullDetailAddress := strings.TrimSpace(fmt.Sprintf("%s %s %s %s %s %s",
		data.Detail.Country,
		data.Detail.Province,
		data.Detail.City,
		data.Detail.District,
		data.Detail.AreaName,
		data.Detail.Detail))
	fmt.Printf("\n📮 完整详细地址: %s\n", fullDetailAddress)
}

// PrintRawResponse 打印原始响应数据
func PrintRawResponse(rawResponse string) {
	if rawResponse != "" && currentOutputLevel >= OutputLevelVerbose {
		fmt.Println("\n📡 ======== 原始响应数据 ======== 📡")
		fmt.Println(rawResponse)
		fmt.Println("📡 ========================== 📡")
	}
}
