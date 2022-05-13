# A Tour of Go
%%t220129%%
[A Tour of Go](https://go.dev/tour/list)
三小时速通完成，我感觉我行了

[go.mod file not found in current directory or any parent directory - Stack Overflow](https://stackoverflow.com/questions/66894200/go-go-mod-file-not-found-in-current-directory-or-any-parent-directory-see-go)
`go mod init helloworld`

## Package
- 为什么 import 要带引号？
- In Go, a name is exported if it begins with a capital letter.

## Variable
- The type comes *after* the variable name.
- `var a, b int`
- `var a, b = 1, "s"`
- Short variable declarations
	`a := 1`
- 
	```go
	var (
		ToBe bool = false
		MaxInt uint64 = 1<<64 - 1
		z complex128 = cmplx.Sqrt(-5 + 12i)
	)
	```
- Constant: `const {name}`
	Numeric constants are high-precision *values*.

## Types
- Basic types:
	```go
	bool
	
	string
	
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	
	byte // alias for uint8
	
	rune // alias for int32
	     // represents a Unicode code point
	 
	float32 float64
	
	complex64 complex128
	```
- Type conversion: `T(v)`
	没有隐式转换
- `v.(type)`
	```go
	func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
	```

- Pointer
	`var p *int = &a`
	- 不支持算术运算
	- `p := &Struct{}`
	- 可以直接 `p.X`（而不用 `(*p).X` 或 `p->X`）
- Struct
	```go
	type Vertex struct {
		X int
		Y int
	}
	```
	`v := Vertex{1, 2}`
	- `v := Vertex{Y:2, X: 1}`
	- 
		```go
		v := struct{
			X int
			Y int
		}{1, 2}
		```
- Array
	`var a [2]int`
	- `[2]int{1, 2}`
	- 多维数组
		```go
		board := [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}
		```
- Slice
	`var s []int = nil`
	- `s := a[0:2]`
	- `s := []int{1, 2}`
	- `s := make([]int, len, cap)`
	
	length: `len(s)`
	capacity: `cap(s)`
	- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	
	Appending: `func append(s []T, vs ...T) []T`
- Map
	`m := make(map[string]Struct)`
	```go
	m := map[string]Struct{
		"A": Struct {...},
		"B": {...}
	}
	```
	`elem, ok := m[key]`
	`delete(m, key)`

## Function
- 不带 * 皆为传值
- Multiple results
	```go
	func swap(x, y string) (string, string) {
		return y, x
	}
	```
- Named return values
	```go
	func split(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}
	```
- Function type
	```go
	func compute(fn func(float64, float64) float64) float64 {
		return fn(3, 4)
	}
	```
- Function closure
	A closure is a function value that references variables from outside its body.
	```go
	func adder() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}
	
	func main() {
		pos, neg := adder(), adder()
			for i := 0; i < 10; i++ {
			fmt.Println(
				pos(i),
				neg(-2*i),
			)
		}
	}
	```
	```go
	func fibonacci() func() int {
		a, b := 0, 1
		return func() int {
			var c int
			c, a, b = a, b, a+b
			return c
		}
	}
	
	func main() {
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}
	}
	```
	Each closure is bound to its own sum variable.

## Method
```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```
```go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{1, 2}
	v.Scale(2)
}
```
You can only declare a method with a receiver whose type is defined in the same package as the method.
As a convenience, Go interprets the statement `v.Scale(2)` as `(&v).Scale(2)` since the Scale method has a pointer receiver. The same for methods with value receivers.

## Interface
An *interface* type is defined as a set of method signatures.
```go
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	
	a = f // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser
	
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v
	
	fmt.Println(a.Abs())
}
```

The interface type that specifies zero methods is known as the *empty interface*:
`interface{}`
An empty interface may hold values of any type. (Every type implements at least zero methods.) Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

Type assertion:
`t, ok := i.(T)`

Type switch:
```go
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
```

Error:
```go
type error interface {
	Error() string
}
```
```go
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted integer:", i)
```

String:
```go
type Stringer interface {
	String() string
}
```

为什么不把 len 设计为 interface 呢？

## Flow control statements
- for
	```go
	for i := 0; i < 10; i++ {
		sum += i
	}
	```
	```go
	// for i
	// for _, v
	for i, v := range []int{1, 2, 4} {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	```
	while:
	```go
	for sum < 1000 {
		sum += sum
	}
	```
	```go
	for {
	}
	```
- if
	```go
	if v := math.Pow(x, n); v < lim {
	}
	```
- switch
	```go
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	```
	```go
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	```
	可以 fallthrough
- defer
	A defer statement defers the execution of a function until the surrounding function returns.
	```go
	func main() {
		defer fmt.Println("world")
		
		fmt.Println("hello")
	}
	```
	last-in-first-out

## Concurrency
Go provides concurrency features as part of the core language.

A *goroutine* is a lightweight thread managed by the Go runtime.
`go f(x, y, z)`

Channel:
```go
ch := make(chan T, bufLen)
ch <- v  // Send v to channel ch.
close(ch)
v, ok := <-ch  // Receive from ch, and assign value to v.
```
By default, sends to a buffered channel block only when the buffer is full and receives block when the buffer is empty. This allows goroutines to synchronize without explicit locks or condition variables.

Select:
```go
The select statement lets a goroutine wait on multiple communication operations.
tick := time.Tick(100 * time.Millisecond)
boom := time.After(500 * time.Millisecond)
for {
	select {
	case <-tick:
		fmt.Println("tick.")
	case <-boom:
		fmt.Println("BOOM!")
		return
	default:
		fmt.Println(" .")
		time.Sleep(50 * time.Millisecond)
	}
}
```
The default case in a select is run if no other case is ready.

Mutual exclusion:
`sync.Mutex`
```go
// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}
```
原来 defer 是用在这里的

## Exercises
exerise-slices.go
```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := 0; i < dy; i++{
		p[i] = make([]uint8, dx)
		for j := 0; j < dx; j++{
			p[i][j] = uint8((i+j)/2)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
```

exercise-maps.go
```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, field := range strings.Fields(s){
		m[field]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
```

exercise-stringer.go
```go
package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```