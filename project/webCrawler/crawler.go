package main

// Big thanks to the auditor of https://jdanger.com/build-a-web-crawler-in-go.html
import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/jackdanger/collectlinks"
)

var visited = map[string]bool{
	"javascript:void(0);": true,
}

func main() {

	// input argument input the url of start page
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specific a start page")
		os.Exit(1)
	}

	queue := make(chan string)

	go func(uri string) {
		queue <- uri
	}(args[0])

	for uri := range queue {
		if len(visited) > 1000 {
			break
		}
		enqueue(uri, queue)
	}
	println(visited)
	// visit page and print content of the page
	// fmt.Println(fetchURLsFromURL(args[0]))

}

func enqueue(uri string, queue chan string) {
	visited[uri] = true
	fmt.Println("fetching", uri)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	// get the address of new created config

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}
	subURIs := fetchURLsFromURI(uri, client)

	for _, subURI := range subURIs {
		absolute := fixURL(subURI, uri)
		if uri != "" && !visited[absolute] {
			go func(link string) {
				queue <- link
			}(absolute)
		}

	}
}

func fixURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		// fmt.Println("invalid1")
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		// fmt.Println("invalid2")
		return ""
	}
	uri = baseURL.ResolveReference(uri)
	// fmt.Println("valid url:", uri)

	return uri.String()
}

func fetchURLsFromURI(uri string, c http.Client) []string {
	resp, _ := c.Get(uri)
	// body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return collectlinks.All(resp.Body)
}
