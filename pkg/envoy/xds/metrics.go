// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package xds

import (
	"github.com/cilium/cilium/pkg/metrics"
	"github.com/cilium/cilium/pkg/metrics/metric"
)

const (
	subsystem       = "xds"
	typeURLLabel    = "type_url"
	statusLabel     = "status"
	statusACKValue  = "ack"
	statusNACKValue = "nack"
)

type Metrics interface {
	IncreaseNACK(string)
	IncreaseACK(string)
}

var _ Metrics = (*XDSMetrics)(nil)

type XDSMetrics struct {
	// ACK is the number of ACK responses from envoy.
	ACKCount metric.Vec[metric.Counter]

	// NACKCount is the number of NACK errors from envoy.
	NACKCount metric.Vec[metric.Counter]
}

func NewXDSMetric() *XDSMetrics {
	return &XDSMetrics{
		ACKCount: metric.NewCounterVec(metric.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: subsystem,
			Name:      "events_count",
			Help:      "The number of event responses from Envoy",
		}, []string{typeURLLabel, statusLabel}),
	}
}

func (x *XDSMetrics) IncreaseNACK(typeURL string) {
	x.NACKCount.WithLabelValues(typeURL, statusNACKValue).Inc()
}

func (x *XDSMetrics) IncreaseACK(typeURL string) {
	x.ACKCount.WithLabelValues(typeURL, statusACKValue).Inc()
}
