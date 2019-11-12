# Go Package Versioning Playground

Verify naming scheme for developing versioned client APIs.

## Usage

Import root `dd` packge:

```go
import "github.com/jirikuncar/go-mod-test/dd" // initial version of root package
import "github.com/jirikuncar/go-mod-test/v1/dd" // version 1.x
```

Import specific API versions:

```go
import ddV1 "github.com/jirikuncar/go-mod-test/dd/v1/dd" // initial version of dd client v1
import ddV2 "github.com/jirikuncar/go-mod-test/dd/v2/dd" // initial version of dd client v2
import ddv1V1 "github.com/jirikuncar/go-mod-test/v1/v1/dd" // version 1.x of dd client v1
import ddv1V2 "github.com/jirikuncar/go-mod-test/v1/v2/dd" // version 1.x of dd client v2
```
