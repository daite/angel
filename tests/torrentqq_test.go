package tests

import (
	"log"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentQQ(t *testing.T) {
	f, err := os.Open("../resources/torrentqq_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find("a.subject").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"온앤오프.E36.210316.720p-NEXT": "https://torrentqq78.com/torrent/med/417306.html",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for TorrentQQ = %q, want %q", got, want)
	}
}

func TestGetMagnetFuncForTorrentQQ(t *testing.T) {
	f, err := os.Open("../resources/torrentqq_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile("[0-9,a-z]{40}")
	got := "no magnet"
	doc.Find("tr > td").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			got = "magnet:?xt=urn:btih:" + re.FindString(s.Text())
		}
	})
	want := "magnet:?xt=urn:btih:e9322c31da47494a31c7f8312c92e7a50a973759"
	if got != want {
		t.Errorf("GetMagnet() for TorrentQQ = %q, want %q", got, want)
	}
}
