package extractor

import (
	"testing"
	"os"
	"github.com/antchfx/xquery/html"
	"github.com/stretchr/testify/assert"
)

func TestXvideo(t *testing.T) {
	f, _ := os.Open(`./examples/xvideo.html`)

	doc, _ := htmlquery.Parse(f)

	x:= &xvideo{doc: doc}

	assert.NotEmpty(t, x.Title())
	assert.NotEmpty(t, x.Thumbnail())
	assert.NotEmpty(t, x.Duration())
	assert.NotEmpty(t, x.Keywords())
	assert.Equal(t, "https://xvideo.com", x.ProviderUrl())
	assert.Equal(t, "xvideo", x.ProviderName())
	assert.Equal(t, "xvideo",x.ProviderDisplay())
	assert.Len(t, x.Thumbnails(), 30)
	assert.Len(t, x.Keywords(), 12)
}

