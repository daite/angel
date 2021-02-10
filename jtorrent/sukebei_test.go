package jtorrent

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForSukeBei(t *testing.T) {
	f, err := os.Open("../resources/sukebei_search.html")
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
		title := strings.TrimSpace(s.Text())[:40]
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"+++ [HD] SIRO-4400 【初撮り】【婚": "/view/3236442",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for Sukebei = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForSukeBei(t *testing.T) {
	f, err := os.Open("../resources/sukebei_bbs.html")
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
	want := "magnet:?xt=urn:btih:92596bbdc0176f523508afbf99550247dba8b35f"
	if got != want {
		t.Errorf("GetMagnet() for Sukebei = %q, want %q", got, want)
	}
}
