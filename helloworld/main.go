package main

import (
	"fmt"
	"helloworld/calculation"
)

func main() {
	fmt.Println("Hello Isekai!")
	res := calculation.Add(5, 6)
	fmt.Println(res)
}
