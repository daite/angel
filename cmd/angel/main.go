package main

import (
	"log"
	"os"

	"github.com/daite/angel/sites"
	"github.com/urfave/cli/v2"
)

var version = "0.1.0"

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}
	app := &cli.App{
		Name:    "angel",
		Usage:   "search torrent magnet!",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "search torrent magnet file",
				Action: func(c *cli.Context) error {
					keyword := c.Args().First()
					s := []sites.Scraping{
						&sites.TToBoGo{},
						&sites.TorrentView{},
						&sites.TorrentMobile{},
						&sites.TorrentTube{},
						&sites.TShare{},
						&sites.Nyaa{},
					}
					data := sites.CollectData(s, keyword)
					sites.PrintData(data, false)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
