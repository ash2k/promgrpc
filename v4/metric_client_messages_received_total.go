package promgrpc

import (
	"context"

	"github.com/piotrkowalczuk/promgrpc/v4/internal/useragent"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/stats"
)

func NewClientMessagesReceivedTotalCounterVec(opts ...CollectorOption) *prometheus.CounterVec {
	labels := []string{
		labelIsFailFast,
		labelMethod,
		labelService,
		labelClientUserAgent,
	}
	return newMessagesReceivedTotalCounterVec("client", labels, opts...)
}

type ClientMessagesReceivedTotalStatsHandler struct {
	baseStatsHandler
	uas useragent.Store
	vec *prometheus.CounterVec
}

// NewClientMessagesReceivedTotalStatsHandler ...
// The GaugeVec must have zero, one, two, three or four non-const non-curried labels.
// For those, the only allowed labelsFn names are "fail_fast", "handler", "service" and "user_agent".
func NewClientMessagesReceivedTotalStatsHandler(vec *prometheus.CounterVec, opts ...StatsHandlerOption) *ClientMessagesReceivedTotalStatsHandler {
	h := &ClientMessagesReceivedTotalStatsHandler{
		vec: vec,
	}
	h.baseStatsHandler = baseStatsHandler{
		collector: vec,
		options: statsHandlerOptions{
			handleRPCLabelFn: h.labels,
		},
	}
	h.applyOpts(opts...)

	return h
}

// HandleRPC implements stats Handler interface.
func (h *ClientMessagesReceivedTotalStatsHandler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	if _, ok := stat.(*stats.InPayload); ok {
		switch {
		case stat.IsClient():
			h.vec.WithLabelValues(h.options.handleRPCLabelFn(ctx, stat)...).Inc()
		}
	}
}

func (h *ClientMessagesReceivedTotalStatsHandler) labels(ctx context.Context, stat stats.RPCStats) []string {
	tag := ctx.Value(tagRPCKey).(rpcTagLabels)
	return []string{
		tag.isFailFast,
		tag.method,
		tag.service,
		h.uas.ClientSide(ctx, stat),
	}
}
