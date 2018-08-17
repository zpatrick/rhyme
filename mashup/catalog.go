package mashup

import "github.com/zpatrick/rhyme/song"

// https://www.ranker.com/list/best-rap-songs-of-all-time/ranker-hip-hop
func Catalog() []song.Song {
	return []song.Song{
		{Title: "Lose Yourself", Artist: "Eminem"},
		{Title: "Juicy", Artist: "The Notorious B.I.G."},
		{Title: "Straight Outta Compton", Artist: "N.W.A."},
		{Title: "Dear Mama", Artist: "Tupac"},
		{Title: "Still D.R.E.", Artist: "Dr. Dre"},
		{Title: "It Was a Good Day", Artist: "Ice Cube"},
		{Title: "Shook Ones, Part II", Artist: "Mobb Deep"},
		{Title: "Stan", Artist: "Eminem"},
		{Title: "C.R.E.A.M.", Artist: "Wu-tang Clan"},
		{Title: "Nuthin' But a 'G' Thang", Artist: "Dr. Dre"},
	}
}
