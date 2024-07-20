package tests

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForNyaa(t *testing.T) {
	f, err := os.Open("../resources/nyaa_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find("a[href*=view]:last-child").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"[MagicStar] Nijiiro Karte EP01 [WEBDL] [1080p]": "/view/1331289",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for Nyaa = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForNyaa(t *testing.T) {
	f, err := os.Open("../resources/nyaa_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	magnet, _ := doc.Find("a.card-footer-item").Attr("href")
	got := strings.Split(magnet, "&")[0]
	want := "magnet:?xt=urn:btih:087858c2626987779f9a3e107e4d12607a6e66aa"
	if got != want {
		t.Errorf("GetMagnet() for Nyaa = %q, want %q", got, want)
	}
}
