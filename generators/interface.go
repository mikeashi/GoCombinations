package generators

type Generator interface {
	Generate(alphabet string, min int, max int, results *[]string)
	GetName() string
}
