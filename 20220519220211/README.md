# Variadic functions

A variadic function takes an aribtrary (or infinite) amount of arguments:
```
func VariadicFunc(args ...int) {
	for _, arg := range args {
		doStuff(arg)
	}
}

VariadicFunc(1, 2)
VariadicFunc(1, 2, 3, 4, ...)
```

    #vocab #args
