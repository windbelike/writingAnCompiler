package evaluator

import (
	"fmt"
	"testing"
)

type Animal interface {
	bark()
}

type Dog struct {
	name string
}

// func (d Dog) bark() {
// 	fmt.Println("wang wang wang")
// }

func (dd *Dog) bark() {
    fmt.Println("wang wang wang *")
}

type Cat struct {
	name string
}

func (c Cat) bark() {
	fmt.Println("miao miao miao")
}

// * and & in Golang
// So, this is how * and the & work in Golang. The * is used for declaring pointer variables as well as de-referencing pointer variables, and the & operator 
// is used for accessing the memory address of the variable.
func TestAnimal(t *testing.T) {
    // var animal Animal = &Dog {"some dog's name"}
    // fmt.Println(animal)
    // t.Log(fmt.Sprintln(animal))
    //
    // // type assertion
    // realDog := animal.(*Dog)
    // fmt.Println(realDog.name)
}
