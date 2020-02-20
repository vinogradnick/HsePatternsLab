package sensors

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PressureSensor struct {
}

func (l PressureSensor) CreateMetric() string {
	r, err := rand.Int(rand.Reader, big.NewInt(80))
	if err != nil {
		panic(err)
	}
	return PressureMetric{pressure: r.Int64() % 10}.getString()
}

type PressureMetric struct {
	pressure int64
}

func (e PressureMetric) getString() string {
	return fmt.Sprintf("%vbar", e.pressure)
}
