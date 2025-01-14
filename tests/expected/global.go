package main

import (
	"fmt"
	"github.com/adsharma/py2many/pygo/runtime"
)

var Code0 int = 0
var Code1 int = 1
var LA []int = []int{Code0, Code1}
var CodeA string = "a"
var CodeB string = "b"
var LB []string = []string{CodeA, CodeB}

func main() {
	for _, i := range LA {
		fmt.Printf("%v\n", i)
	}
	for _, i := range LB {
		fmt.Printf("%v\n", i)
	}
	if pygo.Contains([]string{"a", "b"}, "a") {
		fmt.Printf("%v\n", "OK")
	}
}
