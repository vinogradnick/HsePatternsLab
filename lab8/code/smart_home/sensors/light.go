package sensors

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type LightSensor struct {
}

type LightMetric struct {
	lux int64
}

func (e LightMetric) getString() string {
	return fmt.Sprintf("%vlux", e.lux)
}

func (e LightMetric) CreateMetric() string {
	r, err := rand.Int(rand.Reader, big.NewInt(80))
	if err != nil {
		panic(err)
	}
	return LightMetric{lux: r.Int64() % 100}.getString()
}
