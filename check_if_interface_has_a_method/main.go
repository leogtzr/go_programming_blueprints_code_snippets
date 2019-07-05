package main

import "fmt"

func foo(x interface{}) {
	if _, ok := x.(interface{ cmn() string }); ok {
		fmt.Println("There!")
	} else {
		fmt.Println("No there ... ")
	}
}

type X struct {
}

func (x X) cmn() string {
	return "hello"
}

func main() {
	// var x string = "Hello"
	// x := X{}
	// foo(X{})

	// var x interface{} = X{}
	foo(X{})
}
