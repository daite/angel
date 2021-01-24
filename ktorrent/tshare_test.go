package ktorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestTShareCrawlFunc(t *testing.T) {
	ts := TShare{}
	got := ts.Crawl("처제길들이기")
	want := map[string]string{
		"처제길들이기 2020.720p.HDRip.H264.AAC.mkv": "magnet:?xt=urn:btih:890ef99f886552c2f7d6b1b237509856dc063803",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for TShare = %q, want %q", got, want)
	}
}

func TestTShareMagnetFunc(t *testing.T) {
	ts := TShare{}
	bbsURL := common.TorrentURL["tshare"] + "/movie/11457"
	got := ts.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:890ef99f886552c2f7d6b1b237509856dc063803"
	if got != want {
		t.Errorf("GetMagnet() for TShare = %q, want %q", got, want)
	}
}
