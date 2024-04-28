package main

import (
	"fmt"

	go_console "github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"

	"github.com/mikeashi/GoCombinations/benchmark"
	"github.com/mikeashi/GoCombinations/generators"
)

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const MIN = 1
const MAX = 4

func main() {
	cmd := go_console.NewScript().Build()

	tab := table.
		NewTable().
		AddHeadersFromString(
			[][]string{
				{"Implementation", "Duration", "Rate", "Combinations"},
			},
		)

	implementations := []generators.Generator{
		generators.Closure{},
		generators.Encoder{},
		generators.GPTGenerator{},
		generators.Iterium{},
	}

	fmt.Println("Running benchmarks:")
	fmt.Printf("Min: %d , Max: %d \n", MIN, MAX)
	fmt.Println("Alphabet: ", LETTERS)

	for _, impl := range implementations {
		duration, rate, combinations := benchmark.Benchmark(impl, LETTERS, MIN, MAX)
		tab.AddRowFromString([]string{
			impl.GetName(), fmt.Sprintf("%v", duration), fmt.Sprintf("%f", rate), fmt.Sprintf("%d", combinations),
		})
	}

	render := table.
		NewRender(cmd.Output).
		SetContent(tab)

	render.Render()
}
