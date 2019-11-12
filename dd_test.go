package dd_test

import (
	"testing"

	dd "github.com/jirikuncar/go-mod-test/v1"
	dd1 "github.com/jirikuncar/go-mod-test/v1/v1/dd"
	dd2 "github.com/jirikuncar/go-mod-test/v1/v2/dd"
)

func TestVersion(t *testing.T) {
	expected := "1.0.0"
	if dd.Version != expected {
		t.Errorf("got %s; expected %s", dd.Version, expected)
	}
}

func TestApis(t *testing.T) {
	expected := map[string]string{
		"v1": dd1.ApiVersion,
		"v2": dd2.ApiVersion,
	}
	for v, av := range dd.Apis {
		e := expected[v]
		if av != e {
			t.Errorf("checking %v; got %s; expected %s", v, av, e)
		}
	}
}
