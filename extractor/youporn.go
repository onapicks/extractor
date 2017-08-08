package extractor

import (
	"regexp"
	"fmt"
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

type youporn struct {
	doc *html.Node
}

func (x *youporn) Title() string {
	t := htmlquery.FindOne(x.doc, "//title/text()")
	if t != nil {
		return t.Data
	}
	return ""
}

func (x *youporn) Thumbnails() []string {
	t := x.Thumbnail()

	if t == "" {
		return []string{}
	}

	items := make([]map[int]int, 16)

	ls := []string{}
	re := regexp.MustCompile(`(.+original\/)\d+(\/.+\.jpg)$`)

	for i := range items {
		r := fmt.Sprintf("${1}%d${2}", i+1)
		thumb := re.ReplaceAllString(t, r)
		ls = append(ls, thumb)
	}

	return ls
}

func (x *youporn) Thumbnail() string {
	t := htmlquery.FindOne(x.doc, "//meta[@property='og:image']")
	if t == nil {
		return ""
	}
	return htmlquery.SelectAttr(t, "content")
}

func (x *youporn) ProviderUrl() string {
	return "https://youporn.com"
}

func (x *youporn) ProviderName() string {
	return "youporn"
}

func (x *youporn) ProviderDisplay() string {
	return "youporn"
}

func (x *youporn) Duration() string {
	t := htmlquery.FindOne(x.doc, "//span[@class='video-box-duration']")

	if t == nil {
		return ""
	}

	return t.Data
}

func (x *youporn) Keywords() []string {
	items := []string{}

	expr := `//div[@class='info-column tagsSuggest']//a/text()`
	htmlquery.FindEach(x.doc, expr, func(index int, n *html.Node) {
		items = append(items, n.Data)
	})
	return items
}
