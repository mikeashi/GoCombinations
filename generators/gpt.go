package generators

type GPTGenerator struct {
}

// Generate generates all permutations with repetition of the given alphabet
// from min length to max length and appends them to the results slice.
func (gpt GPTGenerator) Generate(alphabet string, min int, max int, results *[]string) {
	chars := []rune(alphabet)
	var recurse func(prefix []rune, minLength, maxLength int)
	recurse = func(prefix []rune, minLength, maxLength int) {
		if len(prefix) >= minLength {
			*results = append(*results, string(prefix))
		}
		if len(prefix) == maxLength {
			return
		}
		for _, char := range chars {
			recurse(append([]rune{}, append(prefix, char)...), minLength, maxLength)
		}
	}
	recurse([]rune{}, min, max)
}

// GetName returns the name of the generator.
func (gpt GPTGenerator) GetName() string {
	return "GPTGenerator"
}
