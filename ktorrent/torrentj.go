package ktorrent

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/daite/angel/common"
)

// TorrentJ struct is for TorrentJ torrent web site
// It is exactly the same with torrentmobile
type TorrentJ struct {
	Name        string
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// initialize method set keyword and URL based on default url
func (t *TorrentJ) initialize(keyword string) {
	t.Keyword = keyword
	t.Name = "torrentj"
	t.SearchURL = common.TorrentURL[t.Name] + "/bbs/search.php?&stx=" + keyword
}

// Crawl torrent data from web site
func (t *TorrentJ) Crawl(keyword string) map[string]string {
	t.initialize(keyword)
	fmt.Printf("[*] %s starts Crawl!!\n", t.Name)
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
func (t *TorrentJ) getData(url string) *sync.Map {
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
			link = strings.TrimSpace(common.URLJoin(common.TorrentURL["torrentj"], link))
			m.Store(title, t.GetMagnet(link))
		}()
	})
	wg.Wait()
	t.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (t *TorrentJ) GetMagnet(url string) string {
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
