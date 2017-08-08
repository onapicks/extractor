package extractor

import (
	"regexp"
	"fmt"
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

type redtube struct {
	doc *html.Node
}

func (x *redtube) Title() string {
	t := htmlquery.FindOne(x.doc, "//title/text()")
	if t != nil {
		return t.Data
	}
	return ""
}

func (x *redtube) Thumbnails() []string {
	items := make([]map[int]int, 16)

	ls := []string{}

	re := regexp.MustCompile(`(.+\/)\d+(.jpg)$`)

	for i := range items {
		r := fmt.Sprintf("${1}%d${2}", i+1)
		thumb := re.ReplaceAllString(x.Thumbnail(), r)
		ls = append(ls, thumb)
	}
	return ls
}

func (x *redtube) Thumbnail() string {
	t := htmlquery.FindOne(x.doc, "//meta[@property='og:image']")
	return htmlquery.SelectAttr(t,  "content" )
}


func (x *redtube) ProviderUrl() string {
	return "https://redtube.com"
}

func (x *redtube) ProviderName() string {
	return "redtube"
}

func (x *redtube) ProviderDisplay() string {
	return "redtube"
}

func (x *redtube) Duration() string {
	return ""
}

func (x *redtube) Keywords() []string {
	items := []string{}
	expr := `//span[@class='category-list-item-name']/text()`
	htmlquery.FindEach(x.doc, expr, func(index int, n *html.Node) {
		items = append(items, n.Data)
	})
	return items
}
