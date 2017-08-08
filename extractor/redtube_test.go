package extractor

import (
	"testing"
	"os"
	"github.com/antchfx/xquery/html"
	"github.com/stretchr/testify/assert"
)

func TestRedtube(t *testing.T) {
	f, _ := os.Open(`./examples/redtube.html`)

	doc, _ := htmlquery.Parse(f)

	p :=  &redtube{doc: doc}

	assert.NotEmpty(t, p.Title())
	assert.NotEmpty(t, p.Thumbnail())
	//assert.NotEmpty(t, p.Duration())

	assert.Equal(t, "https://redtube.com", p.ProviderUrl())
	assert.Equal(t, "redtube", p.ProviderName())
	assert.Equal(t, "redtube",p.ProviderDisplay())
	assert.Len(t, p.Thumbnails(), 16)
	assert.Len(t, p.Keywords(), 5)
}
