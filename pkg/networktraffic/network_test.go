package networktraffic

import (
	"testing"
	"time"
)

func Test_A(t *testing.T) {
	ph := NewPrometheus("http://n1:31899", "ens192", time.Second)
	m, err := ph.GetNodeBandwidthMeasure("10.0.0.21:9100")
	println(m)
	println(err)
}
