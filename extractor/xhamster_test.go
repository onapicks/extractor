package extractor

import (
	"testing"
	"os"
	"github.com/antchfx/xquery/html"
	"github.com/stretchr/testify/assert"
)

func TestXhamster(t *testing.T) {
	f, _ := os.Open(`./examples/xhamster.html`)

	doc, _ := htmlquery.Parse(f)

	p :=  &xhamster{doc: doc}

	assert.NotEmpty(t, p.Title())
	assert.NotEmpty(t, p.Thumbnail())

	//assert.NotEmpty(t, p.Duration())

	assert.Equal(t, "https://xhamster.com", p.ProviderUrl())
	assert.Equal(t, "xhamster", p.ProviderName())
	assert.Equal(t, "xhamster",p.ProviderDisplay())
	assert.Len(t, p.Thumbnails(), 10)
	assert.Len(t, p.Keywords(), 1)
}
