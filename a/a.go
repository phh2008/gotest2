package a

import "strconv"

type A struct {
	Name string
	Age  int8
}

func (a A) Str() string {
	return a.Name + "：" + strconv.Itoa(int(a.Age))
}

type Integer int64

func (a Integer) Str() string {
	return strconv.Itoa(int(a))
}

type Int *int64

//func (a Int) Str2()  {
// 指针类型不能实现方法
//}
