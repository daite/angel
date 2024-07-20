package main

import (
	"fmt"
	"os"

	"github.com/daite/angel/common"
	"github.com/daite/angel/jtorrent"
	"github.com/daite/angel/ktorrent"
	"github.com/urfave/cli/v2"
)

var version = "0.7.0"

func main() {
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
			// Removed the `version` flag from here
		},
		Action: func(c *cli.Context) error {
			// Handle the -v flag to print the version
			if c.Bool("version") {
				fmt.Printf("%s version %s\n", c.App.Name, c.App.Version)
				return nil
			}

			// Print usage information if no arguments are provided
			if c.Args().Len() == 0 && !c.IsSet("lang") {
				cli.ShowAppHelp(c)
				return nil
			}

			// Print usage information if the -l flag is provided
			if c.IsSet("lang") {
				cli.ShowAppHelp(c)
				return nil
			}

			// Proceed with regular processing if a language is specified
			if lang := c.String("lang"); lang == "kr" {
				s := []common.Scraping{
					&ktorrent.TorrentMobile{},
					&ktorrent.TorrentView{},
					&ktorrent.TorrentSir{},
					&ktorrent.TorrentJ{},
					&ktorrent.TorrentSee{},
					&ktorrent.JuJuTorrent{},
					&ktorrent.TorrentQQ{},
					&ktorrent.TorrentSome{},
					&ktorrent.TorrentRJ{},
				}
				s = common.GetAvailableSites(s)
				fmt.Printf("[*] Angel found %d available site(s) ...\n", len(s))
				data := common.CollectData(s, "")
				common.PrintData(data)
			} else {
				s := []common.ScrapingEx{
					&jtorrent.Nyaa{},
					&jtorrent.SuKeBe{},
				}
				s = common.GetAvailableSitesEx(s)
				fmt.Printf("[*] Angel found %d available site(s) ...\n", len(s))
				data := common.CollectDataEx(s, "")
				common.PrintDataEx(data)
			}
			return nil
		},
		// Use the default error handling for CLI, which avoids the redefined error
	}

	// Run the app
	err := app.Run(os.Args)
	if err != nil {
		// Log errors
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
