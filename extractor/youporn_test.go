package extractor

import (
	"testing"
	"os"
	"github.com/antchfx/xquery/html"
	"github.com/stretchr/testify/assert"
)

func TestYouporn(t *testing.T) {
	f, _ := os.Open(`./examples/youporn.html`)

	doc, _ := htmlquery.Parse(f)

	p :=  &youporn{doc: doc}

	assert.NotEmpty(t, p.Title())
	assert.NotEmpty(t, p.Thumbnail())

	//assert.NotEmpty(t, p.Duration())

	assert.Equal(t, "https://youporn.com", p.ProviderUrl())
	assert.Equal(t, "youporn", p.ProviderName())
	assert.Equal(t, "youporn",p.ProviderDisplay())
	assert.Len(t, p.Thumbnails(), 16)
	assert.Len(t, p.Keywords(), 4)
}
