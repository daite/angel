package tests

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentMobile(t *testing.T) {
	f, err := os.Open("../resources/torrentmobile_search.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := make(map[string]string)
	doc.Find("div.media-heading a").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")
		got[title] = link
	})
	want := map[string]string{
		"O형수박가슴가정부 2020.720p.HDRip.H264.AAC.mp4": "./board.php?bo_table=movie&wr_id=17220",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for TorrentMobile = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTorrentMobile(t *testing.T) {
	f, err := os.Open("../resources/torrentmobile_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := strings.TrimSpace(doc.Find("ul.list-group").Text())
	want := "magnet:?xt=urn:btih:baeffe526ecb61e2e774b2e460a5bdddf3f1e195"
	if got != want {
		t.Errorf("GetMagnet() for TorrentMobile = %q, want %q", got, want)
	}
}
