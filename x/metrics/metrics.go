package metrics

import (
	"github.com/hibiken/asynq"
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "asynq"

type QueueMetricsCollector struct {
	inspector *asynq.Inspector
}

func (qmc *QueueMetricsCollector) collectQueueInfo() ([]*asynq.QueueInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	tasksQueuedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "tasks_enqueued_total"),
		"Number of tasks enqueued; broken down by queue and state.",
		[]string{"queue", "state"}, nil,
	)

	queueSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "queue_size"),
		"Number of tasks in a queue",
		[]string{"queue"}, nil,
	)

	queueLatencyDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "queue_latency_seconds"),
		"Number of seconds the oldest pending task is waiting in pending state to be processed.",
		[]string{"queue"}, nil,
	)

	queueMemUsgDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "queue_memory_usage_approx_bytes"),
		"Number of memory used by a given queue (approximated number by sampling).",
		[]string{"queue"}, nil,
	)

	tasksProcessedTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "tasks_processed_total"),
		"Number of tasks processed (both succeeded and failed); broken down by queue",
		[]string{"queue"}, nil,
	)

	tasksFailedTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "tasks_failed_total"),
		"Number of tasks failed; broken down by queue",
		[]string{"queue"}, nil,
	)

	pausedQueues = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "queue_paused_total"),
		"Number of queues paused",
		[]string{"queue"}, nil,
	)
)

func (qmc *QueueMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}

func (qmc *QueueMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}

func NewQueueMetricsCollector(inspector *asynq.Inspector) *QueueMetricsCollector {
	_ = "STUB: not implemented"
	return nil
}
