package jtorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestSukeBeCrawlFunc(t *testing.T) {
	n := SuKeBe{}
	got := n.Crawl("+++ [HD] SIRO-4400")
	want := map[string]string{"+++ [HD] SIRO-4400 【初撮り】【婚": "magnet:?xt=urn:btih:92596bbdc0176f523508afbf99550247dba8b35f"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for SuKbeBe = %q, want %q", got, want)
	}
}

func TestSuKBeMagnetFunc(t *testing.T) {
	n := Nyaa{}
	bbsURL := common.TorrentURL["sukebe"] + "/view/3236442"
	got := n.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:92596bbdc0176f523508afbf99550247dba8b35f"
	if got != want {
		t.Errorf("GetMagnet() for Sukebe = %q, want %q", got, want)
	}
}
