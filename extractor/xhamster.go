package extractor

import (
	"regexp"
	"fmt"
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

type xhamster struct {
	doc *html.Node
}

func (x *xhamster) Title() string {
	t := htmlquery.FindOne(x.doc, "//title/text()")
	if t != nil {
		return t.Data
	}
	return ""
}

func (x *xhamster) Thumbnails() []string {
	items := make([]map[int]int, 10)
	ls := []string{}
	re := regexp.MustCompile(`(.+\/)\d+(_\d+\.jpg)$`)

	for i := range items {
		r := fmt.Sprintf("${1}%d${2}", i+1)
		thumb := re.ReplaceAllString(x.Thumbnail(), r)
		ls = append(ls, thumb)
	}

	return ls
}

func (x *xhamster) Thumbnail() string {
	t := htmlquery.FindOne(x.doc, `//link[@itemprop="thumbnailUrl"]`)
	return htmlquery.SelectAttr(t,  "href" )
}


func (x *xhamster) ProviderUrl() string {
	return "https://xhamster.com"
}

func (x *xhamster) ProviderName() string {
	return "xhamster"
}

func (x *xhamster) ProviderDisplay() string {
	return "xhamster"
}

func (x *xhamster) Duration() string {
	return ""
}

func (x *xhamster) Keywords() []string {
	items := []string{}

	expr := `//td[@id='channels']//a/text()`
	htmlquery.FindEach(x.doc, expr, func(index int, n *html.Node) {
		items = append(items, n.Data)
	})
	return items
}
