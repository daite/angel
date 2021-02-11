package ktorrent

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"

	"github.com/daite/angel/common"
)

// TorrentTube struct is for TorrentTube torrent web site
type TorrentTube struct {
	Name        string
	Keyword     string
	SearchURL   string
	ScrapedData map[string]string
}

// initialize method set keyword and URL based on default url
func (t *TorrentTube) initialize(keyword string) {
	t.Keyword = keyword
	t.Name = "torrenttube"
	t.SearchURL = common.TorrentURL[t.Name] + "/kt/search?p&q=" + url.QueryEscape(t.Keyword)
}

// Crawl torrent data from web site
func (t *TorrentTube) Crawl(keyword string) map[string]string {
	t.initialize(keyword)
	m := map[string]string{}
	resp, ok := common.GetResponseFromURL(t.SearchURL)
	if !ok {
		return nil
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	doc := string(b)
	re := regexp.MustCompile(`\[[^\]]*\]`)
	s := re.FindAllString(doc, -1)[1]
	jsonStr := strings.Replace(s, "'", `"`, -1)
	jsonMapArr := []map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonStr), &jsonMapArr)
	if err != nil {
		return nil
	}
	for _, d := range jsonMapArr {
		title := d["fn"].(string)
		magnet := "magnet:?xt=urn:btih:" + d["hs"].(string)
		m[title] = magnet
	}
	t.ScrapedData = m
	return m
}
