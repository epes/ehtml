package ehtml_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/epes/ehtml"
	"golang.org/x/net/html"
)

func TestRenderNodeAsHTML(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"<div></div>", "<html><head></head><body><div></div></body></html>"},
		{"hi", "<html><head></head><body>hi</body></html>"},
	}

	for _, tt := range tests {
		doc, err := html.Parse(strings.NewReader(tt.in))
		if err != nil {
			t.Error(tt.in, tt.out, err)
		}

		rendered := ehtml.RenderNodeAsHTML(doc)

		if rendered != tt.out {
			t.Error(
				fmt.Sprintf("got:%s want:%s", rendered, tt.out),
				tt.in,
				tt.out,
			)
		}
	}
}
