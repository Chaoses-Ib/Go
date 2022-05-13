# net
%%t220415~16%%
[net package](https://pkg.go.dev/net)

## Example
Client:
```go
import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:3333")
	if err != nil {
		panic(err)
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	
	for {
		conn.Write([]byte("hello world"))
		
		reply := make([]byte, 1024)
		n, err = conn.Read(reply)
		if err != nil {
			panic(err)
		}
		fmt.Println(reply[:n])
	}
}
```
Server:
```go
import (
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:3333")
	defer ln.Close()
	if err != nil {
		panic(err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		
		go func(conn net.Conn) {
			_, err := io.Copy(conn, conn)
			defer conn.Close()
			if err != nil {
				panic(err)
			}
		}(conn)
	}
}
```
[^felipejfc]

[^felipejfc]: [felipejfc/go-tcp-example: tcp client and server written in golang](https://github.com/felipejfc/go-tcp-example)

## Read
```go
buf := []byte{}
tmp := make([]byte, 256)
for {
    n, err := conn.Read(tmp)
    if err != nil {
        if err != io.EOF {
	        panic(err)
        }
        break
    }
    buf = append(buf, tmp[:n]...)
}
```
[^read-so]

[^read-so]: [go - Read whole data with Golang net.Conn.Read - Stack Overflow](https://stackoverflow.com/questions/24339660/read-whole-data-with-golang-net-conn-read)