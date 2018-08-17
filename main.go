package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/rhyme/controllers"
	"github.com/zpatrick/rhyme/mashup"
	"github.com/zpatrick/rhyme/rhyme"
	"github.com/zpatrick/rhyme/song"
)

const (
	FlagPort        = "port"
	EVPort          = "RM_PORT"
	FlagGeniusToken = "genius-token"
	EVGeniusToken   = "RM_GENIUS_TOKEN"
)

func main() {
	app := cli.NewApp()
	app.Name = "rhyme"
	app.Usage = "Rhyme Generator"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   FlagPort,
			EnvVar: EVPort,
			Value:  80,
		},
		cli.StringFlag{
			Name:   FlagGeniusToken,
			EnvVar: EVGeniusToken,
		},
	}

	app.Before = func(c *cli.Context) error {
		rand.Seed(time.Now().UTC().UnixNano())

		requiredFlags := []string{
			FlagGeniusToken,
		}

		for _, flag := range requiredFlags {
			if !c.IsSet(flag) {
				return fmt.Errorf("Required flag '%s' is not set!", flag)
			}
		}

		return nil
	}

	app.Action = func(c *cli.Context) error {
		rhymer := rhyme.NewRhymeBrainRhymer()
		genius := song.NewGeniusClient(c.String(FlagGeniusToken))
		generator := mashup.NewGenerator(rhymer, genius)
		go generator.Start()

		rootController := controllers.NewRootController()
		verseController := controllers.NewVerseController(generator)
		routes := append(rootController.Routes(), verseController.Routes()...)
		app := fireball.NewApp(routes)
		http.Handle("/", app)

		fs := http.FileServer(http.Dir("static"))
		http.Handle("/static/", http.StripPrefix("/static", fs))

		addr := fmt.Sprintf("0.0.0.0:%d", c.Int(FlagPort))
		log.Printf("[INFO] Listening on %s\n", addr)
		return http.ListenAndServe(addr, nil)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
