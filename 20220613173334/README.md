# Concatenate multiple values with jq

Given the following JSON object:
```json
{
  "abc": "xyz",
  "foo": "bar",
  "testarr": [
    {"name": "item1", "key": "key1", "value": "value1"},
    {"name": "item2", "key": "key2", "value": "value2"}
  ]
}
```

To interpolate values into a string, use the `\(foo)` syntax (similar to
`$(foo)` in bash):
```bash
cat test.json | jq -r '"abc is \(.abc) foo is \(.foo)"'

# Or:
cat test.json | jq -r '[.abc,.foo] | "abc is \(.[0]) foo is \(.[1])"'
```

> NOTE: You need double quotes to actually be sent to `jq` so the result is
> parsed as a string, so put them *inside* the single quotes

Output:
```
abc is xyz foo is bar
```

Show the name, key, and value of each item in `testarr` as a tab-separated
list, use `@tsv` (or `@csv` for a comma separated list):
```bash
cat test.json | jq -r '.testarr[] | [.name,.key,.item] | @tsv'
```

Output:
```
item1   key1
item2   key2
```

Remember to use `-r` to get the raw output or it will be a single string *containing* a tab separated string:
```
"item1\tkey1\t"
"item2\tkey2\t"
```

Show a string and a JSON value (use `tostring`):
```bash
cat test.json | jq -r '.testarr[] | [.abc,.testarr[0]|tostring] | @tsv'
```

Output:
```
xyz     {"name":"item1","key":"key1","value":"value1"}
```

  #json #bash #tips
