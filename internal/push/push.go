package push

import (
	"fmt"

	"Almanac-Push/internal/types"
)

type Pusher interface {
	Push(title, body string) error
}

type Service struct {
	pushers map[string]Pusher
}

func NewService(cfg *types.Config) *Service {
	s := &Service{
		pushers: make(map[string]Pusher),
	}

	// 初始化已启用的推送渠道
	for _, channel := range cfg.Push.EnabledChannels {
		switch channel {
		case "bark":
			if cfg.Push.BarkKey != "" {
				s.pushers["bark"] = NewBarkPusher(cfg)
			}
		case "wecom_bot":
			if cfg.Push.WecomBotURL != "" {
				s.pushers["wecom_bot"] = NewWecomBotPusher(cfg)
			}
		}
	}

	return s
}

func (s *Service) Push(title, body string) error {
	var errs []error

	for name, pusher := range s.pushers {
		if err := pusher.Push(title, body); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("推送失败: %v", errs)
	}
	return nil
}
