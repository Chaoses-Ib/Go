# Dead Code Elimination
Elimination of unused public methods at link time will be inhibited if reflection function calling are used anywhere. [^u-root]

[ssa/deadcode.go](https://github.com/golang/go/blob/master/src/cmd/compile/internal/ssa/deadcode.go)
- early deadcode
- pre-opt deadcode
- opt deadcode
- gcse deadcode
- generic deadcode
- lowered deadcode for cse
- lowered deadcode
- late deadcode

[ld/deadcode.go](https://github.com/golang/go/blob/master/src/cmd/link/internal/ld/deadcode.go)
- `reflect.Value.Method`
- `reflect.Value.MethodByName`
- `reflect.Type.Method`
- `reflect.Type.MethodByName`

只会影响到 ValueOf 和 TypeOf 涉及到的类型。


## 禁用 DCE
没办法直接禁用 DCE，用反射也只能对 ValueOf 涉及到的类型生效。


修改 [ld/deadcode.go](https://go.dev/src/cmd/link/internal/ld/deadcode.go)？
直接修改 `d.ctxt.BuildMode == BuildModeShared` 会导致：
```
_rt0_amd64_lib: relocation target runtime.libpreinit not defined
```
修改其它地方没报错，但也没生效，可能是在 ssa 阶段就被 dce 了。


修改 ssa？
移除所有 deadcode 的 pass 会导致：
```
C:\Temp\go\src\internal\cpu\cpu_x86.go:62:26: internal compiler error: 'doinit': two final stores - simultaneous live st  
ores v127 = Copy <mem> v366 v136 = SelectN <mem> [0] v135
```


[^u-root]: [Enable dead code elimination to reduce binary size · Issue #1477 · u-root/u-root](https://github.com/u-root/u-root/issues/1477)