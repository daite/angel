package ktorrent

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentWiz(t *testing.T) {
	f, err := os.Open("../resources/torrentwiz_search.html")
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
		"핫바디 처제 2020.1080p.FHDRip.H264.AAC.mp4": "./board.php?bo_table=mov&wr_id=16948",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for torrentwiz = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTorrentWiz(t *testing.T) {
	f, err := os.Open("../resources/torrentwiz_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := strings.TrimSpace(doc.Find("ul.list-group").Text())
	want := "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381"
	if got != want {
		t.Errorf("GetMagnet() for torrentwiz = %q, want %q", got, want)
	}
}
