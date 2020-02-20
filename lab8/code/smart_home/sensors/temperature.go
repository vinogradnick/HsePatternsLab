package sensors

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type TermalSensor struct {
}

type ThermalMetric struct {
	temperature int64
}

func (e ThermalMetric) getString() string {
	return fmt.Sprintf("%vC^", e.temperature)
}

func (e ThermalMetric) CreateMetric() string {
	r, err := rand.Int(rand.Reader, big.NewInt(80))
	if err != nil {
		panic(err)
	}
	return ThermalMetric{temperature: r.Int64() % 100}.getString()
}
