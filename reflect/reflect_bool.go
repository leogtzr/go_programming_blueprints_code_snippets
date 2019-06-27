package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(true).Size())
	fmt.Println(reflect.TypeOf(struct{}{}).Size())
}
