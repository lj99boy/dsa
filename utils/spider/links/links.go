package links

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML:%v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for i := n.FirstChild; i != nil; i = i.NextSibling {
		forEachNode(i, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var tokens = make(chan int, 20)

func ReadWithError(url string) []string {
	//针对每个节点的处理代码
	println(url)

	//限流，防止并发数太高引起网络请求错误
	tokens <- 1
	res, err := Extract(url)
	<-tokens

	if err != nil {
		log.Fatal(err)
	}
	return res
}

func BreadthFirst(fn func(url string) []string, workList []string) {
	seen := make(map[string]bool, 0)

	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				workList = append(workList, fn(item)...)
			}
		}
	}
}

func ConcurrentRead(fn func(url string) []string, url string) {
	workList := make(chan []string)

	go func() {
		workList <- []string{url}
	}()

	seen := make(map[string]bool, 0)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				go func(link string) {
					workList <- fn(link)
				}(link)
			}
		}
	}
}
