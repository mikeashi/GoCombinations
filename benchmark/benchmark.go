package benchmark

import (
	"time"

	"github.com/mikeashi/GoCombinations/generators"
)

func Benchmark(generator generators.Generator, alphabet string, min int, max int) (duration time.Duration, rate float64, combinations int) {
	results := []string{}
	startTime := time.Now()
	generator.Generate(alphabet, min, max, &results)
	duration = time.Since(startTime)
	combinations = len(results)
	rate = float64(combinations) / duration.Seconds()
	return
}
