package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"Almanac-Push/internal/calendar"
	"Almanac-Push/internal/push"
	"Almanac-Push/internal/types"

	"github.com/robfig/cron/v3"
)

var logger = log.New(os.Stdout, "[黄历推送] ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	// 加载配置
	cfg, err := loadConfig("config.json")
	if err != nil {
		logger.Fatalf("加载配置失败: %v", err)
	}

	// 初始化服务
	calendarClient := calendar.NewClient(cfg.CalendarAPI.Key)
	pushService := push.NewService(cfg)

	// 测试推送
	if err := sendDailyAlmanac(calendarClient, pushService); err != nil {
		logger.Printf("测试推送失败: %v", err)
	} else {
		logger.Println("测试推送成功")
		// 测试成功后退出
		return
	}

	// 创建定时任务
	c := cron.New()
	spec := fmt.Sprintf("%d %d * * *", cfg.PushTime.Minute, cfg.PushTime.Hour)

	_, err = c.AddFunc(spec, func() {
		if err := sendDailyAlmanac(calendarClient, pushService); err != nil {
			logger.Printf("发送黄历失败: %v", err)
		}
	})
	if err != nil {
		logger.Fatalf("创建定时任务失败: %v", err)
	}

	logger.Printf("服务已启动，将在每天 %02d:%02d 推送黄历",
		cfg.PushTime.Hour,
		cfg.PushTime.Minute)

	c.Start()
	select {}
}

func loadConfig(filename string) (*types.Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var cfg types.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &cfg, nil
}

func sendDailyAlmanac(calendarClient *calendar.Client, pushService *push.Service) error {
	// 获取黄历数据
	data, err := calendarClient.GetAlmanac()
	if err != nil {
		return fmt.Errorf("获取黄历数据失败: %w", err)
	}

	// 构建消息
	message := buildMessage(data)

	// 发送推送
	if err := pushService.Push("今日黄历", message); err != nil {
		return fmt.Errorf("发送推送失败: %w", err)
	}

	return nil
}

func buildMessage(data *types.AlmanacData) string {
	var sb strings.Builder

	// 日期部分
	sb.WriteString(fmt.Sprintf("📅 公历：%s\n", data.Date))
	sb.WriteString(fmt.Sprintf("🏮 农历：%s年%s月%s\n", data.LunarYear, data.LunarMonth, data.LunarDay))
	sb.WriteString(fmt.Sprintf("⭐ 干支：%s\n", data.GanZhi))
	sb.WriteString(fmt.Sprintf("🐾 生肖：%s\n", data.Zodiac))

	if data.SolarTerm != "" {
		sb.WriteString(fmt.Sprintf("\n🌞 节气：%s\n", data.SolarTerm))
	}

	// 宜忌部分
	sb.WriteString(fmt.Sprintf("\n✅ 宜：%s\n", strings.Join(data.Suit, "、")))
	sb.WriteString(fmt.Sprintf("❌ 忌：%s", strings.Join(data.Avoid, "、")))

	return sb.String()
}
