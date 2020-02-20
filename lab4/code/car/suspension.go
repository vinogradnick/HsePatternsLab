package car

import "container/list"

//Демпфирующий
//Стабилизатор
//Стойка
//Рычаг
//Пружина
//Амортизатор

type Damping struct {
}
type Stabilizer struct {
}
type Rack struct {
}
type Lever struct {
}
type Spring struct {
}
type Absorber struct {
}
type Suspenstion struct {
	Container *list.List
}

func (e Suspenstion) Print(num int) {
	for e := e.Container.Front(); e != nil; e = e.Next() {
		for i := 0; i < num; i++ {
			print("-")
		}
		println(TypeFlex(e.Value))
	}
}

func (e Suspenstion) AddPart(part interface{}) {
	e.Container.PushBack(part)

}
func (e Suspenstion) RemovePart(part *interface{}) {
	el := list.Element{Value: part}
	e.Container.Remove(&el)
}

func NewSuspenstion() Suspenstion {
	return Suspenstion{Container: list.New()}
}
