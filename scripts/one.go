package main

import (
	"fmt"

	"./lib"
)

func Run() {
	lib.SetPix(5)
	if lib.GetPix() != 5 {
		fmt.Println("disaster")
	} else {
		fmt.Println("wohoo!")
	}
}
