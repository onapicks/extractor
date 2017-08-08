package extractor

import (
	"regexp"
	"golang.org/x/net/html"
)

type Extractor interface {
	Title() string

	Thumbnails() []string

	ProviderUrl() string

	ProviderName() string

	ProviderDisplay() string

	Duration() string
	//Media() string
	Thumbnail() string

	Keywords() []string
}

func Extract(url string) func(*html.Node) Extractor {
	if regexp.MustCompile(`xvideos.com`).MatchString(url) {
		return func(doc *html.Node) Extractor {
			return &xvideo{doc}
		}
	} else if regexp.MustCompile(`pornhub.com`).MatchString(url) {
		return func(doc *html.Node) Extractor {
			return &pornhub{doc}
		}
	} else if regexp.MustCompile(`dmm.co.jp`).MatchString(url) {
		return func(doc *html.Node) Extractor {
			return &dmm{doc}
		}
	} else if regexp.MustCompile(`xhamster.com`).MatchString(url) {
		return func(doc *html.Node) Extractor {
			return &xhamster{doc}
		}
	} else if regexp.MustCompile(`redtube.com`).MatchString(url) {
		return func(doc *html.Node) Extractor {
			return &redtube{doc}
		}
	} else if regexp.MustCompile(`tube8.com`).MatchString(url) {
		return func(doc *html.Node) Extractor {
			return &tube8{doc}
		}
	} else if regexp.MustCompile(`youporn.com`).MatchString(url) {
		return func(doc *html.Node) Extractor {
			return &youporn{doc}
		}
	}

	return nil
}
