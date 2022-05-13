# Go
> Go is a tool for managing Go source code.

- 2022-01-29
- `scoop install go`
  - [Data](#data) is not persisted.

%%t220507%%
[Installing Go from source](https://go.dev/doc/install/source)
```cmd
set PATH=%PATH%;C:\L\Programs\mingw64\bin

set GOROOT_BOOTSTRAP=C:\Program Files\Go
.\all.bat
```
- ASCII 路径
- 需要 gcc

## Data
Windows:
- `GOPATH`: `%USERPROFILE%\go`
- `%LOCALAPPDATA%\go-build`

The uninstaller *will not* delete these garbage.

[Avoid having a ~/go directory : r/golang](https://www.reddit.com/r/golang/comments/10psufn/avoid_having_a_go_directory/)
