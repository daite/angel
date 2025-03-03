package common

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
)

// Scraping interface is for web scraping
type Scraping interface {
	Crawl(string) map[string]string
}

// ScrapingEx interface is for web scraping
type ScrapingEx interface {
	Crawl(string) map[string][]string
}

// Color constants for progress output
const (
	Reset = "\033[0m"
	Green = "\033[32m"
	Red   = "\033[31m"
)

// variable for scraping
var (
	TorrentURL = map[string]string{
		"nyaa":        "https://nyaa.si",
		"sukebe":      "https://sukebei.nyaa.si",
		"torrentsee":  "torrentsee",
		"torrentqq":   "torrentqq",
		"torrentsome": "torrentsome",
		"torrentrj":   "torrentrj",
		"torrenttop":  "torrenttop",
	}
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"
)

// GetResponseFromURL returns *http.Response from url
func GetResponseFromURL(url string) (resp *http.Response, ok bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resp, false
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err = client.Do(req)
	if err != nil {
		return resp, false
	}
	if resp.StatusCode != 200 {
		return resp, false
	}
	return resp, true
}

// CollectData function executes web scraping based on each scrapper
func CollectData(s []Scraping, keyword string) map[string]string {
	fmt.Println("[*] Angel is collecting data ...")
	var wg sync.WaitGroup
	ch := make(chan map[string]string, len(TorrentURL))
	for _, i := range s {
		wg.Add(1)
		go func(v Scraping) {
			defer wg.Done()
			r := v.Crawl(keyword)
			if r == nil {
				return
			}
			ch <- r
		}(i)
	}
	wg.Wait()
	close(ch)
	m := map[string]string{}
	for elem := range ch {
		for k, v := range elem {
			k = strings.Replace(k, " ", "_", -1)
			if v == "no magnet" {
				continue
			}
			m[k] = v
		}
	}
	return m
}

// CollectDataEx function executes web scraping based on each scrapper
func CollectDataEx(s []ScrapingEx, keyword string) map[string][]string {
	fmt.Println("[*] Angel is collecting data ...")
	var wg sync.WaitGroup
	ch := make(chan map[string][]string, len(TorrentURL))
	for _, i := range s {
		wg.Add(1)
		go func(v ScrapingEx) {
			defer wg.Done()
			r := v.Crawl(keyword)
			if r == nil {
				return
			}
			ch <- r
		}(i)
	}
	wg.Wait()
	close(ch)
	m := map[string][]string{}
	for elem := range ch {
		for k, v := range elem {
			k = strings.Replace(k, " ", "_", -1)
			m[k] = v
		}
	}
	return m
}

// PrintData function prints scraped data to console
func PrintData(data map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Magnet"})
	matrix := [][]string{}
	for k, v := range data {
		matrix = append(matrix, []string{k, v})
	}
	sort.SliceStable(matrix, func(i, j int) bool { return matrix[i][0] > matrix[j][0] })
	for _, v := range matrix {
		table.Append(v)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

// PrintDataEx function prints scraped data to console
func PrintDataEx(data map[string][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Title", "Uploader", "Seeder", "Leecher",
		"Snatch", "FileSize", "Magnet", "Folder",
	})
	for k, v := range data {
		m := make([]string, 0)
		m = append(m, k)
		for _, i := range v {
			m = append(m, i)
		}
		table.Append(m)
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
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

// GetAvailableSites function gets available torrent sites
func GetAvailableSites(oldItems []Scraping) []Scraping {
	fmt.Println("[*] Angel is checking available torrent sites ...")
	newItems := make([]Scraping, 0)
	items := []string{
		"torrentsee", "torrentqq", "torrentsome", "torrentrj", "torrenttop",
	}
	ch := make(chan int, len(items))
	var wg sync.WaitGroup
	for n, title := range items {
		wg.Add(1)
		go func(i int, t string) {
			defer wg.Done()
			ok := CheckNetWorkFromURL(TorrentURL[t])
			if ok {
				ch <- i
			}
		}(n, title)
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		newItems = append(newItems, oldItems[v])
	}
	return newItems
}

// GetAvailableSitesEx function gets available torrent sites
func GetAvailableSitesEx(oldItems []ScrapingEx) []ScrapingEx {
	fmt.Println("[*] Angel is checking available torrent sites ...")
	newItems := make([]ScrapingEx, 0)
	items := []string{"nyaa", "sukebe"}
	ch := make(chan int, len(items))
	var wg sync.WaitGroup
	for n, title := range items {
		wg.Add(1)
		go func(i int, t string) {
			defer wg.Done()
			ok := CheckNetWorkFromURL(TorrentURL[t])
			if ok {
				ch <- i
			}
		}(n, title)
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		newItems = append(newItems, oldItems[v])
	}
	return newItems
}

// RemoveNonAscII remove non-ASCII characters
func RemoveNonAscII(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

// ------------------------------------------------------------------------------
// New functions to scrape torrent URLs from an HTML table and update TorrentURL
// with alignment for Korean site names and a progress indicator.
// ------------------------------------------------------------------------------

// FetchTorrentURLsFromHTML fetches the HTML page at scrapeURL, parses the table rows,
// and returns a map of raw site names (in Korean) to their URLs.
func FetchTorrentURLsFromHTML(scrapeURL string) (map[string]string, error) {
	resp, err := http.Get(scrapeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}
	torrentURLs := make(map[string]string)
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() >= 3 {
			rawSiteName := strings.TrimSpace(tds.Eq(1).Text())
			link, exists := tds.Eq(2).Find("a").Attr("href")
			if exists && rawSiteName != "" && link != "" {
				torrentURLs[rawSiteName] = link
			}
		}
	})
	return torrentURLs, nil
}

// alignSiteName maps the Korean site name to the internal English key.
func alignSiteName(koreanName string) string {
	switch koreanName {
	case "토렌트씨":
		return "torrentsee"
	case "토렌트큐큐":
		return "torrentqq"
	case "토렌트탑":
		return "torrenttop"
	case "토렌트알지":
		return "torrentrj"
	case "토렌트썸":
		return "torrentsome"
	default:
		return ""
	}
}

// generateBar returns a string progress bar given a percentage.
func generateBar(percentage int) string {
	bar := ""
	for i := 0; i < 50; i++ {
		if i < (percentage / 2) {
			bar += "="
		} else {
			bar += " "
		}
	}
	return bar
}

// displayProgress prints the progress of URL updates.
func displayProgress(current, total, successes, failures int) {
	progress := (current * 100) / total
	fmt.Printf("\r[*] Updating torrent URLs: %d/%d [%s] %d%% Successes: %s%d%s Failures: %s%d%s",
		current, total, generateBar(progress), progress, Green, successes, Reset, Red, failures, Reset)
}

// UpdateTorrentURLsFromHTMLWithProgress fetches torrent URLs from the provided scrapeURL,
// aligns the Korean site names to internal keys, and updates the global TorrentURL map
// while displaying progress.
func UpdateTorrentURLsFromHTMLWithProgress(scrapeURL string) error {
	newURLs, err := FetchTorrentURLsFromHTML(scrapeURL)
	if err != nil {
		return err
	}
	var alignedUpdates []struct {
		key string
		url string
	}
	for rawName, url := range newURLs {
		alignedKey := alignSiteName(rawName)
		if alignedKey != "" {
			alignedUpdates = append(alignedUpdates, struct {
				key string
				url string
			}{alignedKey, url})
		}
	}
	total := len(alignedUpdates)
	if total == 0 {
		log.Println("No matching torrent URLs found for update.")
		return nil
	}
	successes := 0
	failures := 0
	for i, update := range alignedUpdates {
		TorrentURL[update.key] = update.url
		successes++
		displayProgress(i+1, total, successes, failures)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\n[*] TorrentURL map has been updated from HTML scraping with progress.")
	return nil
}
