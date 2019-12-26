package main

import "fmt"

func main() {
	var i int
	var m map[string]string
	var a []int

	passIntByValue(i)
	fmt.Printf("passIntByValue: %d\n", i)
	passIntByPointer(&i)
	fmt.Printf("passIntByPointer: %d\n", i)
	passIntByInterface(&i)
	fmt.Printf("passIntByInterface: %d\n", i)
	passIntByInterfaceValue(i)
	fmt.Printf("passIntByInterfaceValue: %d\n", i)

	fmt.Printf("before processing the map %t\n", m != nil)
	passMapByValue(m)
	fmt.Printf("passMapByValue %t\n", m != nil)

	fmt.Printf("before processing the array %t\n", a != nil)
	passArrayByValue(a)
	fmt.Printf("passArrayByValue %t\n", a != nil)
	passArrayByPointer(&a)
	fmt.Printf("passArrayByPointer %t\n", a != nil)

	typeCast(&i)

	// it prints

	// passIntByValue: 0
	// passIntByPointer: 88
	// passIntByInterface: 77
	// passIntByInterfaceValue (inside): 66
	// passIntByInterfaceValue: 77
	// before processing the map false
	// passMapByPointer (inside): true
	// passMapByValue false
	// before processing the array false
	// passArrayByValue (inside): true
	// passArrayByValue false
	// passArrayByPointer true
	// typeCast *int: 77
}

func passIntByValue(v int) {
	v = 999
}

func passIntByPointer(v *int) {
	*v = 88 // equivalent to v = 88
}

func passIntByInterface(v interface{}) {
	temp := v.(*int)
	*temp = 77
}

func passIntByInterfaceValue(v interface{}) {
	temp := v.(int)
	temp = 66
	fmt.Printf("passIntByInterfaceValue (inside): %d\n", temp)
}

func passMapByValue(v map[string]string) {
	v = make(map[string]string)
	fmt.Printf("passMapByPointer (inside): %t\n", v != nil)
}

func passArrayByValue(v []int) {
	v = make([]int, 10)
	fmt.Printf("passArrayByValue (inside): %t\n", v != nil)
}

func passArrayByPointer(v *[]int) {
	*v = make([]int, 10)
}

func typeCast(v interface{}) {
	switch c := v.(type) {
	case int:
		fmt.Printf("typeCast int: %d\n", c)
	case string:
		fmt.Printf("typeCase string: %s\n", c)
	case *int:
		fmt.Printf("typeCast *int: %d\n", *c)
	default:
		fmt.Println("typeCast does not support this type")
	}
}
