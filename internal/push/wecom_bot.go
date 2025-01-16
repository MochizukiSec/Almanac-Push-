package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"Almanac-Push/internal/types"
)

type WecomBotPusher struct {
	webhookURL string
	http       *http.Client
}

func NewWecomBotPusher(cfg *types.Config) *WecomBotPusher {
	return &WecomBotPusher{
		webhookURL: cfg.Push.WecomBotURL,
		http:       &http.Client{},
	}
}

func (w *WecomBotPusher) Push(title, body string) error {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": fmt.Sprintf("%s\n\n%s", title, body),
		},
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("序列化消息失败: %w", err)
	}

	resp, err := w.http.Post(w.webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送消息失败: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("解析响应失败: %w", err)
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("企业微信机器人返回错误: %s", result.ErrMsg)
	}

	return nil
}
