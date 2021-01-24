package ktorrent

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/daite/angel/common"
)

// TShare struct is for ttobogo torrent web site
type TShare struct {
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// Initialize method set keyword and URL based on default url
func (t *TShare) initialize(keyword string) {
	t.Keyword = keyword
	t.SearchURL = common.TorrentURL["tshare"] + "/bbs/search.php?sfl=wr_content&stx=" + keyword
}

// Crawl torrent data from web site
func (t *TShare) Crawl(keyword string) map[string]string {
	fmt.Println("[*] TShare starts Crawl!!")
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
func (t *TShare) getData(url string) *sync.Map {
	var wg sync.WaitGroup
	m := &sync.Map{}
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find("li.list-item-row a").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			title := strings.TrimSpace(s.Find("h1").Text())
			link, _ := s.Attr("href")
			m.Store(title, t.GetMagnet(link))
		}()
	})
	wg.Wait()
	t.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (t *TShare) GetMagnet(url string) string {
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	magnet, ok := doc.Find("td a").Attr("href")
	if !ok {
		// maybe subtitles for movies
		return "NO MAGNET"
	}
	return magnet
}
