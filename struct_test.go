package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Per struct {
	Name string
}

func (a Per) Show() {
	a.Name = a.Name + "ccc"
}

func (a Per) String() string {
	return a.Name
}

func TestPer(t *testing.T) {
	var p = &Per{Name: "xa"}
	p.Show()
	fmt.Println(p.Name)
}

func TestSliceI(t *testing.T) {
	var arr []Per
	arr = append(arr, Per{Name: "n1"})
	arr = append(arr, Per{Name: "n2"})
	arr = append(arr, Per{Name: "n3"})
	for i, v := range arr {
		//v.Name = v.Name + "_***" + strconv.Itoa(i) // 并不能改变元素，切片换成指针元素则可以
		arr[i].Name = v.Name + "_***" + strconv.Itoa(i) // 不是指针元素时，想要修改元素用此方式
	}
	fmt.Println(arr)
}

func TestMapI(t *testing.T) {
	var mp = map[string]Per{}
	mp["a"] = Per{Name: "n1"}
	mp["b"] = Per{Name: "n2"}
	mp["c"] = Per{Name: "n3"}
	var i int
	for _, v := range mp {
		i++
		//mp[k].Name = v.Name + "_***" + k // value是指针时才可以寻址，否则编译就会出错
		v.Name = v.Name + "_***" + strconv.Itoa(i) // value是指针元时，修改才会变化
	}
	fmt.Println(mp)
}
