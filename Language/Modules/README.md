# Modules
> Go Modules (previously called vgo -- versioned Go) were introduced with Go 1.11. Instead of using the GOPATH for storing a single git checkout of every package, Go Modules stores tagged versions with `go.mod` keeping track of each package's version.

## `GO111MODULE`
[Why is `GO111MODULE` everywhere, and everything about Go Modules (updated with Go 1.20) | maelvls dev blog](https://maelvls.dev/go111module-everywhere/)

Symptoms:
- modules disabled by `GO111MODULE=off`

  `go help modules` doesn't answer why it's disabled by default and how to enable it.

- cannot find package "github.com/example"

  [go版本到1.22后，编译碰到的问题和解决经过，希望能给和我一样的小白一点参考。 - Issue #148 - kenzok8/small](https://github.com/kenzok8/small/issues/148)

  [编译官方固件失败，请大神帮看看-OPENWRT专版-恩山无线论坛](https://www.right.com.cn/forum/thread-8397909-1-1.html)
- VS Code failed to install `gopls`, `staticcheck`, `dlv`

Solutions:
- `go env -w GO111MODULE=on`
- Copy the modules to `go/src`

[How to fix "go get: warning: modules disabled by GO111MODULE=auto in GOPATH/src" - Stack Overflow](https://stackoverflow.com/questions/56475313/how-to-fix-go-get-warning-modules-disabled-by-go111module-auto-in-gopath-src)
