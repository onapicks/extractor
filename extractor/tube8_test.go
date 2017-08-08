package extractor

import (
	"testing"
	"os"
	"github.com/antchfx/xquery/html"
	"github.com/stretchr/testify/assert"
)

func TestTube8(t *testing.T) {
	f, _ := os.Open(`./examples/tube8.html`)

	doc, _ := htmlquery.Parse(f)

	p :=  &tube8{doc: doc}

	assert.NotEmpty(t, p.Title())
	assert.NotEmpty(t, p.Thumbnail())
	//assert.NotEmpty(t, p.Duration())

	assert.Equal(t, "https://tube8.com", p.ProviderUrl())
	assert.Equal(t, "tube8", p.ProviderName())
	assert.Equal(t, "tube8",p.ProviderDisplay())

	assert.Len(t, p.Thumbnails(), 16)
	assert.Len(t, p.Keywords(), 1)
}
