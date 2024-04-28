package generators

import (
	"strings"

	"github.com/mowshon/iterium"
)

type Iterium struct {
}

func (i Iterium) Generate(alphabet string, min int, max int, results *[]string) {
	chars := []rune(alphabet)
	// Convert rune slice to a slice of strings for iterium.Product
	stringChars := make([]string, len(chars))
	for i, char := range chars {
		stringChars[i] = string(char)
	}

	for length := min; length <= max; length++ {
		product := iterium.Product[string](stringChars, length)
		permutations, _ := product.Slice()
		for _, perm := range permutations {
			*results = append(*results, strings.Join(perm, ""))
		}
	}
}

func (i Iterium) GetName() string {
	return "Iterium"
}
