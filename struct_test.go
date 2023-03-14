package main

import (
	"fmt"
	"testing"
)

type Per struct {
	Name string
}

func (a Per) Show() {
	a.Name = a.Name + "ccc"
}

func TestPer(t *testing.T) {
	var p = &Per{Name: "xa"}
	p.Show()
	fmt.Println(p.Name)
}
