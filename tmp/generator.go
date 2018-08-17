package mashup

import (
	"log"
	"math/rand"
	"strings"
	"sync"

	"github.com/zpatrick/rhyme/rhyme"
	"github.com/zpatrick/rhyme/song"
)

// todo: use channels
// todo: dynamically create matrix of which lines match with other lines
// Generate() pulls 2 pairs from the matrix

type Generator struct {
	rhymer   rhyme.Rhymer
	searcher song.Searcher
	songs    []*song.Song
	rhymes   map[string][]string
	mutex    *sync.Mutex
}

func NewGenerator(rhymer rhyme.Rhymer, searcher song.Searcher) *Generator {
	return &Generator{
		rhymer:   rhymer,
		searcher: searcher,
		songs:    []*song.Song{},
		rhymes:   map[string][]string{},
		mutex:    &sync.Mutex{},
	}
}

func (g *Generator) Start() {
	go g.lookupSongLyrics()
}

func (g *Generator) lookupSongLyrics() {
	for _, song := range Catalog() {
		log.Printf("[INFO] Looking up lyrics for '%s by %s'", song.Title, song.Artist)
		song, err := g.searcher.Search(song.Title, song.Artist)
		if err != nil {
			log.Println("[ERROR]", err.Error())
			continue
		}

		go g.lookupRhymes(song.Lyrics)
		g.mutex.Lock()
		g.songs = append(g.songs, song)
		g.mutex.Unlock()
	}

	log.Printf("[INFO] Done looking up lyrics")
}

func (g *Generator) lookupRhymes(lines []song.Line) {
	for _, line := range lines {
		word := line[len(line)-1]

		g.mutex.Lock()
		_, ok := g.rhymes[word]
		g.mutex.Unlock()

		if !ok {
			matches, err := g.rhymer.Rhyme(word)
			if err != nil {
				log.Println("[ERROR]", err.Error())
				continue
			}

			g.mutex.Lock()
			g.rhymes[word] = matches
			g.mutex.Unlock()
		}
	}
}

func (g *Generator) randomLine() (Line, string) {
	song := g.songs[rand.Intn(len(g.songs))]
	words := song.Lyrics[rand.Intn(len(song.Lyrics))]
	line := Line{
		Text: strings.Join(words, " "),
		URL:  song.URL,
	}

	return line, words[len(words)-1]
}

func (g *Generator) Generate() Mashup {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	mashup := Mashup{}
	for len(mashup) < 4 {
		lineA, wordA := g.randomLine()
		lineB, wordB := g.randomLine()
		for _, word := range g.rhymes[wordA] {
			if word == wordB {
				mashup = append(mashup, lineA, lineB)
			}
		}
	}

	return mashup
}
