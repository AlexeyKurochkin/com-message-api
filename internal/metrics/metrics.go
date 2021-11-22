package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//HandledEventsTotal metrics for events sent to kafka
var HandledEventsTotal = promauto.NewGauge(prometheus.GaugeOpts{
	Subsystem: "com_message_api",
	Name:      "com_message_api_handled_events_total",
})

//EventsTotal metric for total events count
var EventsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: "com_message_api",
	Name:      "com_message_api_events_total",
}, []string{"event_type"})
