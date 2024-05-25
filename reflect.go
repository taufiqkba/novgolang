package main

import (
	"fmt"
	"reflect"
)

// Accessing Object Variable Property Information
type studentReflect struct {
	Name  string
	Grade int
}

func (s *studentReflect) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("name: ", reflectType.Field(i).Name)
		fmt.Println("data type: ", reflectType.Field(i).Type)
		fmt.Println("value: ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

// access object variable method information
func (s *studentReflect) SetName(name string) {
	s.Name = name
}

func main() {
	// reflect on Golang
	var number = 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("variable type:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("variable value:", reflectValue.Int())
		fmt.Println("variable value:", reflectValue.Interface())
	}
	// var value = reflectValue.Interface().(int)

	// accessing object variable property information
	var s1 = &studentReflect{Name: "john wick", Grade: 2}
	s1.getPropertyInfo()

	// access object variable method information
	fmt.Println("name is:", s1.Name)

	reflectValue = reflect.ValueOf(s1)
	method := reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wicked"),
	})

	fmt.Println("name is:", s1.Name)
}
