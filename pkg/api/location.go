package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/hsn0918/iplocate/pkg/models"
	"github.com/hsn0918/iplocate/pkg/utils"
)

// LocationService 提供位置查询相关的服务
type LocationService struct {
	client *resty.Client
	debug  bool
	logger *utils.HTTPLogger
}

// NewLocationService 创建一个新的位置服务实例
func NewLocationService() *LocationService {
	client := resty.New().
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(500 * time.Millisecond).
		SetRetryMaxWaitTime(3 * time.Second).
		EnableTrace()

	return &LocationService{
		client: client,
		debug:  false,
		logger: utils.NewHTTPLogger(false),
	}
}

// SetDebug 设置是否启用调试模式
func (s *LocationService) SetDebug(debug bool) {
	s.debug = debug
	s.client.SetDebug(debug)
	s.logger = utils.NewHTTPLogger(debug)
}

// GetLocationByIP 通过IP地址获取位置信息
func (s *LocationService) GetLocationByIP(ip string) (*models.IPLocationData, error) {
	// 构造请求URL
	url := "https://apimobile.meituan.com/locate/v2/ip/loc"

	// 记录请求信息
	s.logger.LogRequest("GET", url, map[string][]string{
		"User-Agent": {s.client.Header.Get("User-Agent")},
	})

	// 发起HTTP请求
	resp, err := s.client.R().
		SetQueryParams(map[string]string{
			"rgeo": "true",
			"ip":   ip,
		}).
		Get(url)

	// 处理请求错误
	if err != nil {
		s.logger.LogError(err, "GET", url)
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	// 记录响应信息
	s.logger.LogResponse(resp.StatusCode(), resp.Time(), len(resp.Body()))

	// 检查HTTP状态码
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("接收到非成功状态码: %d", resp.StatusCode())
	}

	// 解析响应体
	var response models.IPResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	// 保存原始响应
	response.RawResponse = resp
	response.Data.RawResponse = resp

	// 如果响应体解析失败，Data 会为空
	if response.Data.IP == "" {
		return nil, fmt.Errorf("解析响应失败或返回数据为空")
	}

	return &response.Data, nil
}

// GetDetailByLatLngWithTags 通过经纬度获取tag从0到10的所有详细位置信息
func (s *LocationService) GetDetailByLatLngWithTags(lat, lng float64) ([]models.TaggedLatLngDetail, error) {
	var results []models.TaggedLatLngDetail
	baseURL := fmt.Sprintf("https://apimobile.meituan.com/group/v1/city/latlng/%f,%f", lat, lng)

	// 遍历tag从0到10
	for tag := 0; tag <= 10; tag++ {
		// 记录请求信息
		url := baseURL
		s.logger.LogRequest("GET", url, map[string][]string{
			"User-Agent": {s.client.Header.Get("User-Agent")},
		})

		// 发起HTTP请求
		resp, err := s.client.R().
			SetQueryParam("tag", fmt.Sprintf("%d", tag)).
			Get(url)

		// 处理请求错误
		if err != nil {
			s.logger.LogError(err, "GET", url)
			break // 继续下一个tag，不中断整个过程
		}

		// 记录响应信息
		s.logger.LogResponse(resp.StatusCode(), resp.Time(), len(resp.Body()))

		// 检查HTTP状态码
		if resp.StatusCode() != http.StatusOK {
			break // 继续下一个tag，不中断整个过程
		}

		// 解析响应体
		var response models.LatLngResponse
		err = json.Unmarshal(resp.Body(), &response)
		if err != nil {
			break // 继续下一个tag，不中断整个过程
		}

		// 保存原始响应
		response.RawResponse = resp
		response.Data.RawResponse = resp

		// 如果响应体解析成功且Data不为空
		if response.Data.City != "" {
			taggedDetail := models.TaggedLatLngDetail{
				Tag:    tag,
				Detail: response.Data,
			}
			results = append(results, taggedDetail)
		}
	}

	// 如果没有获取到任何有效结果
	if len(results) == 0 {
		return nil, fmt.Errorf("未能获取到任何有效的位置信息")
	}

	return results, nil
}
