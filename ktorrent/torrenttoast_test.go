package ktorrent

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentToast(t *testing.T) {
	f, err := os.Open("../resources/torrenttoast_search.html")
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
		"O형수박가슴가정부 2020.720p.HDRip.H264.AAC.mp4": "./board.php?bo_table=movie&wr_id=7053",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for TorrentToast = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTorrentToast(t *testing.T) {
	f, err := os.Open("../resources/torrenttoast_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := "no magnet"
	doc.Find("a.list-group-item").Each(func(i int, s *goquery.Selection) {
		if m, _ := s.Attr("href"); strings.Contains(m, "magnet:?xt=urn:btih:") {
			got = m
		}
	})
	want := "magnet:?xt=urn:btih:baeffe526ecb61e2e774b2e460a5bdddf3f1e195"
	if got != want {
		t.Errorf("GetMagnet() for TorrentToast = %q, want %q", got, want)
	}
}
