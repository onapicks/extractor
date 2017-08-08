package extractor

import (
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

type dmm struct {
	doc *html.Node
}

func (x *dmm) Title() string {
	t := htmlquery.FindOne(x.doc, "//title/text()")
	if t != nil {
		return t.Data
	}
	return ""
}

func (x *dmm) Thumbnails() []string {
	items := []string{}

	items = append(items, x.Thumbnail())

	htmlquery.FindEach(x.doc, `//a[@name="sample-image"]/img`, func(index int, n *html.Node) {
		src := htmlquery.SelectAttr(n,  "src" )
		items = append(items, src)
	})

	return items
}

func (x *dmm) ProviderUrl() string {
	return "http://dmm.co.jp"
}

func (x *dmm) ProviderName() string {
	return "DMM"
}

func (x *dmm) ProviderDisplay() string {
	return "DMM"
}

func (x *dmm) Duration() string {
	return ""
}

func (x *dmm) Media() string {
	return `not implemented`
}

func (x *dmm) Thumbnail() string {
	t := htmlquery.FindOne(x.doc, "//meta[@property='og:image']")
	return htmlquery.SelectAttr(t,  "content" )
}


func (x *dmm) Keywords() []string {
	items := []string{}
	htmlquery.FindEach(x.doc, `//a[starts-with(@href,'/digital/videoa/-/list/=/article=keyword')]/text()`, func(index int, n *html.Node) {
		items = append(items, n.Data)
	})

	return items
}
