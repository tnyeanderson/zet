package pkgvars

import (
	"fmt"
	"pkgvars/a"
	"pkgvars/b"
	"pkgvars/c"
)

func RunTest() {
	fmt.Println(a.Val)
	b.SetVal("b")
	fmt.Println(a.Val)
	c.SetVal("c")
	fmt.Println(a.Val)
}
