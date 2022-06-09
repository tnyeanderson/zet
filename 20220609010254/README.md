# Adding query params to a go request

Go has `Set` and `Add` methods on `URL` objects... but those don't actually add or set
the values on their own (they only parse `RawQuery` and return that with your
changes). The `RawQuery` has to be explicitly set:
```go
// Get the current query object and add/set the parameter key/value
query := req.URL.Query()
query.Set("access_token", token)
// This is what actually sets it
req.URL.RawQuery = query.Encode()
```

    #go #http #url
