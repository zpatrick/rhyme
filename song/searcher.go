package song

type Searcher interface {
	Search(title, artist string) (*Song, error)
}
