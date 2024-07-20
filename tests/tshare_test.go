package tests

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTShare(t *testing.T) {
	f, err := os.Open("../resources/tshare_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find("li.list-item-row a").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h1").Text())
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"처제길들이기 2020.720p.HDRip.H264.AAC.mkv": "https://tshare.org/movie/11457",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for TShare = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTShare(t *testing.T) {
	f, err := os.Open("../resources/tshare_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got, _ := doc.Find("td a").Attr("href")
	want := "magnet:?xt=urn:btih:890ef99f886552c2f7d6b1b237509856dc063803"
	if got != want {
		t.Errorf("GetMagnet() for TShare = %q, want %q", got, want)
	}
}
