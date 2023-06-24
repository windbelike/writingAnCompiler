package object

import (
	"fmt"
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}
	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}

	var someInt64 int64 = -100
	fmt.Printf("in64:%d to unit64:%d \n", someInt64, uint64(someInt64))

    // Struct values are comparable if all their fields are comparable.
    // Two struct values are equal if their corresponding non-blank fields are equal.
	name1 := &String{Value: "name"}
	monkey := &String{Value: "Monkey"}

	pairs := map[HashKey]Object{}
	pairs[name1.HashKey()] = monkey
	fmt.Printf("pairs[name1.HashKey()]=%+v\n", pairs[name1.HashKey()])
	// => pairs[name1.HashKey()]=&{Value:Monkey}

	name2 := &String{Value: "name"}
	fmt.Printf("pairs[name2.HashKey()]=%+v\n", pairs[name2.HashKey()])
	// => pairs[name2.HashKey()]=&{Value:Monkey}
	fmt.Printf("(name1 == name2)=%t\n", name1 == name2)
	// => (name1 == name2)=false
	fmt.Printf("(name1.HashKey() == name2.HashKey())=%t\n",
		name1.HashKey() == name2.HashKey())
	// => (name1.HashKey() == name2.HashKey())=true

}
