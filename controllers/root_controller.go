package controllers

import (
	"net/http"

	"github.com/zpatrick/fireball"
)

type RootController struct{}

func NewRootController() *RootController {
	return &RootController{}
}

func (r *RootController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": r.redirect,
			},
		},
	}

	return routes
}

func (r *RootController) redirect(c *fireball.Context) (fireball.Response, error) {
	return fireball.Redirect(http.StatusTemporaryRedirect, "/verse"), nil
}
