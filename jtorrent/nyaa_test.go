package jtorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestNyaaCrawlFunc(t *testing.T) {
	n := Nyaa{}
	got := n.Crawl("nijiiro+EP01+Magic")
	want := map[string]string{
		"[MagicStar] Nijiiro Karte EP01 [WEBDL] [1080p]": "magnet:?xt=urn:btih:087858c2626987779f9a3e107e4d12607a6e66aa",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for Nyaa = %q, want %q", got, want)
	}
}

func TestNyaaMagnetFunc(t *testing.T) {
	n := Nyaa{}
	bbsURL := common.TorrentURL["nyaa"] + "/view/1331289"
	got := n.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:087858c2626987779f9a3e107e4d12607a6e66aa"
	if got != want {
		t.Errorf("GetMagnet() for Nyaa = %q, want %q", got, want)
	}
}
