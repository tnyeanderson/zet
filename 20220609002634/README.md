# Read piped input in go

Reading input from stdin that has been piped is kind of annoying in go:
```go
func readFromStdin() (text string) {
	scanner := bufio.NewScanner(os.Stdin)
	// Scan each line
	for scanner.Scan() {
		// Add back the line break
		text += scanner.Text() + "\n"
	}
	// Remove the last extra line break
	text = strings.TrimSuffix(text, "\n")
	return
}
```

> NOTE: This will hang if no input was piped in (it will wait for user input)

    #go #stdin #scripts
