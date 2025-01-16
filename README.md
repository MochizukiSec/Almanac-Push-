# Almanac-Push

一个简单的黄历推送服务，支持每日定时推送黄历信息到多个渠道。

## 功能特点

- 自动获取当日黄历信息（农历日期、干支、节气等）
- 支持多种推送渠道
- Bark（iOS 推送）
- 企业微信机器人
- 可配置推送时间
- 支持自定义推送内容格式

## 快速开始

### 1. 获取代码

```bash
git clone https://github.com/yourusername/Almanac-Push.git
cd Almanac-Push
```

### 2. 配置

复制 `config.json.example` 到 `config.json` 并修改相关配置：

```json
{
"calendar_api": {
  "key": "YOUR_CALENDAR_API_KEY"
},
"push": {
  "bark_key": "YOUR_BARK_KEY",
  "bark_sound": "minuet.caf",
  "bark_group": "Almanac",
  "wecom_bot_url": "YOUR_WECOM_BOT_WEBHOOK_URL",
  "enabled_channels": ["bark", "wecom_bot"],
  "archive": true
},
"push_time": {
  "hour": 7,
  "minute": 30
}
}
```

### 3. 编译运行

```bash
go build -o almanac-push cmd/almanac/main.go
./almanac-push
```

## 配置说明

### 推送渠道配置

#### Bark
- `bark_key`: Bark 推送的 Key
- `bark_sound`: 推送提示音
- `bark_group`: 消息分组名称

#### 企业微信机器人
- `wecom_bot_url`: 企业微信机器人的 Webhook URL

### 推送时间配置
- `hour`: 推送小时（24小时制）
- `minute`: 推送分钟

### 其他配置
- `enabled_channels`: 启用的推送渠道列表
- `archive`: 是否保存推送记录

## 推送内容格式

每日推送的内容包括：
- 公历日期
- 农历日期
- 干支纪年
- 生肖
- 节气（如果有）
- 宜忌事项

示例：
```
📅 公历：2025-01-16
🏮 农历：乙巳年八月初十
⭐ 干支：乙巳年
🐾 生肖：蛇
🌞 节气：小寒
✅ 宜：祭祀、沐浴、理发、冠笄
❌ 忌：开光、嫁娶、入宅、动土
```

## 开发计划

- [ ] 添加更多推送渠道
- [ ] 支持自定义消息模板
- [ ] 添加 Web 管理界面
- [ ] 支持多时区配置

## 贡献指南

欢迎提交 Issue 和 Pull Request！在提交之前，请确保：

1. 代码遵循 Go 语言规范
2. 添加必要的测试用例
3. 更新相关文档

## 许可证

本项目采用 MIT License 开源许可证。

## 致谢

- [黄历 API](https://cn.apihz.cn/)
- [Bark](https://github.com/Finb/Bark)

## 如何贡献

1. Fork 本仓库
2. 创建新的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request
