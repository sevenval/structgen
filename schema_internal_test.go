package structgen

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	schema, err := NewSchema([]byte("{}"))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		Have, Want string
	}{
		{"prop", "Prop"},
		{"propAbc", "PropAbc"},
		{"prop_list", "PropList"},
		{"Prop_map", "PropMap"},
		{"Prop-whatever", "PropWhatever"},
		{"PropMop", "PropMop"},
		{"", ""},
	}

	for _, c := range cases {
		if got := schema.toCamelCase(c.Have); got != c.Want {
			t.Errorf("Expected %q, got: %q", c.Want, got)
		}
	}
}
