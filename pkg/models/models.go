package models

import "github.com/go-resty/resty/v2"

// IPResponse 定义IP定位API的响应结构
type IPResponse struct {
	Data IPLocationData `json:"data"`
	// 保存原始响应
	RawResponse *resty.Response `json:"-"`
}

// IPLocationData 定义IP位置数据结构
type IPLocationData struct {
	Lng       float64   `json:"lng"`
	FromWhere string    `json:"fromwhere"`
	IP        string    `json:"ip"`
	Rgeo      RegionGeo `json:"rgeo"`
	Lat       float64   `json:"lat"`
	// 保存原始响应
	RawResponse *resty.Response `json:"-"`
}

// RegionGeo 定义区域地理信息结构
type RegionGeo struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	Adcode   string `json:"adcode"`
	City     string `json:"city"`
	District string `json:"district"`
}

// LatLngResponse 定义经纬度详细信息API的响应结构
type LatLngResponse struct {
	Data LatLngDetail `json:"data"`
	// 保存原始响应
	RawResponse *resty.Response `json:"-"`
}

// LatLngDetail 定义经纬度详细位置数据结构
type LatLngDetail struct {
	Detail       string  `json:"detail"`
	ParentArea   int     `json:"parentArea"`
	CityPinyin   string  `json:"cityPinyin"`
	Lng          float64 `json:"lng"`
	IsForeign    bool    `json:"isForeign"`
	DpCityId     int     `json:"dpCityId"`
	Country      string  `json:"country"`
	IsOpen       bool    `json:"isOpen"`
	City         string  `json:"city"`
	Id           int     `json:"id"`
	OpenCityName string  `json:"openCityName"`
	OriginCityID int     `json:"originCityID"`
	Area         int     `json:"area"`
	AreaName     string  `json:"areaName"`
	Province     string  `json:"province"`
	District     string  `json:"district"`
	Lat          float64 `json:"lat"`
	// 保存原始响应
	RawResponse *resty.Response `json:"-"`
}

// TaggedLatLngDetail 定义带有标签的经纬度详细位置数据结构
type TaggedLatLngDetail struct {
	Tag    int          `json:"tag"`
	Detail LatLngDetail `json:"detail"`
}
