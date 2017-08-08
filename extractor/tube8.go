package extractor

import (
	"regexp"
	"fmt"
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

type tube8 struct {
	doc *html.Node
}

func (x *tube8) Title() string {
	t := htmlquery.FindOne(x.doc, "//title/text()")
	if t != nil {
		return t.Data
	}
	return ""
}

func (x *tube8) Thumbnails() []string {
	items := make([]map[int]int, 16)

	ls := []string{}

	re := regexp.MustCompile(`(.+\/originals\/)\d+`)

	for i := range items {
		r := fmt.Sprintf("${1}%d${2}", i+1)
		thumb := re.ReplaceAllString(x.Thumbnail(), r)
		ls = append(ls, thumb)
	}

	return ls
}

func (x *tube8) Thumbnail() string {
	t := htmlquery.FindOne(x.doc, "//meta[@property='og:image']")
	return htmlquery.SelectAttr(t,  "content" )
}


func (x *tube8) ProviderUrl() string {
	return "https://tube8.com"
}

func (x *tube8) ProviderName() string {
	return "tube8"
}

func (x *tube8) ProviderDisplay() string {
	return "tube8"
}

func (x *tube8) Duration() string {
	return ""
}

func (x *tube8) Keywords() []string {
	items := []string{}

	expr := `//li[@class='video-category']/a/text()`
	htmlquery.FindEach(x.doc, expr, func(index int, n *html.Node) {
		items = append(items, n.Data)
	})
	return items
}
