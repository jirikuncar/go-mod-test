package dd

import (
	dd1 "github.com/jirikuncar/go-mod-test/v1/v1/dd"
	dd2 "github.com/jirikuncar/go-mod-test/v1/v2/dd"
)

const Version string = "1.0.0"

var Apis = map[string]string{
	"v1": dd1.ApiVersion,
	"v2": dd2.ApiVersion,
}
