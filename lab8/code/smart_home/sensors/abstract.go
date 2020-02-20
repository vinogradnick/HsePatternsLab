package sensors

import (
	"math/rand"
	"time"
)

type AbstractSensor interface {
	CreateMetric()
	SendMetric() *struct{}
}

func RandFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

type MetricStringify interface {
	getString() string
}
