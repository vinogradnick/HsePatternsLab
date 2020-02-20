package car

import "fmt"

type Proxy struct {
	service Service
}

func NewProxy(service Service) *Proxy {
	return &Proxy{service: service}
}

func (p Proxy) SpeedUp() {

	p.service.SpeedUp()
}

func (p Proxy) SpeedDown() {
	p.service.SpeedDown()
}

func (p Proxy) GearUp() {
	p.service.GearUp()

}

func (p Proxy) GearDown() {
	p.service.GearDown()

}

type Service interface {
	SpeedUp()
	SpeedDown()
	GearUp()
	GearDown()
}
type CarService struct {
	engine       Engine
	transmission Transmission
}

func (c CarService) SpeedUp() {
	c.engine.fuel--
	if c.engine.fuel >= 0 {
		c.engine.speed++
	}
	fmt.Println("engine speed:=> ", c.engine.speed)

}

func (c CarService) SpeedDown() {
	if c.engine.speed > 0 {
		c.engine.speed--
	}
	fmt.Println("engine speed:=> ", c.engine.speed)
}

func (c CarService) GearUp() {
	c.transmission.gear++
	fmt.Println("transmission gear:=> ", c.transmission.gear)
}

func (c CarService) GearDown() {
	c.transmission.gear--
	fmt.Println("transmission gear:=> ", c.transmission.gear)
}

func NewCarService() *CarService {
	return &CarService{engine: Engine{fuel: 100}, transmission: Transmission{gear: 0}}
}

type Engine struct {
	fuel  int
	speed int
}
type Transmission struct {
	gear int
}
