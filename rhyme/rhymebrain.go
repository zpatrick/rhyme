package rhyme

import (
	"net/url"

	"github.com/zpatrick/rclient"
)

// The amount in which two words rhyme is given a score between 0 and 300.
// Words that don't meet MinRequiredScore are not returned.
const MinRequiredScore = 250

type RhymeBrainMatch struct {
	Word  string
	Score int
}

func NewRhymeBrainRhymer() RhymerFunc {
	client := rclient.NewRestClient("http://rhymebrain.com")
	return func(word string) ([]string, error) {
		q := url.Values{}
		q.Set("function", "getRhymes")
		q.Set("word", word)

		var matches []RhymeBrainMatch
		if err := client.Get("/talk", &matches, rclient.Query(q)); err != nil {
			return nil, err
		}

		words := make([]string, 0, len(matches))
		for i := 0; i < len(matches); i++ {
			if matches[i].Score >= MinRequiredScore {
				words = append(words, matches[i].Word)
			}
		}

		return words, nil
	}
}
