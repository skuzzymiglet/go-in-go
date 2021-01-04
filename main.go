package main

import (
	"log"
	"path/filepath"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	x := 0

	files, err := filepath.Glob("scripts/*.go")
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range files {
		log.Printf("running %s", filename)

		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols)
		i.Use(interp.Exports(map[string]map[string]reflect.Value{
			"lib": map[string]reflect.Value{
				"IncX": reflect.ValueOf(func() { x++ }),
				"DecX": reflect.ValueOf(func() { x-- }),
				"SetX": reflect.ValueOf(func(n int) { x = n }),
				"GetX": reflect.ValueOf(func() int { return x }),
			},
		}))

		_, err := i.EvalPath(filename)
		if err != nil {
			log.Fatalf("error running %s: %s", filename, err)
		}
		symbols := i.Symbols("")
		main, ok := symbols["main"]
		if !ok {
			panic("main not found")
		}
		runFuncValue, ok := main["Func"]
		if !ok {
			log.Fatalf("%s: Func not found", filename)
		}
		runFunc, ok := runFuncValue.Interface().(func() func())
		if !ok {
			log.Fatalf("%s: Run is not of type func() func()", filename)
		}
		runFunc()()
		log.Printf("x is now %d", x)
	}
}
