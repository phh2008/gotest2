package main

import (
	"fmt"
	"github.com/imdario/mergo"
	"log"
	"reflect"
	"testing"
	"time"
)

type Foo struct {
	A         string
	B         int64
	CreatedAt time.Time
	amount    float64
}

type timeTransformer struct {
}

func (t timeTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(time.Time{}) {
		return func(dst, src reflect.Value) error {
			if dst.CanSet() {
				isZero := dst.MethodByName("IsZero")
				result := isZero.Call([]reflect.Value{})
				if result[0].Bool() {
					dst.Set(src)
				}
			}
			return nil
		}
	}
	return nil
}

func TestMergo1(t *testing.T) {
	src := Foo{
		A:         "one",
		B:         2,
		CreatedAt: time.Now(),
		amount:    3.14,
	}
	dest := Foo{
		A: "two",
	}
	//mergo.Merge(&dest, src, mergo.WithTransformers(timeTransformer{}))
	mergo.Merge(&dest, src, mergo.WithOverride)
	fmt.Printf("%+v", dest)
}

func TestMergo2(t *testing.T) {
	src := Foo{
		//A:         "one",
		B:         2,
		CreatedAt: time.Now(),
		amount:    3.14,
	}
	map1 := map[string]interface{}{
		"a": "tow2",
		//"createdAt": time.Now(),
	}
	mergo.Map(&map1, src)
	fmt.Println(map1)
	fmt.Println(src)
}

type User1 struct {
	Name string
}

type User2 struct {
	Name string
	Age  int
}

func TestMergo3(t *testing.T) {
	var des = User1{Name: "li"}
	var src = User2{Name: "Li2", Age: 33}
	err := mergo.Map(&des, &src)
	log.Println(err)
	log.Println(des)
}
