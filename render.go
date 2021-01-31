package ehtml

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

// RenderNodeAsHTML takes an html Node and returns a string
// representation of the rendered html tree.
func RenderNodeAsHTML(node *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, node)
	return buf.String()
}
