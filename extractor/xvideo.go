package extractor


import (
	"regexp"
	"fmt"
	"strings"
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

type xvideo struct {
	doc *html.Node
}

func (x *xvideo) Title() string {
	t := htmlquery.FindOne(x.doc, "//title/text()")
	if t != nil {
		return t.Data
	}
	return ""
}

func (x *xvideo) Duration() string {
	duration := htmlquery.FindOne(x.doc, "//meta[@property='og:duration']")
	n := htmlquery.SelectAttr(duration,  "content" )
	return n
}

func (x *xvideo) Thumbnail() string {
	t := htmlquery.FindOne(x.doc, "//meta[@property='og:image']")
	return htmlquery.SelectAttr(t,  "content" )
}

func (x *xvideo) Thumbnails() []string {
	items := make([]map[int]int, 30)

	ls := []string{}

	re := regexp.MustCompile(`(.+\.)\d+(.jpg)$`)

	for i := range items {
		r := fmt.Sprintf("${1}%d${2}", i+1)
		thumb := re.ReplaceAllString(x.Thumbnail(), r)
		ls = append(ls, thumb)
	}
	return ls
}
func (x *xvideo) Keywords() []string {
	duration := htmlquery.FindOne(x.doc, "//meta[@name='keywords']")
	n := htmlquery.SelectAttr(duration,  "content" )
	return strings.Split(n, ",")
}

func (x *xvideo) Description() string {
	duration := htmlquery.FindOne(x.doc, "//meta[@name='description']")
	return htmlquery.SelectAttr(duration,  "content" )
}

func (x *xvideo) ProviderUrl() string {
	return "https://xvideo.com"
}

func (x *xvideo) ProviderName() string {
	return "xvideo"
}

func (x *xvideo) ProviderDisplay() string {
	return "xvideo"
}

func (x *xvideo) Media() string {
	return `<iframe class="embedly-embed" src="//cdn.embedly.com/widgets/media.html?src=https%3A%2F%2Fwww.youtube.com%2Fembed%2FtnBQmEqBCY0%3Ffeature%3Doembed&url=http%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DtnBQmEqBCY0&image=https%3A%2F%2Fi.ytimg.com%2Fvi%2FtnBQmEqBCY0%2Fhqdefault.jpg&key=internal&type=text%2Fhtml&schema=youtube" width="500" height="281" scrolling="no" frameborder="0" allowfullscreen></iframe>`
}


