package main

import (
	"fmt"
	"lib"
	"math/rand"
)

func Func() func() {
	return func() {
		fmt.Println(lib.GetX())
		lib.SetX(rand.Int())
	}
}
