# Anonymous struct in go

To create an anonymous struct:
```go
myStruct := struct {
  prop1 string
  prop2 string
}{
  prop1: "val1",
  prop2: "val2",
}
```

The first set of curly braces defines the name and type of the structs properties.
It has no colons or commas.

The second set of curly braces assigns values to the properties. Colons and commas
are required.

    #go #struct #oo #tips
