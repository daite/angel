package common

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

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

// ANSI escape codes for colored output
const (
	Reset = "\033[0m"
	Green = "\033[32m"
	Red   = "\033[31m"
)

// variable for scraping
var (
	TorrentURL = map[string]string{
		"torrentmobile": "https://torrentmobile100.com",
		"torrentview":   "https://torrentview39.com",
		"nyaa":          "https://nyaa.si",
		"sukebe":        "https://sukebei.nyaa.si",
		"torrentsir":    "https://torrentsir86.com",
		"torrentj":      "https://torrentj84.com",
		"torrentsee":    "https://torrentsee236.com",
		"jujutorrent":   "https://torrentjuju14.com",
		"torrentqq":     "https://torrentqq323.com",
		"torrentwiz":    "https://torrentwiz48.com",
		"torrentgram":   "https://torrentgram47.com",
		"torrentsome":   "https://torrentsome150.com",
		"ktxtorrent":    "https://ktxtorrent37.com",
		"torrentrj":     "https://torrentrj156.com",
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
		"torrentmobile", "torrentview", "torrentsir",
		"torrentj", "torrentsee", "jujutorrent",
		"torrentqq", "torrentsome", "torrentrj",
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

/////////////////////////////////////////
// http://jaewook.net/archives/2613 /////
/////////////////////////////////////////

// Function to increment the URL number using regex
func incrementURL(url string) string {
	re := regexp.MustCompile(`(https://[a-zA-Z]+)(\d+)(\.(com|net))`)
	matches := re.FindStringSubmatch(url)
	if matches == nil {
		return url // Return the original URL if regex doesn't match
	}

	// Convert the number to an integer and increment it
	num, err := strconv.Atoi(matches[2])
	if err != nil {
		return url // Return the original URL if conversion fails
	}
	num++

	// Construct the new URL
	newURL := matches[1] + strconv.Itoa(num) + matches[3]
	return newURL
}

// Function to check the response code of the URL
func checkURL(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

// Function to update the URL with a maximum number of retries
func updateTorrentURL(key string, url string, maxRetries int, wg *sync.WaitGroup, resultChan chan<- struct {
	key string
	url string
}) {
	defer wg.Done()
	for i := 0; i < maxRetries; i++ {
		if checkURL(url) {
			resultChan <- struct {
				key string
				url string
			}{key, url}
			return
		}
		url = incrementURL(url)
	}
	resultChan <- struct {
		key string
		url string
	}{key, ""} // Indicate failure with an empty string
}

// Function to display a progress bar
func displayProgress(current, total, successes, failures int) {
	progress := (current * 100) / total
	fmt.Printf("\r[*] angel is updating torrent URLs: %d/%d [%s] %d%% Successes: %s%d%s Failures: %s%d%s",
		current, total, generateBar(progress), progress, Green, successes, Reset, Red, failures, Reset)
}

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

func init() {
	total := len(TorrentURL)
	successes := 0
	failures := 0
	current := 0
	maxRetries := 10 // Define a maximum number of retries

	var wg sync.WaitGroup
	resultChan := make(chan struct {
		key string
		url string
	})

	for key, url := range TorrentURL {
		wg.Add(1)
		go updateTorrentURL(key, url, maxRetries, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		current++
		if result.url != "" {
			TorrentURL[result.key] = result.url
			successes++
		} else {
			//fmt.Printf("Failed to update URL for key %s\n", result.key)
			failures++
		}
		displayProgress(current, total, successes, failures)
	}
	fmt.Println("\n[*] Angel completed to update torrent URLs")
}
