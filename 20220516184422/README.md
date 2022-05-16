# Named return values in go

Avoid needlessly instantiating return variables!

Instead of:
```go
// Sum each integer in a string of only integers
func countLetterA(input string) int {
  v := 0
  for _, s := range input {
		if s == 'a' {
			v++
		}
  }
  return v
```

Do:
```go
func countLetterA(input string) (v int) {
  for _, s := range input {
		if s == 'a' {
			v++
		}
  }
  return
```

    #golang #tips

