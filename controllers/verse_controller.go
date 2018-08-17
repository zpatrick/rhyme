package controllers

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/rhyme/mashup"
)

type VerseController struct {
	generator *mashup.Generator
}

func NewVerseController(g *mashup.Generator) *VerseController {
	return &VerseController{
		generator: g,
	}
}

func (v *VerseController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/verse",
			Handlers: map[string]fireball.Handler{
				"GET": v.getVerse,
			},
		},
	}

	return routes
}

func (v *VerseController) getVerse(c *fireball.Context) (fireball.Response, error) {
	return c.HTML(200, "index.html", v.generator.Generate())
}
