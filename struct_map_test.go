package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

type People struct {
	Name     string `mapstruct:"name"`
	Age      int    `mapstruct:"age"`
	Birthday time.Time
}

type A struct {
	list []*People
}

func TestSlice1(t *testing.T) {
	tom := People{Name: "tom", Age: 21, Birthday: time.Now()}
	tom2 := People{Name: "tom2", Age: 22, Birthday: time.Now()}
	var list = []*People{&tom, &tom2}
	var a = A{}
	a.list = list
	fmt.Println(a.list)
	for i, v := range a.list {
		v.Name = v.Name + "-" + strconv.Itoa(i)
		//a.list[i] = v
	}
	fmt.Println(a.list)
}

func TestStruct2Map(t *testing.T) {
	tom := People{Name: "tom", Age: 22, Birthday: time.Now()}
	typ := reflect.TypeOf(tom)
	val := reflect.ValueOf(tom)
	num := typ.NumField()
	mapInfo := make(map[string]interface{}, num)
	for i := 0; i < num; i++ {
		tag := typ.Field(i).Tag.Get("mapstruct")
		if tag != "" {
			mapInfo[tag] = val.Field(i).Interface()
		} else {
			if val.Field(i).CanInterface() {
				mapInfo[typ.Field(i).Name] = val.Field(i).Interface()
			}
		}
	}
	fmt.Println(mapInfo)
}

func TestLength(t *testing.T) {
	var list []People
	var m map[string]People
	fmt.Println(len(list))
	fmt.Println(len(m))
	for i, v := range list {
		fmt.Println(i, v)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("-----------------")
}

func TestList2(t *testing.T) {
	var list []int64
	var list2 []int64 = make([]int64, 1)
	fmt.Println(len(list2), cap(list2))
	list = append(list, list2...)
	fmt.Println(list, len(list), cap(list), list == nil)
	fmt.Println("--------------------")
	fmt.Println(len([]int64(nil)))
}
