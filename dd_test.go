package dd_test

import (
	"testing"

	dd "github.com/jirikuncar/go-mod-test"
	dd1 "github.com/jirikuncar/go-mod-test/v1/dd"
	dd2 "github.com/jirikuncar/go-mod-test/v2/dd"
)

func TestVersion(t *testing.T) {
	expected := "0.0.1"
	if dd.Version != expected {
		t.Errorf("got %s; expected %s", dd.Version, expected)
	}
}

func TestApis(t *testing.T) {
	expected := map[int]string{
		1: dd1.ApiVersion,
		2: dd2.ApiVersion,
	}
	for v, av := range dd.Apis {
		e := expected[v]
		if av != e {
			t.Errorf("checking %v; got %s; expected %s", v, av, e)
		}
	}
}
