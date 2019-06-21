package main

import (
	"fmt"
	"github.com/mvozaar/ellipse"
)

func main() {
	fmt.Println("Hello World!!!")
	e := ellipse.Init{
		9, 2,
	}
	fmt.Println(e.A)

	fmt.Println(e.GetEccentricity())
}
