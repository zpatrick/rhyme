package mashup

type Generator interface {
	Generate() (Verse, error)
}

type GeneratorFunc func() (Verse, error)

func (g GeneratorFunc) Generate() (Verse, error) {
	return g()
}
