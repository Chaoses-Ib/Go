# Binary Size
[^xaionaro]
## -ldflags "-s -w"
`go build -ldflags "-s -w"`
- -s: disable symbol table
- -w: disable DWARF generation

[Guide to compress GoLang binaries](https://github.com/akshaybharambe14/golang-examples/tree/master/binaryCompression)

[Size comparison of Go executables | by Petr Jahoda | ITNEXT](https://itnext.io/size-comparison-of-go-executables-9b4ae2aaebb5)


## Strip function names
```sh
TOOL_LINK="$(readlink -f "$(go env GOROOT)"/pkg/tool/*/link)"
pushd "$(go env GOROOT)"/src/cmd/link
sed -re 's/(start := len\(ftab.P\))/\1; return int32(start)+1/' \
    -i "$(go env GOROOT)"/src/cmd/link/internal/ld/pcln.go
go build
sudo mv "$TOOL_LINK" "$TOOL_LINK".orig
sudo mv link "$TOOL_LINK"
popd
GOARCH=386 go build -a -gcflags=all="-l -B" -ldflags="-w -s"
sudo mv "$TOOL_LINK".orig "$TOOL_LINK"
stat -c %s helloworld
```


## Other flags
- Disable function inlining: `-gcflags=all=-l`
- Disable bounds checks: `-gcflags=all=-B`
- ~~Disable write-barrier: `-gcflags=all=-wb=false`~~


## [→Dead code elimination](Dead%20Code%20Elimination/README.md)

## [→Dynamic libraries](Dynamic%20Libraries.md)

## Others
- 32-bit instead of 64-bit
- Use GCCGo
- UPX
- Change version of Go



[^xaionaro]: [documentation/reduce-binary-size.md](https://github.com/xaionaro/documentation/blob/master/golang/reduce-binary-size.md)
[^moonlightwatch]: [Golang 编译优化 - 浮生偶记](https://blog.moonlightwatch.com/%E4%BB%A3%E7%A0%81%E7%AC%94%E8%AE%B0/2021/12/02/golang-%E7%BC%96%E8%AF%91%E4%BC%98%E5%8C%96.html)
