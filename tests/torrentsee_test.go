package tests

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentSee(t *testing.T) {
	f, err := os.Open("../resources/torrentsee_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find("li.tit > a").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"광서열차 2021.1080p.KOR.FHDRip.H264.AAC-JTC": "/topic/106997",
	}
	if got["광서열차 2021.1080p.KOR.FHDRip.H264.AAC-JTC"] != "/topic/106997" {
		t.Errorf("GetData() for TorrentSee = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTorrentSee(t *testing.T) {
	f, err := os.Open("../resources/torrentsee_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := strings.TrimSpace(doc.Find("a[target].bbs_btn2").Text())
	want := "magnet:?xt=urn:btih:80788dd173e48e5eb139758c165a89c3c048d458"
	if got != want {
		t.Errorf("GetMagnet() for TorrentSee = %q, want %q", got, want)
	}
}
