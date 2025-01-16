package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"Almanac-Push/internal/types"
)

type BarkMessage struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Group     string `json:"group"`
	Sound     string `json:"sound"`
	IsArchive int    `json:"isArchive"`
}

type BarkPusher struct {
	key     string
	sound   string
	group   string
	archive bool
	http    *http.Client
}

func NewBarkPusher(cfg *types.Config) *BarkPusher {
	return &BarkPusher{
		key:     cfg.Push.BarkKey,
		sound:   cfg.Push.BarkSound,
		group:   cfg.Push.BarkGroup,
		archive: cfg.Push.Archive,
		http:    &http.Client{},
	}
}

func (b *BarkPusher) Push(title, body string) error {
	msg := BarkMessage{
		Title: title,
		Body:  body,
		Group: b.group,
		Sound: b.sound,
	}
	if b.archive {
		msg.IsArchive = 1
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("序列化消息失败: %w", err)
	}

	url := fmt.Sprintf("https://api.day.app/%s/", b.key)
	resp, err := b.http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送推送失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("推送服务返回错误状态码: %d", resp.StatusCode)
	}

	return nil
}
