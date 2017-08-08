package extractor

import (
	"regexp"
	"fmt"
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

type pornhub struct {
	doc *html.Node
}

func (x *pornhub) Title() string {
	t := htmlquery.FindOne(x.doc, "//title/text()")
	if t != nil {
		return t.Data
	}
	return ""
}

func (x *pornhub) Thumbnails() []string {
	t := x.Thumbnail()

	if t == "" {
		return []string{}
	}

	items := make([]map[int]int, 16)

	ls := []string{}

	re := regexp.MustCompile(`(.+\))\d+(.jpg)$`)

	for i := range items {
		r := fmt.Sprintf("${1}%d${2}", i+1)
		thumb := re.ReplaceAllString(t, r)
		ls = append(ls, thumb)
	}
	return ls
}

func (x *pornhub) ProviderUrl() string {
	return "https://pornhub.com"
}

func (x *pornhub) ProviderName() string {
	return "pornhub"
}

func (x *pornhub) ProviderDisplay() string {
	return "pornhub"
}

func (x *pornhub) Duration() string {
	duration := htmlquery.FindOne(x.doc, `//div[@id="js-shareData"]`)
	n := htmlquery.SelectAttr(duration,  "data-duration" )
	return n
}

func (x *pornhub) Thumbnail() string {
	t := htmlquery.FindOne(x.doc, "//div[@class='thumbnail']/img")
	return htmlquery.SelectAttr(t,  "src" )
}

func (x *pornhub) Keywords() []string {
	items := []string{}
	expr := `//div[@class='categoriesWrapper']//a[starts-with(@href,'/video?c=')]/text()`
	htmlquery.FindEach(x.doc, expr, func(index int, n *html.Node) {
		items = append(items, n.Data)
	})
	return items
}
