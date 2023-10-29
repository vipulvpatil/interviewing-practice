package main

import "fmt"

func main() {
	t := NewTest(1)
	fmt.Println("Should print \"simple: 1, ptr: {address}, value: 1\"")
	fmt.Println(t)

	t.ChangeInt(2)
	fmt.Println("Should print \"simple: 1, ptr: {same_address_as_before}, value: 1\"")
	fmt.Println(t)

	t.UpdateInt(2)
	fmt.Println("Should print \"simple: 2, ptr: {same_address_as_before}, value: 1\"")
	fmt.Println(t)

	n := 3

	t.ChangeIntPtr(&n)
	fmt.Println("Should print \"simple: 2, ptr: {same_address_as_before}, value: 1\"")
	fmt.Println(t)

	t.UpdateIntPtr(&n)
	fmt.Println("Should print \"simple: 2, ptr: {different_from_before}, value: 3\"")
	fmt.Println(t)
}

func NewTest(d int) Test {
	return Test{
		simpleVar:    d,
		pointerToVar: &d,
	}
}

type Test struct {
	simpleVar    int
	pointerToVar *int
}

func (t Test) ChangeInt(d int) {
	t.simpleVar = d
}

func (t *Test) UpdateInt(d int) {
	t.simpleVar = d
}

func (t Test) ChangeIntPtr(d *int) {
	t.pointerToVar = d
}

func (t *Test) UpdateIntPtr(d *int) {
	t.pointerToVar = d
}

func (t Test) String() string {
	return fmt.Sprintf("simple: %d, ptr: %v, value: %d", t.simpleVar, t.pointerToVar, *t.pointerToVar)
}
