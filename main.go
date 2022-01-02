package main

import (
	"ci-cd-test-go/compute"
	"fmt"
)

func main() {
	x, y := 20, 10
	fmt.Println(compute.Add(x, y))
}
