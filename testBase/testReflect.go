package main

import (
	"fmt"
	"reflect"
)

type Bago struct {
	Name string
	Age  int
}

func main() {
	bago := Bago{"teeth", 22}
	//typeOfBago := reflect.TypeOf(bago)
	//fmt.Println(typeOfBago.Name())
	//fmt.Println(typeOfBago.Kind())
	valueOfBago := reflect.ValueOf(bago)
	fmt.Println(valueOfBago.Field(0))
	fmt.Println(valueOfBago)
	valueOfBago = reflect.ValueOf(&bago)
	v := valueOfBago.Elem()
	v.FieldByName("Name").SetString("teethex")
	v.FieldByName("Age").SetInt(95)
	fmt.Println(bago)
	//bagoptr := &Bago{"dzq087", 15}
	//typeOfBago = reflect.TypeOf(bagoptr)
	//fmt.Println(typeOfBago.Name())
	//fmt.Println(typeOfBago.Kind())
	//typeOfBago = typeOfBago.Elem()
	//fmt.Println(typeOfBago.Name())
	//fmt.Println(typeOfBago.Kind())
}
