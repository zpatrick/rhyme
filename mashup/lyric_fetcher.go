package mashup

type LyricFetcher interface {
	FetchLyrics(songID int) (Lyrics, error)
}

type LyricFetcherFunc func(songID int) (Lyrics, error)

func (l LyricFetcherFunc) FetchLyrics(songID int) (Lyrics, error) {
	return l(songID)
}
