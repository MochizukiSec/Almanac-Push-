package types

type Weather struct {
	City        string `json:"city"`        // 城市
	Temperature string `json:"temperature"` // 温度
	Weather     string `json:"weather"`     // 天气状况
	Wind        string `json:"wind"`        // 风向风力
	Humidity    string `json:"humidity"`    // 湿度
	AQI         string `json:"aqi"`         // 空气质量指数
	UpdateTime  string `json:"update_time"` // 更新时间
}

// 在 Config 结构体中添加天气配置
type WeatherConfig struct {
	City string `json:"city"` // 城市名称
	Key  string `json:"key"`  // API密钥
}
