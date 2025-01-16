package types

type AlmanacData struct {
	Date       string   `json:"date"`       // 公历日期
	LunarYear  string   `json:"lunarYear"`  // 农历年
	LunarMonth string   `json:"lunarMonth"` // 农历月
	LunarDay   string   `json:"lunarDay"`   // 农历日
	GanZhi     string   `json:"ganZhi"`     // 天干地支
	Zodiac     string   `json:"zodiac"`     // 生肖
	SolarTerm  string   `json:"term"`       // 节气
	Suit       []string `json:"suit"`       // 宜
	Avoid      []string `json:"avoid"`      // 忌
}

type Config struct {
	CalendarAPI struct {
		Key string `json:"key"`
	} `json:"calendar_api"`
	Push struct {
		// Bark 配置
		BarkKey   string `json:"bark_key"`
		BarkSound string `json:"bark_sound"`
		BarkGroup string `json:"bark_group"`

		// 企业微信机器人配置
		WecomBotURL string `json:"wecom_bot_url"` // 机器人 webhook URL

		EnabledChannels []string `json:"enabled_channels"`
		Archive         bool     `json:"archive"`
	} `json:"push"`
	PushTime struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"push_time"`
}
