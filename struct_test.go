package structgen_test

import (
	"testing"

	"github.com/sevenval/structgen"
)

func TestStructString(t *testing.T) {
	cases := []struct {
		Have *structgen.Struct
		Want string
	}{
		{&structgen.Struct{Name: "foo", JSONName: "foo_bar", Type: "string"}, "foo\tstring\t`json:\"foo_bar,omitempty\"`"},
		{&structgen.Struct{Name: "bar", JSONName: "bar_foo", Type: "int", Required: true}, "bar\tint\t`json:\"bar_foo\"`"},
		{&structgen.Struct{Name: "A", JSONName: "a", Type: "interface{}", Comment: "string or list"}, "A\tinterface{}\t`json:\"a,omitempty\"` // string or list"},
	}
	for _, c := range cases {
		if s := c.Have.String(); s != c.Want {
			t.Errorf("Expected: %s; got: %s", c.Want, s)
		}
	}
}
