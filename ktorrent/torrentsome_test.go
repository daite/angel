package ktorrent

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentSome(t *testing.T) {
	f, err := os.Open("../resources/torrentsome_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find("div.flex-auto a").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			title, _ := s.Attr("title")
			title = strings.TrimSpace(title)
			link, _ := s.Attr("href")
			got[title] = link
		}
	})
	want := map[string]string{
		"동상이몽2 너는 내운명_E186_210301": "/v/123335",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for TorrentSome = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTorrentSome(t *testing.T) {
	f, err := os.Open("../resources/torrentsome_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	magnet, _ := doc.Find("a.ml-3").Attr("href")
	got := strings.TrimSpace(magnet)
	want := "magnet:?xt=urn:btih:08a1b53bcb809a94c2ce9582bb4d70b6a4ad4460"
	if got != want {
		t.Errorf("GetMagnet() for TorrentSome = %q, want %q", got, want)
	}
}
