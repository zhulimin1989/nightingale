package astats

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "n9e"
	subsystem = "alert"
)

type Stats struct {

	// Counter	channel	告警通知发送总数（按通知渠道分类，如Email、Slack等）
	AlertNotifyTotal *prometheus.CounterVec

	// Counter	channel	告警通知发送失败次数（按通知渠道分类）
	AlertNotifyErrorTotal *prometheus.CounterVec

	// Counter	cluster, type, busi_group	产生的告警事件总数（按集群、类型、业务组分类）
	CounterAlertsTotal *prometheus.CounterVec

	// Gauge - 内存中告警事件队列的当前大小
	GaugeAlertQueueSize prometheus.Gauge

	// Counter - 规则评估总数（所有规则评估的总次数）
	CounterRuleEval *prometheus.CounterVec

	// Counter	datasource	数据查询失败次数（按数据源分类）
	CounterQueryDataErrorTotal *prometheus.CounterVec

	// Counter	datasource, rule_id	数据查询总次数（按数据源、规则ID分类）
	CounterQueryDataTotal *prometheus.CounterVec

	// Counter	rule_id, datasource_id, ref, type 变量填充查询次数（按规则ID、数据源ID、引用、类型分类）
	CounterVarFillingQuery *prometheus.CounterVec

	//Counter	datasource	记录评估总次数（按数据源分类）
	CounterRecordEval *prometheus.CounterVec

	// Counter	datasource	记录评估失败次数（按数据源分类）
	CounterRecordEvalErrorTotal *prometheus.CounterVec

	// Counter	group, rule_id, mute_rule_id, datasource_id	告警静默触发次数（按分组、规则ID、静默规则ID、数据源ID分类）
	CounterMuteTotal *prometheus.CounterVec

	// Counter	datasource, stage, busi_group, rule_id	规则评估失败次数（按数据源、阶段、业务组、规则ID分类）
	CounterRuleEvalErrorTotal *prometheus.CounterVec

	// Counter	-	心跳检测失败次数
	CounterHeartbeatErrorTotal *prometheus.CounterVec

	// Counter	group	子事件处理总数（按分组分类）
	CounterSubEventTotal *prometheus.CounterVec

	// Gauge	rule_id, datasource_id, ref	评估查询返回的时间序列数量（按规则ID、数据源ID、查询引用分类）
	GaugeQuerySeriesCount *prometheus.GaugeVec

	// Gauge	通知记录队列的当前大小
	GaugeNotifyRecordQueueSize prometheus.Gauge
}

func NewSyncStats() *Stats {
	CounterRuleEval := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "rule_eval_total",
		Help:      "Number of rule eval.",
	}, []string{})

	CounterRuleEvalErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "rule_eval_error_total",
		Help:      "Number of rule eval error.",
	}, []string{"datasource", "stage", "busi_group", "rule_id"})

	CounterQueryDataErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "query_data_error_total",
		Help:      "Number of rule eval query data error.",
	}, []string{"datasource"})

	CounterQueryDataTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "query_data_total",
		Help:      "Number of rule eval query data.",
	}, []string{"datasource", "rule_id"})

	CounterRecordEval := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "record_eval_total",
		Help:      "Number of record eval.",
	}, []string{"datasource"})

	CounterRecordEvalErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "record_eval_error_total",
		Help:      "Number of record eval error.",
	}, []string{"datasource"})

	AlertNotifyTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alert_notify_total",
		Help:      "Number of send msg.",
	}, []string{"channel"})

	AlertNotifyErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alert_notify_error_total",
		Help:      "Number of send msg.",
	}, []string{"channel"})

	// 产生的告警总量
	CounterAlertsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alerts_total",
		Help:      "Total number alert events.",
	}, []string{"cluster", "type", "busi_group"})

	// 内存中的告警事件队列的长度
	GaugeAlertQueueSize := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alert_queue_size",
		Help:      "The size of alert queue.",
	})

	CounterMuteTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "mute_total",
		Help:      "Number of mute.",
	}, []string{"group", "rule_id", "mute_rule_id", "datasource_id"})

	CounterSubEventTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "sub_event_total",
		Help:      "Number of sub event.",
	}, []string{"group"})

	CounterHeartbeatErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "heartbeat_error_count",
		Help:      "Number of heartbeat error.",
	}, []string{})

	GaugeQuerySeriesCount := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "eval_query_series_count",
		Help:      "Number of series retrieved from data source after query.",
	}, []string{"rule_id", "datasource_id", "ref"})
	// 通知记录队列的长度
	GaugeNotifyRecordQueueSize := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "notify_record_queue_size",
		Help:      "The size of notify record queue.",
	})

	CounterVarFillingQuery := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "var_filling_query_total",
		Help:      "Number of var filling query.",
	}, []string{"rule_id", "datasource_id", "ref", "typ"})

	prometheus.MustRegister(
		CounterAlertsTotal,
		GaugeAlertQueueSize,
		AlertNotifyTotal,
		AlertNotifyErrorTotal,
		CounterRuleEval,
		CounterQueryDataTotal,
		CounterQueryDataErrorTotal,
		CounterRecordEval,
		CounterRecordEvalErrorTotal,
		CounterMuteTotal,
		CounterRuleEvalErrorTotal,
		CounterHeartbeatErrorTotal,
		CounterSubEventTotal,
		GaugeQuerySeriesCount,
		GaugeNotifyRecordQueueSize,
		CounterVarFillingQuery,
	)

	return &Stats{
		CounterAlertsTotal:          CounterAlertsTotal,
		GaugeAlertQueueSize:         GaugeAlertQueueSize,
		AlertNotifyTotal:            AlertNotifyTotal,
		AlertNotifyErrorTotal:       AlertNotifyErrorTotal,
		CounterRuleEval:             CounterRuleEval,
		CounterQueryDataTotal:       CounterQueryDataTotal,
		CounterQueryDataErrorTotal:  CounterQueryDataErrorTotal,
		CounterRecordEval:           CounterRecordEval,
		CounterRecordEvalErrorTotal: CounterRecordEvalErrorTotal,
		CounterMuteTotal:            CounterMuteTotal,
		CounterRuleEvalErrorTotal:   CounterRuleEvalErrorTotal,
		CounterHeartbeatErrorTotal:  CounterHeartbeatErrorTotal,
		CounterSubEventTotal:        CounterSubEventTotal,
		GaugeQuerySeriesCount:       GaugeQuerySeriesCount,
		GaugeNotifyRecordQueueSize:  GaugeNotifyRecordQueueSize,
		CounterVarFillingQuery:      CounterVarFillingQuery,
	}
}
