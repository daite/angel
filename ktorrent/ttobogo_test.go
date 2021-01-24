package ktorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestTToBoGoCrawlFunc(t *testing.T) {
	tbg := TToBoGo{}
	got := tbg.Crawl("핫바디")
	want := map[string]string{
		"핫바디 처제 2020.1080p.FHDRip.H264.AAC": "magnet:?xt=urn:btih:1cc7a302e8402c48a76962d6b8f15fa4aab70381",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for TToBogo = %q, want %q", got, want)
	}
}

func TestTToBoGoGetMagnetFunc(t *testing.T) {
	tbg := TToBoGo{}
	bbsURL := common.TorrentURL["ttobogo"] + "/post/150327"
	got := tbg.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:167f61c08c130b4e9ddff003af81ecb6177c47b8"
	if got != want {
		t.Errorf("GetMagnet() for TToBogo = %q, want %q", got, want)
	}
}
