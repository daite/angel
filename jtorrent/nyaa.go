package jtorrent

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/daite/angel/common"
)

// Client struct is for sending tasks
type Client struct {
	title string
	link  string
}

// Data struct is for receiving final data
type Data struct {
	title  string
	magnet string
}

var (
	clients = make(chan Client, 100)
	data    = make(chan Data, 100)
)

func create(doc *goquery.Document, baseURL string) {
	doc.Find("a[href*=view]:last-child").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")
		link = baseURL + link
		c := Client{title, link}
		clients <- c
	})
	close(clients)
}

// Nyaa struct is for Nyaa torrent web site
type Nyaa struct {
	Name        string
	Keyword     string
	SearchURL   string
	ScrapedData map[string]string
}

// initialize method set keyword and URL based on default url
func (n *Nyaa) initialize(keyword string) {
	n.Keyword = keyword
	n.Name = "nyaa"
	n.SearchURL = common.TorrentURL[n.Name] + "/?f=0&c=0_0&q=" + keyword
}

// Crawl torrent data from web site
// NOTE: status code error: 429 429 Too Many Requests for goroutines
// Max concurrent request: 5
func (n *Nyaa) Crawl(keyword string) map[string]string {
	n.initialize(keyword)
	fmt.Printf("[*] %s starts Crawl!!\n", n.Name)
	data := n.getData(n.SearchURL)
	return data
}

// GetData method returns map(title, bbs url)
func (n *Nyaa) getData(url string) map[string]string {
	resp := common.GetResponseFromURL(url)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	go create(doc, common.TorrentURL[n.Name])
	n.makeWP(5)
	m := make(map[string]string, 0)
	for d := range data {
		m[d.title] = d.magnet
	}
	n.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (n *Nyaa) GetMagnet(url string) string {
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

func (n *Nyaa) worker(wg *sync.WaitGroup) {
	for c := range clients {
		title := c.title
		magnet := n.GetMagnet(c.link)
		data <- Data{title, magnet}
	}
	wg.Done()
}

func (n *Nyaa) makeWP(num int) {
	// Max concurrent should be "5"
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go n.worker(&wg)
	}
	wg.Wait()
	close(data)
}
