package aconf

import (
	"path"
)

// Alert 定义警报配置模型
// swagger:model
type Alert struct {
	// 是否禁用警报功能，默认关闭
	// example: false
	Disable bool `json:"disable"`

	// 引擎延迟时间（秒），默认30秒
	// example: 30
	EngineDelay int64 `json:"engine_delay"`

	// 心跳检测配置
	// required: true
	Heartbeat HeartbeatConfig `json:"heartbeat"`

	// 警报通知配置
	// required: true
	Alerting Alerting `json:"alerting"`
}

// HeartbeatConfig 定义心跳检测配置
// swagger:model
type HeartbeatConfig struct {
	// 检测节点IP地址
	// example: 127.0.0.1
	IP string `json:"ip"`

	// 检测间隔时间（毫秒），默认1000ms
	// example: 1000
	Interval int64 `json:"interval"`

	// 心跳检测端点地址
	// example: /api/heartbeat
	Endpoint string `json:"endpoint"`

	// 引擎实例名称
	// example: default-engine
	EngineName string `json:"engine_name"`
}

// Alerting 定义警报通知配置
// swagger:model
type Alerting struct {
	// 警报超时时间（毫秒），默认30000ms
	// example: 30000
	Timeout int64 `json:"timeout"`

	// 模板文件目录，默认会自动设置为configDir/template
	// example: /etc/alert/templates
	TemplatesDir string `json:"templates_dir"`

	// 通知并发数，默认10
	// example: 10
	NotifyConcurrency int `json:"notify_concurrency"`

	// 是否启用Webhook批量发送
	// example: true
	WebhookBatchSend bool `json:"webhook_batch_send"`
}

// SMTPConfig 定义SMTP邮件配置
// swagger:model
type SMTPConfig struct {
	// SMTP服务器地址
	// example: smtp.example.com
	Host string `json:"host"`

	// SMTP服务端口
	// example: 587
	Port int `json:"port"`

	// SMTP认证用户名
	// example: user@example.com
	User string `json:"user"`

	// SMTP认证密码
	Pass string `json:"pass"`

	// 发件人邮箱地址
	// example: no-reply@example.com
	From string `json:"from"`

	// 是否跳过证书验证
	// example: false
	InsecureSkipVerify bool `json:"insecure_skip_verify"`

	// 邮件发送批处理数量
	// example: 100
	Batch int `json:"batch"`
}

// CallPlugin 定义插件调用配置
// swagger:model
type CallPlugin struct {
	// 是否启用插件调用
	// example: true
	Enable bool `json:"enable"`

	// 插件执行路径
	// example: /plugins/
	PluginPath string `json:"plugin_path"`

	// 插件调用方标识
	// example: alert-manager
	Caller string `json:"caller"`
}

// RedisPub 定义Redis发布配置
// swagger:model
type RedisPub struct {
	// 是否启用Redis发布
	// example: true
	Enable bool `json:"enable"`

	// 发布频道前缀
	// example: alert:
	ChannelPrefix string `json:"channel_prefix"`

	// 发布频道密钥
	// example: alert_channel
	ChannelKey string `json:"channel_key"`
}

func (a *Alert) PreCheck(configDir string) {
	if a.Alerting.TemplatesDir == "" {
		a.Alerting.TemplatesDir = path.Join(configDir, "template")
	}

	if a.Alerting.NotifyConcurrency == 0 {
		a.Alerting.NotifyConcurrency = 10
	}

	if a.Heartbeat.Interval == 0 {
		a.Heartbeat.Interval = 1000
	}

	if a.EngineDelay == 0 {
		a.EngineDelay = 30
	}
}
