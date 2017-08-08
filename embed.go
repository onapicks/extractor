package xtract

import (
	"fmt"
	"regexp"
)

type Oembed struct {
	ProviderId string `json:"provider_id"`
	ProviderName string `json:"provider_name"`
	ProviderUrl string `json:"provider_url"`
	AuthorUrl string `json:"author_url,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Version string `json:"version"`
	Html string `json:"Html"`
	Height int `json:"Height"`
	Width        int `json:"Width"`
	ThumbnailUrl string `json:"thumbnail_url"`
	ThumbnailUrls []string `json:"thumbnail_urls"`
	ThumbnailHeight int `json:"thumbnail_height"`
	ThumbnailWidth int `json:"thumbnail_width"`
	Title string `json:"title"`
	Type string `json:"type"`
}

type provider interface {
	Id() string
	Embed(s *Size) *embed
}

func Provider(url string) provider {
	if regexp.MustCompile(`xvideos.com`).MatchString(url) {
		return &xvideo{url: url}
	} else if regexp.MustCompile(`pornhub.com`).MatchString(url) {
		return &pornhub{url: url}
	} else if regexp.MustCompile(`dmm.co.jp`).MatchString(url) {
		return &dmm{url: url}
	} else if regexp.MustCompile(`xhamster.com`).MatchString(url) {
		return &xhamster{url: url}
	} else if regexp.MustCompile(`redtube.com`).MatchString(url) {
		return &redtube{url: url}
	} else if regexp.MustCompile(`tube8.com`).MatchString(url) {
		return &tube8{url: url}
	} else if regexp.MustCompile(`youporn.com`).MatchString(url) {
		return &youporn{url: url}
	} else {
		return nil
	}
}

func firstMatch(url, expr string) string {
	r, _ := regexp.Compile(expr)
	m := r.FindStringSubmatch(url)

	if len(m) <= 1 {
		return ""
	}

	return m[1]
}

type embed struct {
	Html   string
	Width  int
	Height int
}

type Size struct {
	width int
	height int
}

func (s *Size) Guard(w int, h int) {
	if s.width == 0 {
		s.width = w
	}

	if s.height == 0 {
		s.height = h
	}
}


type pornhub struct {
	url string
}

func (o *pornhub) Id() string {
	return firstMatch(o.url,`viewkey=(.+$)`)
}

func (o *pornhub) Embed(s *Size) *embed {
	s.Guard(560, 315)

	f := `<iframe src="https://jp.pornhub.com/embed/%s" frameborder="0" Width="%d" Height="%d" scrolling="no" allowfullscreen></iframe>`

	return &embed{
		fmt.Sprintf(f, o.Id(), s.width, s.height),
		s.width,
		s.height,
	}
}

type xvideo struct {
	url string
}

func (o *xvideo) Id() string {
	return firstMatch(o.url, `\/video(\d+)\/`)
}

func (o *xvideo) Embed(s *Size) *embed {
	s.Guard(510,400)
	f := `<iframe src="https://flashservice.xvideos.com/embedframe/%s" frameborder=0 Width=%d Height=%d scrolling=no allowfullscreen=allowfullscreen></iframe>`

	return &embed{
		fmt.Sprintf(f, o.Id(), s.width, s.height),
		s.width,
		s.height,
	}
}

type dmm struct {
	url string
}

func (o *dmm) Id() string {
	return firstMatch(o.url, `\/cid\=(.+)\/`)
}

func (o *dmm) Embed(s *Size) *embed {
	s.Guard(476,306)
	f := `<iframe src="http://www.dmm.co.jp/litevideo/-/part/=/cid=%s/size=%d_%d/" Width="%d" Height="%d" scrolling="no" frameborder="0" allowfullscreen></iframe>`
	return &embed{
		fmt.Sprintf(f, o.Id(), s.width, s.height, s.width, s.height),
		s.width,
		s.height,
	}
}

type xhamster struct {
	url string
}

func (o *xhamster) Id() string {
	s := firstMatch(o.url,`\-(\d+)$`)

	if s == "" {
		s = firstMatch(o.url, `\/movies\/(\d+)\/`)
	}

	return s
}

func (o *xhamster) Embed(s *Size) *embed {
	s.Guard(510,400)
	f := `<iframe src="https://xhamster.com/xembed.php?video=%s" Width="%d" Height="%d" frameborder="0" scrolling="no" allowfullscreen></iframe>`

	return &embed{
		fmt.Sprintf(f, o.Id(), s.width, s.height),
		s.width,
		s.height,
	}
}

type redtube struct {
	url string
}

func (o *redtube) Id() string {
	return firstMatch(o.url,`\/(\d+)$`)
}

func (o *redtube) Embed(s *Size) *embed {
	s.Guard(560,315)
	f := `<iframe src="https://embed.redtube.com/?id=%s&bgcolor=000000" frameborder="0" Width="%d" Height="%d" scrolling="no" allowfullscreen></iframe>`

	return &embed{
		fmt.Sprintf(f, o.Id(), s.width, s.height),
		s.width,
		s.height,
	}
}

type tube8 struct {
	url string
}

func (o *tube8) Id() string {
	return firstMatch(o.url,`tube8.com\/(.+\/.+\/\d+)`)
}

func (o *tube8) Embed(s *Size) *embed {
	s.Guard(608,342)
	f := `<iframe src="https://www.tube8.com/embed/%s" frameborder="0" Width="%d" Height="%d" scrolling="no" allowfullscreen="true" webkitallowfullscreen="true" mozallowfullscreen="true" name="t8_embed_video"></iframe>`
	return &embed{
		fmt.Sprintf(f, o.Id(), s.width, s.height),
		s.width,
		s.height,
	}
}

type youporn struct {
	url string
}

func (o *youporn) Id() string {
	return firstMatch(o.url, `watch\/(\d+)\/`)
}

func (o *youporn) Embed(s *Size) *embed {
	s.Guard(560,315)

	f := `<iframe src='https://www.youporn.com/embed/%s' frameborder=0 Width='%d' Height='%d' scrolling=no name='yp_embed_video'></iframe>`

	return &embed{
		fmt.Sprintf(f, o.Id(), s.width, s.height),
		s.width,
		s.height,
	}
}

