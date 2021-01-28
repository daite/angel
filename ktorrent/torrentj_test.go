package ktorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestTorrentJCrawlFunc(t *testing.T) {
	tj := TorrentJ{}
	got := tj.Crawl("핫바디 처제")
	want := map[string]string{
		"핫바디 처제 2020.1080p.FHDRip.H264.AAC": "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for TorrentJ = %q, want %q", got, want)
	}
}

func TestTorrentJMagnetFunc(t *testing.T) {
	tj := TorrentJ{}
	bbsURL := common.TorrentURL["torrentj"] + "/bbs/board.php?bo_table=movie&wr_id=15340"
	got := tj.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381"
	if got != want {
		t.Errorf("GetMagnet() for TorrentJ = %q, want %q", got, want)
	}
}
