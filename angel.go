package angel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Scraping interface is for web scraping
type Scraping interface {
	Crawl(string) map[string]string
}

// TToBoGo struct is for ttobogo torrent web site
type TToBoGo struct {
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// initialize function set keyword and URL based on default url
func (t *TToBoGo) initialize(keyword string) {
	t.Keyword = keyword
	t.SearchURL = TTOBoGoURL + "/search?skeyword=" + keyword
}

// Crawl torrent data from web site
func (t *TToBoGo) Crawl(keyword string) map[string]string {
	fmt.Println("[*] TToBoGo starts Crawl!!")
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
func (t *TToBoGo) getData(url string) *sync.Map {
	var wg sync.WaitGroup
	m := &sync.Map{}
	resp := GetResponseFromURL(url)
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
	resp := GetResponseFromURL(url)
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

// TorrentMobile struct is for TorrentMobile torrent web site
type TorrentMobile struct {
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// initialize function set keyword and URL based on default url
func (t *TorrentMobile) initialize(keyword string) {
	t.Keyword = keyword
	t.SearchURL = TorrentMobileURL + "/bbs/search.php?&stx=" + keyword
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
	resp := GetResponseFromURL(url)
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
			link = strings.TrimSpace(urlJoin(TorrentMobileURL, link))
			m.Store(title, t.GetMagnet(link))
		}()
	})
	wg.Wait()
	t.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (t *TorrentMobile) GetMagnet(url string) string {
	resp := GetResponseFromURL(url)
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

// TorrentView struct is for TorrentView torrent web site
// It is exactly the same with torrentmobile
type TorrentView struct {
	Keyword     string
	SearchURL   string
	ScrapedData *sync.Map
}

// initialize function set keyword and URL based on default url
func (t *TorrentView) initialize(keyword string) {
	t.Keyword = keyword
	t.SearchURL = TorrentViewURL + "/bbs/search.php?&stx=" + keyword
}

// Crawl torrent data from web site
func (t *TorrentView) Crawl(keyword string) map[string]string {
	fmt.Println("[*] TorrentView starts Crawl!!")
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
func (t *TorrentView) getData(url string) *sync.Map {
	var wg sync.WaitGroup
	m := &sync.Map{}
	resp := GetResponseFromURL(url)
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
			link = strings.TrimSpace(urlJoin(TorrentViewURL, link))
			m.Store(title, t.GetMagnet(link))
		}()
	})
	wg.Wait()
	t.ScrapedData = m
	return m
}

// GetMagnet method returns torrent magnet
func (t *TorrentView) GetMagnet(url string) string {
	resp := GetResponseFromURL(url)
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

// TorrentTube struct is for TorrentView torrent web site
// It is exactly the same with torrentmobile
type TorrentTube struct {
	Keyword     string
	SearchURL   string
	ScrapedData map[string]string
}

// initialize function set keyword and URL based on default url
func (t *TorrentTube) initialize(keyword string) {
	t.Keyword = keyword
	t.SearchURL = TorrentTubeURL + "/kt/search?p&q=" + keyword
}

// Crawl torrent data from web site
func (t *TorrentTube) Crawl(keyword string) map[string]string {
	fmt.Println("[*] TorrentTube starts Crawl!!")
	t.initialize(keyword)
	m := map[string]string{}
	resp := GetResponseFromURL(t.SearchURL)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	doc := string(b)
	re := regexp.MustCompile(`\[[^\]]*\]`)
	s := re.FindAllString(doc, -1)[1]
	jsonStr := strings.Replace(s, "'", `"`, -1)
	jsonMapArr := []map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonStr), &jsonMapArr)
	if err != nil {
		log.Fatalln(err)
	}
	for _, d := range jsonMapArr {
		title := d["fn"].(string)
		magnet := "magnet:?xt=urn:btih:" + d["hs"].(string)
		m[title] = magnet
	}
	return m
}
