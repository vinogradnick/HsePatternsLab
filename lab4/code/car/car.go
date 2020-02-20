package car

import (
	"container/list"
	"reflect"
)

func TypeFlex(v interface{}) string {
	return reflect.TypeOf(v).String()
}

type Car struct {
	Container *list.List
}

func (e Car) Print(num int) {
	for e := e.Container.Front(); e != nil; e = e.Next() {

		switch str := e.Value.(type) {
		case CartPart:
			println(TypeFlex(e.Value))


			str.Print(num + 1)
			break
		}
	}
}
func (e Car) AddPart(part interface{}) {

	e.Container.PushBack(part)
	println(e.Container.Len())
}
func (e Car) RemovePart(part *interface{}) {
	el := list.Element{Value: part}
	e.Container.Remove(&el)
}

type CartPart interface {
	Print(num int)
}
