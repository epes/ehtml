package ehtml_test

import (
	"testing"

	"github.com/epes/ehtml"
)

func TestFindJ(t *testing.T) {
	ehtml.FindJ(nil, ".hello a .world")
}
