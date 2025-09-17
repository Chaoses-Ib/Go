# FFI
## cgo
[cgo command - cmd/cgo - Go Packages](https://pkg.go.dev/cmd/cgo)

[Cgo: When and (Usually) When Not to Use it - Repeatable Systems](https://relistan.com/cgo-when-and-when-not-to-use-it)

```go
/*
#include <stdbool.h>
#include <stdlib.h>
#include "HAP_SDK.h"
#cgo amd64 LDFLAGS: -L./lib -lHAP_SDK64
#cgo 386 LDFLAGS: -L./lib -lHAP_SDK32
*/
import "C"
```

## [→C-shared](../Build/Dynamic%20Libraries.md#c-shared)
