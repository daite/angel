package ktorrent

import (
	"testing"

	"github.com/daite/angel/common"
)

func TestTorrentSeeCrawlFunc(t *testing.T) {
	ts := TorrentSee{}
	got := ts.Crawl("광서열차 2021")
	title := "광서열차 2021.1080p.KOR.FHDRip.H264.AAC-JTC"
	want := map[string]string{
		title: "magnet:?xt=urn:btih:80788dd173e48e5eb139758c165a89c3c048d458",
	}
	if got[title] != want[title] {
		t.Errorf("Crawl() for TorrentSee = %q, want %q", got, want)
	}
}

func TestTorrentSeeMagnetFunc(t *testing.T) {
	ts := TorrentSee{}
	bbsURL := common.TorrentURL["torrentsee"] + "/topic/87009"
	got := ts.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:eee4d6fdf36ba112523cc48315ac5300cd84c77f"
	if got != want {
		t.Errorf("GetMagnet() for TorrentSee = %q, want %q", got, want)
	}
}
