package main

import (
	"container/list"
	"hse.pattern/lab4/code/car"
)

func main() {
	engine := car.Engine{Container: list.New()}
	engine.AddPart(car.Crankshaft{})
	engine.AddPart(car.Sump{})
	engine.AddPart(car.Rod{})
	engine.AddPart(car.CylinderBlock{})

	susp := car.NewSuspenstion()
	susp.AddPart(car.Damping{})
	susp.AddPart(car.Rack{})

	carFlex := car.Car{Container: list.New()}
	carFlex.AddPart(engine)
	carFlex.AddPart(susp)
	carFlex.Print(2)

}
