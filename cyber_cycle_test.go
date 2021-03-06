package go_ehlers_indicators

import (
	"fmt"
	"github.com/MathisWellmann/go_timeseries_generator"
	"testing"
)

func TestCyberCycle(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cc := CyberCycle(vals, 16)
	for i := 0; i < len(cc); i++ {
		fmt.Printf("cc[%d]: %f\n", i, cc[i])
	}
}

func TestCyberCycleGraph(t *testing.T) {
	vals := go_timeseries_generator.GaussianProcess(1024)
	cc := CyberCycle(vals, 16)

	filename := fmt.Sprintf("./img/cyber_cycle.png")
	err := Plt(cc, filename)
	if err != nil {
		t.Error(err)
	}
}
