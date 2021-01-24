package ktorrent

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/daite/angel/common"
)

// TorrentMobile struct is for TorrentMobile torrent web site
type TorrentMobile struct {
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// initialize method set keyword and URL based on default url
func (t *TorrentMobile) initialize(keyword string) {
	t.Keyword = keyword
	t.SearchURL = common.TorrentURL["torrentmobile"] + "/bbs/search.php?&stx=" + keyword
}

// Crawl torrent data from web site
func (t *TorrentMobile) Crawl(keyword string) map[string]string {
	fmt.Println("[*] TorrentMobile starts Crawl!!")
	t.initialize(keyword)
	data := t.getData(t.SearchURL)
	m := map[string]string{}
	data.Range(
		func(key, value interface{}) bool {
			m[fmt.Sprint(key)] = fmt.Sprint(value)
			return true
		})
	return m
}

// GetData method returns map(title, bbs url)
func (t *TorrentMobile) getData(url string) *sync.Map {
	var wg sync.WaitGroup
	m := &sync.Map{}
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find("div.media-heading a").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			title := strings.TrimSpace(s.Text())
			link, _ := s.Attr("href")
			link = strings.TrimSpace(common.URLJoin(common.TorrentURL["torrentmobile"], link))
			m.Store(title, t.GetMagnet(link))
		}()
	})
	wg.Wait()
	t.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (t *TorrentMobile) GetMagnet(url string) string {
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	magnet := strings.TrimSpace(doc.Find("ul.list-group").Text())
	if magnet == "" {
		return "NO MAGNET"
	}
	return magnet
}
