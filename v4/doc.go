// Package promgrpc is an instrumentation package that allows capturing metrics of your gRPC based services, both the server and the client side.
// The main goal of version 4 was to make it modular without sacrificing the simplicity of use.
//
// It is still possible to integrate the package in just a few lines.
// However, if necessary, metrics can be added, removed or modified freely.
//
// Design
//
// The package does not introduce any new concepts to an already complicated environment.
// Instead, it focuses on providing implementations of interfaces exported by gRPC and Prometheus libraries.
//
// It causes no side effects nor has global state.
// Instead, it comes with handy one-liners to reduce integration overhead.
//
// The package achieved high modularity by using Inversion of Control.
// We can define three layers of abstraction, where each is configurable or if necessary replaceable.
//
// Collectors serve one purpose, storing metrics.
// These are types well known from Prometheus ecosystem, like counters, gauges, histograms or summaries.
// This package comes with a set of predefined functions that create a specific instances for each use case. For example:
//
//  func NewRequestsTotalCounterVec(Subsystem, ...CollectorOption) *prometheus.CounterVec
//
// Level higher consist of stats handlers. This layer is responsible for metrics collection.
// It is aware of a collector and knows how to use it to record event occurrences.
// Each implementation satisfies stats.Handler and prometheus.Collector interface and knows how to monitor a single dimension, e.g. a total number of received/sent requests:
//
//  func NewRequestsStatsHandler(Subsystem, *prometheus.GaugeVec, ...StatsHandlerOption) *RequestsStatsHandler {
//
// Above all, there is a coordinator.
// StatsHandler combines multiple stats handlers into a single instance.
package promgrpc