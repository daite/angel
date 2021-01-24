package ktorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestTorrentViewCrawlFunc(t *testing.T) {
	tv := TorrentView{}
	got := tv.Crawl("핫바디 처제")
	want := map[string]string{
		"핫바디 처제 2020.1080p.FHDRip.H264.AAC.mp4": "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for TorrentView = %q, want %q", got, want)
	}
}

func TestTorrentViewMagnetFunc(t *testing.T) {
	tv := TorrentView{}
	bbsURL := common.TorrentViewURL + "/bbs/board.php?bo_table=mov&wr_id=17725"
	got := tv.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:baeffe526ecb61e2e774b2e460a5bdddf3f1e195"
	if got != want {
		t.Errorf("GetMagnet() for TorrentView = %q, want %q", got, want)
	}
}
