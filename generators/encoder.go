package generators

import (
	"github.com/mikeashi/GoCombinations/utils"
)

type Encoder struct {
}

func (e Encoder) Generate(alphabet string, min int, max int, results *[]string) {
	*results = []string{}
	start := calcStart(alphabet, min)
	end := calcEnd(alphabet, min, max)
	for i := start; i < end; i++ {
		*results = append(*results, getCombination(alphabet, i))
	}
}

func (e Encoder) GetName() string {
	return "Encoder"
}

func getCombination(alphabet string, n int) string {
	base := len(alphabet)
	n++
	var str string = ""

	for n > 0 {
		n--
		index := n % base
		str = string(alphabet[index]) + str
		n /= base
	}
	return str
}

func calcStart(alphabet string, min int) int {
	if min-1 == 0 {
		return 0
	}
	return utils.IntPow(len(alphabet), min-1)
}

func calcEnd(alphabet string, min int, max int) int {
	var e int
	for length := min - 1; length <= max; length++ {
		e += utils.IntPow(len(alphabet), length)
	}
	return e - 1
}
