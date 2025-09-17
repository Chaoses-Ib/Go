# fmt
[fmt package - fmt - pkg.go.dev](https://pkg.go.dev/fmt)

## Scanning
For strings, byte slices and byte arrays, precision limits the length of the input to be formatted (not the size of the output), truncating if necessary. Normally it is measured in runes, but for these types when formatted with the `%x` or `%X` format it is measured in bytes.

`func Scanf(format string, a ...interface{}) (n int, err error)`
It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.