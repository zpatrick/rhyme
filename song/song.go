package song

type Song struct {
	Title  string
	Artist string
	Lyrics Lyrics
}

type Lyrics []Verse

type Verse []Word

type Word string
