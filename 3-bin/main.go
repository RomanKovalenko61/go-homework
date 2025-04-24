package main

import (
	"app/hw-3/bins"
	"fmt"
)

func main() {
	fmt.Println("project 3-bin")
	bin := bins.NewBin("1", true, "newBin")
	fmt.Printf("%v", *bin)
}
