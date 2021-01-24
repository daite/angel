package ktorrent

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/daite/angel/common"
)

// Nyaa struct is for ttobogo torrent web site
type Nyaa struct {
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// initialize method set keyword and URL based on default url
func (t *Nyaa) initialize(keyword string) {
	t.Keyword = keyword
	t.SearchURL = common.NyaaURL + "/?f=0&c=4_4&q=" + keyword
}

// Crawl torrent data from web site
// NOTE: status code error: 429 429 Too Many Requests for goroutines
func (t *Nyaa) Crawl(keyword string) map[string]string {
	fmt.Println("[*] Nyaa starts Crawl!!")
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
func (t *Nyaa) getData(url string) *sync.Map {
	var wg sync.WaitGroup
	m := &sync.Map{}
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find("a[href*=view]:last-child").Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			title := strings.TrimSpace(s.Text())
			link, _ := s.Attr("href")
			link = common.NyaaURL + link
			m.Store(title, t.GetMagnet(link))
		}()
	})
	wg.Wait()
	t.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (t *Nyaa) GetMagnet(url string) string {
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	magnet, ok := doc.Find("a.card-footer-item").Attr("href")
	if !ok {
		// maybe subtitles for movies
		return "NO MAGNET"
	}
	magnet = strings.Split(magnet, "&")[0]
	return magnet
}
