package song

type Song struct {
	Title  string
	Artist string
	URL    string
	Lyrics []Line
}

type Line []string
