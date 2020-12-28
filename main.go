package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	files, err := filepath.Glob("scripts/*.go")
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range files {
		log.Printf("running %s", filename)
		start := time.Now()
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		i := interp.New(interp.Options{Stdin: file})
		i.Use(stdlib.Symbols)
		// sym := i.Symbols("go-in-go/lib")
		// pretty.Println(sym)
		// i.Use(sym)
		i.REPL()
		symbols := i.Symbols("main")
		main, ok := symbols["main"]
		if !ok {
			panic("main not found")
		}
		runFuncValue, ok := main["Run"]
		if !ok {
			log.Fatalf("%s: Run not found", filename)
		}
		runFunc, ok := runFuncValue.Interface().(func())
		if !ok {
			log.Fatalf("%s: Run is not of type func()", filename)
		}
		runFunc()
		end := time.Now()
		log.Printf("%s took %s", filename, end.Sub(start))
	}
}
