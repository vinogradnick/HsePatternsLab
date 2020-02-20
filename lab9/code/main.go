package main

import "hse.pattern/lab6/code/car"

func main() {

	proxy := car.NewProxy(car.NewCarService())
	proxy.SpeedUp()
	proxy.SpeedDown()
	proxy.GearUp()
	proxy.GearUp()

}
