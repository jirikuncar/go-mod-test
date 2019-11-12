package dd_test

import (
	"testing"

	dd0 "github.com/jirikuncar/go-mod-test"
	dd "github.com/jirikuncar/go-mod-test/v1"
	dd1 "github.com/jirikuncar/go-mod-test/v1/v1/dd"
	dd2 "github.com/jirikuncar/go-mod-test/v1/v2/dd"
)

func TestInitialVersion(t *testing.T) {
	expected := "0.0.1"
	if dd0.Version != expected {
		t.Errorf("got %s; expected %s", dd0.Version, expected)
	}
}


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
