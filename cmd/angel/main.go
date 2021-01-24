package main

import (
	"log"
	"os"

	"github.com/daite/angel/common"
	"github.com/daite/angel/jtorrent"
	"github.com/daite/angel/ktorrent"
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Usage:   "choose torrent sites (kr or jp)",
			},
		},
		Action: func(c *cli.Context) error {
			keyword := "동상이몽2"
			if c.NArg() > 0 {
				keyword = c.Args().Get(0)
			}
			if c.String("lang") == "kr" {
				s := []common.Scraping{
					&ktorrent.TToBoGo{},
					&ktorrent.TorrentView{},
					&ktorrent.TorrentMobile{},
					&ktorrent.TorrentTube{},
					&ktorrent.TShare{},
				}
				s = common.GetAvailableSites(s)
				data := common.CollectData(s, keyword)
				common.PrintData(data, false)
			} else {
				s := []common.Scraping{
					&jtorrent.Nyaa{},
					&jtorrent.SuKeBe{},
				}
				data := common.CollectData(s, keyword)
				common.PrintData(data, false)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
