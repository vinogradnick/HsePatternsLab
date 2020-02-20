package car

//блока цилиндров с картером,
//
//- головки блока цилиндров,
//
//- поддона картера двигателя,
//
//- поршней с кольцами и пальцами,
//
//- шатунов,
//
//- коленчатого вала,
//
//- маховика.
import (
	"container/list"
)

type Engine struct {
	Container *list.List
}

func (e Engine) Print(num int) {
	for e := e.Container.Front(); e != nil; e = e.Next() {
		for i := 0; i < num; i++ {
			print("-")
		}
		println(TypeFlex(e.Value))
	}
}

type CylinderBlock struct {
}
type Sump struct {
}
type Crankshaft struct {
}
type Rod struct {
}

func (e Engine) AddPart(part interface{}) {
	e.Container.PushBack(part)
}
func (e Engine) RemovePart(part *interface{}) {
	el := list.Element{Value: part}
	e.Container.Remove(&el)
}

type EnginePart struct {
}
