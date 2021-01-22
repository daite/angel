package angel

import (
	"reflect"
	"testing"
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
	bbsURL := TTOBoGoURL + "/post/150327"
	got := tbg.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:167f61c08c130b4e9ddff003af81ecb6177c47b8"
	if got != want {
		t.Errorf("GetMagnet() for TToBogo = %q, want %q", got, want)
	}
}

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
	bbsURL := TorrentMobileURL + "/bbs/board.php?bo_table=movie&wr_id=17220"
	got := tmb.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:baeffe526ecb61e2e774b2e460a5bdddf3f1e195"
	if got != want {
		t.Errorf("GetMagnet() for TorrentMobile = %q, want %q", got, want)
	}
}

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
	bbsURL := TorrentViewURL + "/bbs/board.php?bo_table=mov&wr_id=17725"
	got := tv.GetMagnet(bbsURL)
	want := "magnet:?xt=urn:btih:baeffe526ecb61e2e774b2e460a5bdddf3f1e195"
	if got != want {
		t.Errorf("GetMagnet() for TorrentView = %q, want %q", got, want)
	}
}

func TestTorrentTubeCrawlFunc(t *testing.T) {
	tv := TorrentTube{}
	got := tv.Crawl("동상이몽2 너는 내운명 E179")
	want := map[string]string{
		"동상이몽2 너는 내운명.E179.210111.720p-NEXT": "magnet:?xt=urn:btih:2d5eee0743df5950f0193ca53017ecb99409b5c4",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Crawl() for TorrentTube = %q, want %q", got, want)
	}
}
