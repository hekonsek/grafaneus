package main

import (
	"github.com/grafaneus"
	"fmt"
)

func main() {
	fmt.Println(grafaneus.ListTemplates())

	fmt.Println()

	json := grafaneus.GenerateGraph("Goroutines number", "go_goroutines")
	fmt.Println(json)
}