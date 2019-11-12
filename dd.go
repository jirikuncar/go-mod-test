package dd

import (
	dd1 "github.com/jirikuncar/go-mod-test/v1/dd"
	dd2 "github.com/jirikuncar/go-mod-test/v2/dd"
)

const Version string = "0.0.1"

var Apis = map[int]string{
	1: dd1.ApiVersion,
	2: dd2.ApiVersion,
}
