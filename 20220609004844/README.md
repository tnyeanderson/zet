# Convert struct to JSON and vice versa

To convert a `struct` to JSON data, the structs properties **must be exported (public)!**
In the following example, `prop3` can never be marshaled (converted from JSON) or
unmarshaled (converted to JSON) since it is not an exported (capitalized) property:
```go
type MyStruct struct {
  Prop1 string
  Prop2 string
  prop3 string
}
```

But does that mean that the JSON key also have to be capitalized? By default,
yes. but the JSON key that is linked to the property can be overwritten:
```go
type MyStruct struct {
  Prop1 string `json:"prop1"`
  Prop2 string `json:"prop_2"`
  prop3 string `json:"Property3"`
}
```

> NOTE: Even though `prop3` has a capitalized JSON key, it still can never be
> marshaled or unmarshaled!

    #go #struct #json
