package tests

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
func TestGetInfoFuncForSukeBei(t *testing.T) {
	f, err := os.Open("../resources/sukebei_bbs.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	got := doc.Find("div.col-md-5").Map(func(i int, s *goquery.Selection) string {
		return strings.TrimSpace(s.Text())
	})
	want := "9801ef1cf9ad6a3dd788d13df45471dbf2a29271"
	if got[8] != want {
		t.Errorf("GetInfo() for Sukebei = %q, want %q", got, want)
	}
	_, ok := doc.Find("a.folder").Attr("class")
	if !ok {
		t.Errorf("GetInfo() for Sukebei = %q, want %q", got, want)
	}
}
