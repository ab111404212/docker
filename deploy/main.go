package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Test struct {
	Arg1 int        `wyj:"arg1"`
	Arg2 int        `wyj:"arg2"`
	Add1 func() int `wyj:"add1"`
}

func (t *Test) Add() int {
	return t.Arg1 + t.Arg2
}

type AddFunc func() int

func main() {
	p := fmt.Println
	tt := Test{10, 20, func() int {
		return 10
	}}
	v := reflect.Indirect(reflect.ValueOf(tt))
	t := reflect.TypeOf(tt)
	for i := 0; i < v.NumField(); i++ {
		v1 := v.Field(i)
		if v1.Kind() == reflect.Int {
			p("===v:", v1.Int())
		}
		if v1.Kind() == reflect.Func {
			var pp *AddFunc
			pp = (*AddFunc)(*(unsafe.Pointer)(v1.Pointer()))
			p("===v:", (*pp)())
		}
	}

	p(t.NumMethod(), v.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		t1 := t.Method(i)
		p("=====m:", t1.Name)
	}

	for i := 0; i < t.NumField(); i++ {
		t1 := t.Field(i)
		p("=====t:", t1.Tag.Get("wyj"), t1.Name, t1.Type)
	}
	p(tt.Add())
	p("first github...")
}
