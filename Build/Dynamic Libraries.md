# Dynamic Libraries
Go 默认会使用部分动态库。[^moonlightwatch]

## Shared
shared [^shared][^shared-individual] 只支持 x64-linux。

```cmd
set GOOS=linux
set PATH=%PATH%;C:\L\Programs\mingw64\bin
go install -buildmode=shared -linkshared std
```
如果用 MinGW-w64 会导致：
```cmd
C:\Program Files\Go\pkg\tool\windows_amd64\link.exe: running gcc failed: exit status 1  
gcc: error: unrecognized command line option '-rdynamic'
```
看来只能在 Linux 上编译了。

shared 本来要被移除，但因为遭到反对而没有彻底移除。[^shared-remove]


[^shared]: [Shared library in Go? - Stack Overflow](https://stackoverflow.com/questions/1757090/shared-library-in-go)
[^shared-individual]: [Dynamic linking against individual golang standard libraries](https://groups.google.com/g/golang-nuts/c/-xrX-lzq-vY)
[^shared-remove]: [cmd/go: remove -buildmode=shared (not c-shared) · Issue #47788 · golang/go](https://github.com/golang/go/issues/47788)


## C-shared
c-shared 会在编译后调用 gcc 的链接器，不过似乎没办法导出 imported package 中的函数 [^cgo-export]。

[golang编译为dll与调用dll简单样例 - 大墨垂杨 - 博客园](https://www.cnblogs.com/quchunhui/p/16963248.html)

[Recommended way to distribute pre-compiled C-dependencies for modules using CGO ?](https://groups.google.com/g/golang-nuts/c/ahZXdoClBGg)

[^cgo-export]: [go - cgo doesn't export functions in imported packages - Stack Overflow](https://stackoverflow.com/questions/58433624/cgo-doesnt-export-functions-in-imported-packages)


## Plugin
plugin [^plugin] 不支持 Windows。

[^plugin]: [一文搞懂Go语言的plugin | Tony Bai](https://tonybai.com/2021/07/19/understand-go-plugin/)