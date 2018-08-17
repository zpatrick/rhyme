package song

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zpatrick/rclient"
)

type GeniusResponse struct {
	Response GeniusSearchResponse `json:"response"`
}

type GeniusSearchResponse struct {
	Hits []GeniusHit `json:"hits"`
}

type GeniusHit struct {
	Type   string          `json:"type"`
	Result GeniusHitResult `json:"result"`
}

type GeniusHitResult struct {
	ID            int          `json:"id"`
	Title         string       `json:"title"`
	URL           string       `json:"url"`
	PrimaryArtist GeniusArtist `json:"primary_artist"`
}

type GeniusArtist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GeniusClient struct {
	client *rclient.RestClient
}

func NewGeniusClient(token string) *GeniusClient {
	header := rclient.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	return &GeniusClient{
		client: rclient.NewRestClient("https://api.genius.com", rclient.RequestOptions(header)),
	}
}

func (g *GeniusClient) Search(title, artist string) (*Song, error) {
	hit, err := g.findSong(title, artist)
	if err != nil {
		return nil, err
	}

	lyrics, err := g.scrapeSongLyrics(hit.URL)
	if err != nil {
		return nil, err
	}

	song := &Song{
		Title:  hit.Title,
		Artist: hit.PrimaryArtist.Name,
		URL:    hit.URL,
		Lyrics: lyrics,
	}

	return song, nil
}

func (g *GeniusClient) findSong(title, artist string) (*GeniusHitResult, error) {
	q := url.Values{}
	q.Set("q", fmt.Sprintf("%s by %s", title, artist))

	var resp GeniusResponse
	if err := g.client.Get("/search", &resp, rclient.Query(q)); err != nil {
		return nil, err
	}

	for _, hit := range resp.Response.Hits {
		if hit.Type == "song" {
			switch {
			case strings.ToLower(title) == strings.ToLower(hit.Result.Title):
				return &hit.Result, nil
			case strings.ToLower(artist) == strings.ToLower(hit.Result.PrimaryArtist.Name):
				return &hit.Result, nil
			}
		}
	}

	return nil, fmt.Errorf("No songs found matching '%s by %s'", title, artist)
}

func (g *GeniusClient) scrapeSongLyrics(songURL string) ([]Line, error) {
	response, err := http.Get(songURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	result := strings.Split(doc.Find(".lyrics").First().Text(), "\n")
	lines := make([]Line, 0, len(result))
	for _, r := range result {
		words := strings.Split(r, " ")
		line := make([]string, 0, len(words))
		for _, word := range words {
			if len(word) > 0 {
				line = append(line, word)
			}
		}

		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	// todo: replace words so they can rhyme

	return lines, nil
}
