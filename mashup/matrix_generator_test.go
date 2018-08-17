package mashup

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLines(t *testing.T) {
	result := []Line{}
	update := func(l Line) error {
		result = append(result, l)
		return nil
	}

	songIDs := []int{123, 456}
	lyricFetcher := LyricFetcherFunc(func(songID int) (Lyrics, error) {
		var lyrics Lyrics
		switch songID {
		case 123:
			lyrics = Lyrics{
				strings.Split("Hey Bonita, glad to meet ya", " "),
				strings.Split("For the kind of stunning newness, I must beseech ya", " "),
			}
		case 456:
			lyrics = Lyrics{
				strings.Split("Can I kick it? To all the people who can Quest like A Tribe does", " "),
				strings.Split("Before this, did you really know what live was?", " "),
			}
		default:
			t.Fatalf("Unexpected songID: %d", songID)
		}

		return lyrics, nil
	})

	if err := generateLines(songIDs, lyricFetcher, update); err != nil {
		t.Fatal(err)
	}

	expected := []Line{
		strings.Split("Hey Bonita, glad to meet ya", " "),
		strings.Split("For the kind of stunning newness, I must beseech ya", " "),
		strings.Split("Can I kick it? To all the people who can Quest like A Tribe does", " "),
		strings.Split("Before this, did you really know what live was?", " "),
	}
	assert.Equal(t, result, expected)
}

func TestGenerateRhymes(t *testing.T) {
	result := map[string][]string{}
	update := func(word string, rhymes []string) error {
		result[word] = rhymes
		return nil
	}

	rhymer := RhymerFunc(func(word string) ([]string, error) {
		switch word {
		case "amazing":
			return []string{"matrix", "vacation"}, nil
		case "matrix":
			return []string{"amazing", "vacation"}, nil
		case "9":
			return []string{"mind"}, nil
		case "vacation":
			return []string{"amazing", "matrix"}, nil
		default:
			t.Errorf("Unexpected word: %s", word)
			return nil, nil
		}
	})

	for _, line := range []string{
		"Yeah - I feel amazing",
		"Yeah -I'm in the matrix",
		"My mind is living on cloud 9",
		"and this nine is never on vacation"} {
		if err := generateRhymes(strings.Split(line, " "), rhymer, update); err != nil {
			t.Fatal(err)
		}
	}

	expected := map[string][]string{
		"amazing":  []string{"matrix", "vacation"},
		"matrix":   []string{"amazing", "vacation"},
		"9":        []string{"mind"},
		"vacation": []string{"amazing", "matrix"},
	}

	assert.Equal(t, expected, result)
}
