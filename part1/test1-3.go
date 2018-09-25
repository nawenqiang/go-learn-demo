package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	Life int
}

func (b Bird) Fly() {
	fmt.Println("bird fly")
}

func main() {
	bird := &Bird{Name: "bird", Life: 2}
	s := reflect.ValueOf(bird).Elem()
	s_type := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		//fmt.Printf("%d: %s %s = %v\n", i, s_type.Field(i).Name, f.Type(), f.Interface())
		fmt.Printf("%d: %s %s = %v\n", i, s_type.Field(i).Name, s_type.Field(i).Type, f.Interface())
	}
}
