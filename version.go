package main

import (
	"fmt"
	"runtime"
)

func main() {
	v := runtime.Version()
	fmt.Println("go version : ", v)
}
