package ktorrent

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/daite/angel/common"
)

// TToBoGo struct is for ttobogo torrent web site
type TToBoGo struct {
	Name        string
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// initialize method set keyword and URL based on default url
func (t *TToBoGo) initialize(keyword string) {
	t.Keyword = keyword
	t.Name = "ttobogo"
	t.SearchURL = common.TorrentURL[t.Name] + "/search?skeyword=" + keyword
}

// Crawl torrent data from web site
func (t *TToBoGo) Crawl(keyword string) map[string]string {
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
func (t *TToBoGo) getData(url string) *sync.Map {
	var wg sync.WaitGroup
	m := &sync.Map{}
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find("a.subject").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			title := s.Text()
			link, _ := s.Attr("href")
			m.Store(title, t.GetMagnet(link))
		}()
	})
	wg.Wait()
	t.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (t *TToBoGo) GetMagnet(url string) string {
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	magnet, ok := doc.Find(".btn.btn-blue").Attr("onclick")
	if !ok {
		// maybe subtitles for movies
		return "NO MAGNET"
	}
	return strings.Split(magnet, "'")[1]
}
