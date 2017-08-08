package extractor

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/antchfx/xquery/html"
)

func TestPorhub(t *testing.T) {
	f, _ := os.Open(`./examples/pornhub.html`)

	doc, _ := htmlquery.Parse(f)

	p :=  &pornhub{doc: doc}

	assert.NotEmpty(t, p.Title())
	assert.NotEmpty(t, p.Thumbnail())
	assert.NotEmpty(t, p.Duration())

	assert.Equal(t, "https://pornhub.com", p.ProviderUrl())
	assert.Equal(t, "pornhub", p.ProviderName())
	assert.Equal(t, "pornhub",p.ProviderDisplay())
	assert.Len(t, p.Thumbnails(), 16)
	assert.Len(t, p.Keywords(), 3)
}
