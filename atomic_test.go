package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomic1(t *testing.T) {
	var int = atomic.Int64{}
	int.Add(1)
	fmt.Println("add: ", int.Load())

	int.Store(5)
	fmt.Println("store: ", int.Load())

	old := int.Swap(6)
	fmt.Println("old: ", old)
	fmt.Println("new: ", int.Load())

	b := int
	b.Add(2)
	fmt.Println(b.Load())
	fmt.Println(int.Load())

}

type Person struct {
	Name string
	Age  int
}

func TestAtomic2(t *testing.T) {
	p := atomic.Pointer[Person]{}
	person := Person{Name: "tom", Age: 22}
	p.Store(&person)
	person2 := p.Load()
	person.Age = 23
	fmt.Println(person2.Name, person2.Age)
}

func TestAtomic3(t *testing.T) {
	person := Person{Name: "tom", Age: 22}
	v := &atomic.Value{}
	v.Store(person)
	person.Age = 33
	fmt.Println(v.Load())
}
