package jtorrent

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/daite/angel/common"
)

// SClient struct is for sending tasks
type SClient struct {
	title string
	link  string
}

// SData struct is for receiving final data
type SData struct {
	title  string
	magnet string
}

var (
	sclients = make(chan SClient, 100)
	sdata    = make(chan SData, 100)
)

func screate(doc *goquery.Document, baseURL string) {
	doc.Find("a[href*=view]:last-child").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")
		link = baseURL + link
		c := SClient{title, link}
		sclients <- c
	})
	close(sclients)
}

// SuKeBe struct is for sukebei torrent web site
type SuKeBe struct {
	Name        string
	Keyword     string
	SearchURL   string
	ScrapedData map[string]string
}

// initialize method set keyword and URL based on default url
func (s *SuKeBe) initialize(keyword string) {
	s.Keyword = keyword
	s.Name = "sukebe"
	s.SearchURL = common.TorrentURL[s.Name] + "/?f=0&c=0_0&q=" + url.QueryEscape(s.Keyword)
}

// Crawl torrent data from web site
// NOTE: status code error: 429 429 Too Many Requests for goroutines
// Max concurrent request: 5
func (s *SuKeBe) Crawl(keyword string) map[string]string {
	s.initialize(keyword)
	fmt.Printf("[*] %s starts Crawl!!\n", s.Name)
	data := s.getData(s.SearchURL)
	return data
}

// GetData method returns map(title, bbs url)
func (s *SuKeBe) getData(url string) map[string]string {
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	go screate(doc, common.TorrentURL[s.Name])
	s.makeWP(5)
	m := make(map[string]string, 0)
	for d := range sdata {
		title := d.title
		if len(title) > 40 {
			title = title[:40]
		}
		m[title] = d.magnet
	}
	s.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (s *SuKeBe) GetMagnet(url string) string {
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

func (s *SuKeBe) worker(wg *sync.WaitGroup) {
	for c := range sclients {
		title := c.title
		magnet := s.GetMagnet(c.link)
		sdata <- SData{title, magnet}
	}
	wg.Done()
}

func (s *SuKeBe) makeWP(num int) {
	// Max concurrent should be "5"
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go s.worker(&wg)
	}
	wg.Wait()
	close(sdata)
}
