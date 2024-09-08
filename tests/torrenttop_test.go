package tests

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentTop(t *testing.T) {
	f, err := os.Open("../resources/torrenttop_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find(".py-4.flex.flex-row.border-b.topic-item a").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		title = strings.TrimSpace(title)
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"서진이네 2.E01.240628.720p-NEXT": "/torrent/vzla5mq.html",
	}
	if got["서진이네 2.E01.240628.720p-NEXT"] != "/torrent/vzla5mq.html" {
		t.Errorf("GetData() for TorrentTop = %q, want %q", got, want)
	}
}

func TestGetMagnetFuncForTorrentTop(t *testing.T) {
	f, err := os.Open("../resources/torrenttop_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got, _ := doc.Find(".fas.fa-magnet + a").Attr("href")
	want := "magnet:?xt=urn:btih:6cf65299b5b48b077370f5675ce34b666e82cc3f"
	if got != want {
		t.Errorf("GetMagnet() for TorrentTop = %q, want %q", got, want)
	}
}
