package ehtml_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/epes/ehtml"
	"golang.org/x/net/html"
)

func TestFindJ(t *testing.T) {
	tests := []struct {
		in         string
		expression string
		len        int
	}{
		{"", "", 0},
		{"", ".hello", 0},
		{"", ".hello a *", 0},

		{"<div>foobar</div>", "div", 1},
		{"<div>foobar</div>", "div *", 1},
		{"<div>foobar</div>", "div div", 0},
		{"<div>foobar</div>", "*", 1},

		{`<div class="one two"><script>alert(1)</script><script>alert(2)</script></div>`, "", 0},
		{`<div class="one two"><script>alert(1)</script><script>alert(2)</script></div>`, "*", 2},
		{`<div class="one two"><script>alert(1)</script><script>alert(2)</script></div>`, "div script", 2},
		{`<div class="one two"><script>alert(1)</script><script>alert(2)</script></div>`, "div script *", 2},
		{`<div class="one two"><script>alert(1)</script><script>alert(2)</script></div>`, ".one script *", 2},

		{`<div class="one"><div class="two three">foo<div class="four">bar</div></div></div>`, "", 0},
		{`<div class="one"><div class="two three">foo<div class="four">bar</div></div></div>`, ".one .three .four", 1},
		{`<div class="one"><div class="two three">foo<div class="four">bar</div></div></div>`, "*", 2},
		{`<div class="one"><div class="two three">foo<div class="four">bar</div></div></div>`, ".one .three .four *", 1},
		{`<div class="one"><div class="two three">foo<div class="four">bar</div></div></div>`, ".two *", 2},
		{`<div class="one"><div class="two three">foo<div class="four">bar</div></div></div>`, ".one div", 2},
	}

	for _, tt := range tests {
		doc, err := html.Parse(strings.NewReader(tt.in))
		if err != nil {
			t.Error(tt.len, tt.in, tt.expression, err)
		}

		found := ehtml.FindJ(doc, tt.expression)
		if len(found) != tt.len {
			t.Error(
				fmt.Sprintf("got:%d want:%d", len(found), tt.len),
				tt.in,
				tt.expression,
			)
		}
	}
}
