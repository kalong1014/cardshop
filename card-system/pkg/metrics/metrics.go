// 后端指标暴露
package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	orderCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cardshop_orders_total",
			Help: "总订单数",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(orderCounter)
}

// 在订单创建时记录指标
func RecordOrder(status string) {
	orderCounter.WithLabelValues(status).Inc()
}

// 暴露监控端点
func RegisterMetricsHandler(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}