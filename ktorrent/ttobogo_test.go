package ktorrent

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTToBoGo(t *testing.T) {
	f, err := os.Open("../resources/ttobogo_search.html")
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
		"핫바디 처제 2020.1080p.FHDRip.H264.AAC": "https://www1.ttobogo.net/post/158930",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for TToBoGo = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTToBoGo(t *testing.T) {
	f, err := os.Open("../resources/ttobogo_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	magnet, _ := doc.Find(".btn.btn-blue").Attr("onclick")
	got := strings.Split(magnet, "'")[1]
	want := "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381"
	if got != want {
		t.Errorf("GetMagnet() for TToBoGo = %q, want %q", got, want)
	}
}
