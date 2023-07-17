package comm

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

var metricsRequest prometheus.Gauge

// MetricAdd 指标计数
func MetricAdd() {
	// prometheus 指标
	metricsRequest.Add(1)
	metric := &dto.Metric{}
	metricsRequest.Write(metric)
	fmt.Println(proto.MarshalTextString(metric))
}

// MetricInit 初始化
func MetricInit(router *gin.Engine) {
	// prometheus client Create non-global registry.
	reg := prometheus.NewRegistry()
	metricsRequest = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "gin",
		Subsystem: "v1",
		Name:      "requests",
		Help:      "gin vi request gauge",
		ConstLabels: map[string]string{
			"grpc":   "method",
			"server": "UserGrowth",
		},
	})
	// Add go runtime metrics and process collectors.
	reg.MustRegister(
		collectors.NewGoCollector(),
		metricsRequest,
	)
	// Expose /metrics HTTP endpoint using the created custom registry.
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})))
}
