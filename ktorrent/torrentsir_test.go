package ktorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestTorrentSirCrawlFunc(t *testing.T) {
	ts := TorrentSir{}
	got := ts.Crawl("핫바디 처제")
	want := map[string]string{
		"핫바디 처제 2020.1080p.FHDRip.H264.AAC": "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for TorrentSir = %q, want %q", got, want)
	}
}

func TestTorrentSirMagnetFunc(t *testing.T) {
	ts := TorrentSir{}
	bbsURL := common.TorrentURL["torrentsir"] + "/bbs/board.php?bo_table=movie&wr_id=15338"
	got := ts.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381"
	if got != want {
		t.Errorf("GetMagnet() for TorrentSir = %q, want %q", got, want)
	}
}
