package mashup

type MatrixGenerator struct{}

func (m MatrixGenerator) Generate() Verse {
	return Verse{}
}

/*
Start at the end:
-We have 4 lines
-We have 2 pairs of matching lines
-We match the last word of each line
-We know words match by seeing if they have eachother in a rhymep map
-Quittable


// Step 1: Build a wordRhymes map[string][]string for the last word in each line of each song
As each song's lyrics come in (range over channel):
  For each of the last words in each line:
    If that word is not already in our map:
      We lookup that word's rhymes using a helper

*/

//func generateRhymes(c <-chan Line, r Rhymer) error {

func generateLines(songIDs []int, l LyricFetcher, update func(Line)) error {
	for _, songID := range songIDs {
		lyrics, err := l.FetchLyrics(songID)
		if err != nil {
			return err
		}

		for _, line := range lyrics {
			update(line)
		}
	}

	return nil
}

func generateRhymes(c <-chan Line, r Rhymer, update func(string, []string)) error {
	wordRhymes := map[string][]string{}
	for line := range c {
		word := line.Last()
		if _, ok := wordRhymes[word]; !ok {
			rhymes, err := r.Rhyme(word)
			if err != nil {
				return err
			}

			update(word, rhymes)
		}
	}

	return nil
}
