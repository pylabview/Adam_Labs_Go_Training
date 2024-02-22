//go:build ignore

package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var feedURLs = []string{
	// BBC
	"https://feeds.bbci.co.uk/news/rss.xml",
	"https://feeds.bbci.co.uk/news/world/rss.xml",
	"https://feeds.bbci.co.uk/news/politics/rss.xml",
	"https://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml",

	// CNN
	"http://rss.cnn.com/rss/cnn_topstories.rss",
	"http://rss.cnn.com/rss/cnn_world.rss",
	"http://rss.cnn.com/rss/cnn_us.rss",
	"http://rss.cnn.com/rss/cnn_allpolitics.rss",

	// NYT
	"https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	"https://rss.nytimes.com/services/xml/rss/nyt/US.xml",
	"https://rss.nytimes.com/services/xml/rss/nyt/Politics.xml",
	"https://rss.nytimes.com/services/xml/rss/nyt/Business.xml",
}

type Item struct {
	Title   string
	Content string
	Link    string
}

func parseRSS(r io.Reader) ([]Item, error) {
	var doc struct {
		Items []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
		} `xml:"channel>item"`
	}

	if err := xml.NewDecoder(r).Decode(&doc); err != nil {
		return nil, err
	}

	items := make([]Item, len(doc.Items))
	for i, item := range doc.Items {
		items[i].Title = item.Title
		items[i].Content = item.Description
		items[i].Link = item.Link
	}

	return items, nil
}

func fetchFeed(ctx context.Context, url string) ([]Item, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%q - %s", url, resp.Status)
	}

	return parseRSS(resp.Body)
}

func fetchAllFeeds(ctx context.Context) ([]Item, error) {
	var allItems []Item
	for _, url := range feedURLs {
		items, err := fetchFeed(ctx, url)
		if err != nil {
			return nil, err
		}
		allItems = append(allItems, items...)
	}

	return allItems, nil
}

var header = `
package main

var dbData = []Record{
`

var itemText = `
{
	Title: %q,
	Content: %q,
	Link: %q,
},
`

func main() {
	log.Printf("INFO: fetching items")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	items, err := fetchAllFeeds(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	log.Printf("INFO: %d items", len(items))

	out, err := os.Create("db_data.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	defer out.Close()

	fmt.Fprint(out, header)
	for _, item := range items {
		fmt.Fprintf(out, itemText, item.Title, item.Content, item.Link)
	}
	fmt.Fprintf(out, "}")
}
