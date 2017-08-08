package extractor

import (
	"testing"
	"os"
	"github.com/antchfx/xquery/html"
	assert "github.com/stretchr/testify/assert"
)

func TestDmm(t *testing.T) {
	f, _ := os.Open(`./examples/dmm.html`)

	doc, _ := htmlquery.Parse(f)

	d :=  &dmm{doc: doc}
	assert.NotEmpty(t, d.Title())
	assert.NotEmpty(t, d.Thumbnail())

	assert.Equal(t, "http://dmm.co.jp", d.ProviderUrl())
	assert.Equal(t, "DMM", d.ProviderName())
	assert.Equal(t, "DMM",d.ProviderDisplay())
	assert.Len(t, d.Thumbnails(), 11)
	assert.Len(t, d.Keywords(), 19)
}
