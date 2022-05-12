package ktorrent

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentRJ(t *testing.T) {
	f, err := os.Open("../resources/torrentrj_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find("a.tit").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"광서열차 2021.1080p.KOR.FHDRip.H264.AAC-JTC": "/v/106444",
	}
	if got["광서열차 2021.1080p.KOR.FHDRip.H264.AAC-JTC"] != "/v/106444" {
		t.Errorf("GetData() for TorrentRJ = %q, want %q", got, want)
	}
}

func TestGetMagnetFuncForTorrentRJ(t *testing.T) {
	f, err := os.Open("../resources/torrentrj_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got, _ := doc.Find("a.ml-3").Attr("href")
	want := "magnet:?xt=urn:btih:80788dd173e48e5eb139758c165a89c3c048d458"
	if got != want {
		t.Errorf("GetMagnet() for TorrentRJ = %q, want %q", got, want)
	}
}
