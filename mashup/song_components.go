package mashup

type Lyrics []Line

type Verse []Line

type Line []string

func (l Line) Last() string {
	if len(l) == 0 {
		return ""
	}

	return l[len(l)-1]
}
