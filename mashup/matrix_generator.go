package mashup

import (
	"math/rand"
	"strings"
)

type Generator func() Verse

func NewMatrixGenerator(songIDs []int, l LyricFetcher, r Rhymer) (Generator, error) {
	wordRhymes := map[string][]string{}
	updateWordRhymes := func(word string, rhymes []string) error {
		wordRhymes[word] = rhymes
		return nil
	}

	lines := []Line{}
	updateLines := func(line Line) error {
		lines = append(lines, line)
		return generateRhymes(line, r, updateWordRhymes)
	}

	if err := generateLines(songIDs, l, updateLines); err != nil {
		return nil, err
	}

	/*
		if err := generateLines(songIDs, l, func(line Line) {
			lines = append(lines, line)
			generateRhymes(line, r, func(word string, rhymes []string) {
				wordRhymes[word] = rhymes
			})
		}); if err != nil {
			return err
		}
	*/

	matrix := map[int][]Line{}
	generateMatrix(lines, wordRhymes, func(lineIndex int, lines []Line) {
		matrix[lineIndex] = lines
	})

	return func() Verse {
		verse := Verse{}
		for len(verse) < 4 {
			lineIndex := rand.Intn(len(lines))
			line := lines[lineIndex]
			matches := matrix[lineIndex]
			if len(matches) == 0 {
				continue
			}

			verse = append(verse, line, matches[rand.Intn(len(matches))])
		}

		return verse
	}, nil
}

func generateMatrix(lines []Line, wordRhymes map[string][]string, update func(lineIndex int, lines []Line)) {
	for lineIndex, line := range lines {
		word := strings.ToLower(line.Last())
		rhymes, ok := wordRhymes[word]
		if !ok {
			continue
		}

		acceptableWords := map[string]bool{}
		for _, rhyme := range rhymes {
			acceptableWords[rhyme] = true
		}

		matches := []Line{}
		for i, potentialMatch := range lines {
			if i == lineIndex {
				continue
			}

			word := strings.ToLower(potentialMatch.Last())
			if _, ok := acceptableWords[word]; ok {
				matches = append(matches, line)
			}
		}

		update(lineIndex, lines)
	}
}

func generateLines(songIDs []int, l LyricFetcher, update func(Line) error) error {
	for _, songID := range songIDs {
		lyrics, err := l.FetchLyrics(songID)
		if err != nil {
			return err
		}

		for _, line := range lyrics {
			if err := update(line); err != nil {
				return err
			}
		}
	}

	return nil
}

func generateRhymes(line Line, r Rhymer, update func(word string, rhymes []string) error) error {
	word := line.Last()
	rhymes, err := r.Rhyme(word)
	if err != nil {
		return err
	}

	return update(word, rhymes)
}
