package ktorrent

import (
	"reflect"
	"testing"
)

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
