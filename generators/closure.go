package generators

type Closure struct {
}

func (c Closure) Generate(alphabet string, min int, max int, results *[]string) {
	*results = []string{}
	generate("", min, max, alphabet, results)
}

func (c Closure) GetName() string {
	return "Closure"
}

func letterGenerator(alphabet string) func() string {
	letters := alphabet
	return func() string {
		if len(letters) > 0 {
			letter := string(letters[0])
			letters = letters[1:]
			return letter
		}
		return ""
	}
}

func generate(current string, min int, max int, alphabet string, results *[]string) {
	if len(current) >= min {
		*results = append(*results, current)
	}

	if len(current) >= max {
		return
	}

	generator := letterGenerator(alphabet)
	for nextLetter := generator(); nextLetter != ""; nextLetter = generator() {
		generate(current+nextLetter, min, max, alphabet, results)
	}
}
