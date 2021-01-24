package common

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

// Scraping interface is for web scraping
type Scraping interface {
	Crawl(string) map[string]string
}

// variable for scraping
var (
	TTOBoGoURL       = "https://ttobogo.net"
	TorrentMobileURL = "https://torrentmobile16.com"
	TorrentViewURL   = "https://torrentview29.com"
	TorrentTubeURL   = "https://torrentube.to"
	TShareURL        = "https://tshare.org"
	NyaaURL          = "https://nyaa.si"
	UserAgent        = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0)"
)

// GetResponseFromURL returns *http.Response from url
func GetResponseFromURL(url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return resp
}

// CollectData function executes web scraping based on each scrapper
func CollectData(s []Scraping, keyword string) map[string]string {
	var wg sync.WaitGroup
	ch := make(chan map[string]string, 5)
	for _, i := range s {
		wg.Add(1)
		go func(v Scraping) {
			defer wg.Done()
			r := v.Crawl(keyword)
			ch <- r
		}(i)
	}
	wg.Wait()
	close(ch)
	m := map[string]string{}
	for elem := range ch {
		for k, v := range elem {
			k = strings.Replace(k, " ", "_", -1)
			m[k] = v
		}
	}
	return m
}

// PrintData function prints scraped data to console
func PrintData(data map[string]string, reverse bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Magnet"})
	matrix := [][]string{}
	for k, v := range data {
		matrix = append(matrix, []string{k, v})
	}
	sort.Slice(matrix[:], func(i, j int) bool {
		for x := range matrix[i] {
			if matrix[i][x] == matrix[j][x] {
				continue
			}
			if reverse == true {
				return matrix[i][x] < matrix[j][x]
			}
			return matrix[i][x] > matrix[j][x]
		}
		return false
	})
	for _, v := range matrix {
		table.Append(v)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

// URLJoin function join baseURL and relURL
func URLJoin(baseURL string, relURL string) string {
	u, err := url.Parse(relURL)
	if err != nil {
		log.Fatal(err)
	}
	base, err := url.Parse(baseURL + "/bbs/")
	if err != nil {
		log.Fatal(err)
	}
	return base.ResolveReference(u).String()
}

// CheckNetWorkFromURL function checks network status
func CheckNetWorkFromURL(url string) bool {
	client := &http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		return false
	}
	return true
}
