package tests

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDataFuncForTorrentMax(t *testing.T) {
	f, err := os.Open("../resources/torrentmax_search.html")
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
		"동상이몽2 너는 내운명.E176.201221.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/22523",
		"동상이몽2 너는 내운명.E177.201228.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/22715",
		"동상이몽2 너는 내운명.E177.201228.720p-NEXT.mp4": "https://torrentmax15.com/max/VARIETY/22718",
		"동상이몽2 너는 내운명.E178.210104.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/22916",
		"동상이몽2 너는 내운명.E179.210111.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/23087",
		"동상이몽2 너는 내운명.E179.210111.720p-NEXT.mp4": "https://torrentmax15.com/max/VARIETY/23083",
		"동상이몽2 너는 내운명.E180.210118.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/23268",
		"동상이몽2 너는 내운명.E181.210125.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/23443",
		"동상이몽2 너는 내운명.E182.210201.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/23613",
		"동상이몽2 너는 내운명.E183.210208.720p-NEXT":     "https://torrentmax15.com/max/VARIETY/23713",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetData() for TorrentMax = %q, want %q", got, want)
	}
}
func TestGetMagnetFuncForTorrentMax(t *testing.T) {
	f, err := os.Open("../resources/torrentmax_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := strings.TrimSpace(doc.Find("ul.list-group").Text())
	want := "magnet:?xt=urn:btih:cbed3a226963bba284cc056a4ee2e1257ff71725"
	if got != want {
		t.Errorf("GetMagnet() for TorrentMax = %q, want %q", got, want)
	}
}
