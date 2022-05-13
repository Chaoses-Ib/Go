package main

import (
	"fmt"
	"go/importer"
	"go/types"
	"reflect"

	test "./test"
)

func main() {
	test.A()

	v := test.Vertex{X: 1, Y: 2}
	v.C()

	pkg, err := importer.Default().Import("fmt")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		fmt.Println(name)

		obj := scope.Lookup(name)
		if tn, ok := obj.Type().(*types.Named); ok {
			fmt.Printf("%#v\n", tn.NumMethods())
		}
	}

	//fmt.Println(reflect.ValueOf(pkg).Method(0).Call([]reflect.Value{}))
	//fmt.Println(reflect.ValueOf(pkg).MethodByName("Close").Call([]reflect.Value{}))

	fmt.Println(reflect.TypeOf(pkg).Method(0))
	fmt.Println(reflect.TypeOf(pkg).MethodByName("Close"))
}
