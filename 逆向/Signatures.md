# Signatures
%%t220506~07%%
## Making signatures
### Executable
[[Dead Code Elimination]]


### Dynamic libraries
shared 只支持 x64-linux。

c-shared 需要 `//export`，而且没办法导出 imported package 中的函数。


### Static libraries
archive 是专有格式，IDA 识别不出函数。

c-archive 与 c-shared 有同样的问题。