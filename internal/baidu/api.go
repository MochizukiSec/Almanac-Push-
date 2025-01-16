package baidu

import (
	"net/http"
	"time"

	"Almanac-Push/internal/types"
)

type Client struct {
	key    string
	secret string
	http   *http.Client
}

func NewClient(key, secret string) *Client {
	return &Client{
		key:    key,
		secret: secret,
		http:   &http.Client{},
	}
}

func (c *Client) GetAlmanac() (*types.AlmanacData, error) {
	// 临时返回测试数据
	return &types.AlmanacData{
		Date:       time.Now().Format("2006-01-02"),
		LunarYear:  "甲辰",
		LunarMonth: "二月",
		LunarDay:   "初五",
		GanZhi:     "甲辰年 壬寅月 丙子日",
		Zodiac:     "龙",
		SolarTerm:  "春分",
		Suit:       []string{"祭祀", "开光", "出行"},
		Avoid:      []string{"动土", "开张"},
	}, nil
}
