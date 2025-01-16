package calendar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"Almanac-Push/internal/types"
)

type Client struct {
	key  string
	http *http.Client
}

func NewClient(key string) *Client {
	return &Client{
		key:  key,
		http: &http.Client{},
	}
}

// GetAlmanac 获取黄历数据
func (c *Client) GetAlmanac() (*types.AlmanacData, error) {
	uri := "https://cn.apihz.cn/api/time/getday.php?id=10002074&key=7948bc89efc2c7f1ea23cd3fa6ff42cb"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求日历API失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	fmt.Printf("API响应: %s\n", string(body))

	var result struct {
		Code       int    `json:"code"`
		Ynian      string `json:"ynian"`      // 公历年
		Yyue       string `json:"yyue"`       // 公历月
		Yri        string `json:"yri"`        // 公历日
		Nyue       string `json:"nyue"`       // 农历月
		Nri        string `json:"nri"`        // 农历日
		Ganzhinian string `json:"ganzhinian"` // 干支年
		Yi         string `json:"yi"`         // 宜
		Ji         string `json:"ji"`         // 忌
		Shengxiao  string `json:"shengxiao"`  // 生肖
		Jieqi      string `json:"jieqi"`      // 节气
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w, 响应内容: %s", err, string(body))
	}

	if result.Code != 200 {
		return nil, fmt.Errorf("API返回错误码: %d, 响应内容: %s", result.Code, string(body))
	}

	return &types.AlmanacData{
		Date:       fmt.Sprintf("%s-%s-%s", result.Ynian, result.Yyue, result.Yri),
		LunarYear:  result.Ynian,
		LunarMonth: result.Nyue,
		LunarDay:   result.Nri,
		GanZhi:     result.Ganzhinian,
		Zodiac:     result.Shengxiao,
		SolarTerm:  result.Jieqi,
		Suit:       strings.Split(result.Yi, "|"),
		Avoid:      strings.Split(result.Ji, "|"),
	}, nil
}
