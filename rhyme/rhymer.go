package rhyme

type Rhymer interface {
	Rhyme(word string) ([]string, error)
}

type RhymerFunc func(word string) ([]string, error)

func (r RhymerFunc) Rhyme(word string) ([]string, error) {
	return r(word)
}
