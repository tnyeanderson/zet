package main

import (
	"fmt"
	"pkgvars"
	"pkgvars/a"
)

func ExamplePackageVarsScope() {
	fmt.Println(a.Val)
	pkgvars.RunTest()
	fmt.Println(a.Val)
	a.Val = "outside"
	fmt.Println(a.Val)
	pkgvars.RunTest()
	fmt.Println(a.Val)

	// Output:
	// a
	// a
	// b
	// c
	// c
	// outside
	// outside
	// b
	// c
	// c
}
