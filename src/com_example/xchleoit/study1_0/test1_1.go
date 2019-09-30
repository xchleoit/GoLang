package main

import (
	user "com_example/xchleoit/study1_0/module/"
	"fmt"
	"math"
)

func main() {
	var key = 20.2
	var result = math.Pi * math.Pow(key, 2)
	fmt.Println("the circle result is ", result)
	user.Add()
	user.BAdd()
}
