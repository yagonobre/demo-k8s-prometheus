package delay

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	delayDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "delay_durations_histogram_seconds",
		Help: "The latency distributions by func",
	}, []string{"func"})
)

func init() {
	prometheus.MustRegister(delayDuration)
}

func Delay1() {
	now := time.Now()

	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)

	delayDuration.WithLabelValues("deploy1").Observe(time.Since(now).Seconds())
}

func Delay2() {
	now := time.Now()

	r := rand.Intn(1000)
	time.Sleep(time.Duration(r) * time.Millisecond)

	delayDuration.WithLabelValues("deploy2").Observe(time.Since(now).Seconds())

}

func Delay3() {
	now := time.Now()

	r := rand.Intn(1500)
	time.Sleep(time.Duration(r) * time.Millisecond)

	delayDuration.WithLabelValues("deploy3").Observe(time.Since(now).Seconds())
}
