package ktorrent

import (
	"reflect"
	"testing"

	"github.com/daite/angel/common"
)

func TestTorrentMobileCrawlFunc(t *testing.T) {
	tmb := TorrentMobile{}
	got := tmb.Crawl("수박")
	want := map[string]string{
		"O형수박가슴가정부 2020.720p.HDRip.H264.AAC.mp4": "magnet:?xt=urn:btih:baeffe526ecb61e2e774b2e460a5bdddf3f1e195",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for TorrentMobile = %q, want %q", got, want)
	}
}

func TestTorrentMobileMagnetFunc(t *testing.T) {
	tmb := TorrentMobile{}
	bbsURL := common.TorrentMobileURL + "/bbs/board.php?bo_table=movie&wr_id=17220"
	got := tmb.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:baeffe526ecb61e2e774b2e460a5bdddf3f1e195"
	if got != want {
		t.Errorf("GetMagnet() for TorrentMobile = %q, want %q", got, want)
	}
}
