package sensors

import (
	"fmt"
)

type ElectricSensor struct {
}

func (e ElectricSensor) CreateMetric() string {
	return EletricMetric{
		voltage:     RandFloat(110, 220),
		amperage:    RandFloat(10, 100),
		power:       RandFloat(1, 15),
		consumption: RandFloat(0, 100),
	}.getString()
}

type EletricMetric struct {
	voltage     float64
	amperage    float64
	power       float64
	consumption float64
}

func (e EletricMetric) getString() string {
	return fmt.Sprintf("%vV%vA%vW%vW", e.voltage, e.amperage, e.power, e.consumption)
}
