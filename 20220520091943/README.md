# Spread operator in go (pack and unpack)

In go, there is something similar to javascript's spread operator that can be used for variadic functions:
```go
func VariadicFunc(args ...int) {
  for _, arg := range args {
    doStuff(arg)
  }
}

func main() {
	arr := []int{4,3,2,1}
	VariadicFunc(arr...)
}
```

The `pack` operator is used in the argument definition (`func VariadicFunc(args ...int)`). Note the ellipse *preceeds*
the type. The `unpack` operator is used when calling the variadic function (`VariadicFunc(arr...)`). Note the ellipse
*follows* the type.

Be aware of type mismatches:
```go
// ERROR: Type mismatch
func VariadicFunc(args ...any) {
  for _, arg := range args {
    doStuff(arg)
  }
}

func main() {
	arr := []int{4,3,2,1}
	VariadicFunc(arr...)
}
```

This won't work, since `arr` is an array of primitive integers, which do not implement `any`.

>Remember, `any` is an alias for `interface{}`

First, convert the primitives array to an array of empty interfaces:
```go
func VariadicFunc(args ...any) {
  for _, arg := range args {
    doStuff(arg)
  }
}

func main() {
  var args []interface{}
  arr := []int{4,3,2,1}
  for _, v := range arr {
    args = append(args, v)
  }
  VariadicFunc(args...)
}
```

    #go #tips #vocab
