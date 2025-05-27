package prometheus

import "github.com/prometheus/client_golang/prometheus"

var rpcResourceCounter prometheus.Counter

func MetricsInit() error {
	rpcResourceCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name:      "rpc_counter",
		Subsystem: "resource",
		Help:      "Number of rpc calls",
	})
	if err := prometheus.Register(rpcResourceCounter); err != nil {
		return err
	}
	return nil
}

func RpcResourceCounterInc() {
	rpcResourceCounter.Inc()
}
