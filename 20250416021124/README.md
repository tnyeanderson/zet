# Maximum character repeat using regexp vs io.Reader in Golang

Say you have a string (or byte slice) and you want to transform it by ensuring
that a given character is never repeated more than `n` times in a row within
that string. For example, you want to ensure that there are never more than two
newlines in a row, to ensure that the paragraphs will be spaced evenly. For
instance, if the original string contained 5 `\n` characters in a row, the
returned result would only have two `\n` characters instead.

This is really easy with a regular expression function:

```golang
func removeRepeatsWithRegex(in []byte, char byte, maxRepeats int) []byte {
	repl := strings.Repeat(string(char), maxRepeats)
	re := regexp.MustCompile(repl + "+")
	return re.ReplaceAllLiteral(in, []byte(repl))
}
```

However, I was worried about performance with this. I considered using an
io.Reader which would just ignore any repeated characters that shouldn't be
there. So I ran some benchmarks:

```txt
BenchmarkRemoveRepeatsWithRegex-12                        170218        6484 ns/op
BenchmarkRemoveRepeatsWithReader-12                      3058422         397.4 ns/op
BenchmarkRemoveRepeatsWithRegex_LoremIpsum-12              19786       53568 ns/op
BenchmarkRemoveRepeatsWithReader_LoremIpsum-12           3374498         330.2 ns/op
BenchmarkRemoveRepeatsWithRegex_LoremIpsum10x-12            1044     1132094 ns/op
BenchmarkRemoveRepeatsWithReader_LoremIpsum10x-12        3042264         390.6 ns/op
BenchmarkRemoveRepeatsWithRegex_LoremIpsum100x-12            192     6709638 ns/op
BenchmarkRemoveRepeatsWithReader_LoremIpsum100x-12       2943762         413.6 ns/op
PASS
ok    maxrepeats  9.393s
```

The first two tests use a very short string as the input. However, subsequent
tests use five paragraphs of Lorem Ipsum text, then 50 paragraphs, then 500
paragraphs.

Clearly, the solution using regular expressions uses much more time when the
input is larger. In contrast, the time that the io.Reader implementation takes
doesn't really change at all as the input size increases, and is 16x faster
even on the smallest input!

The io.Reader implementation used in these benchmarks is rather
straightforward, but is certainly more complex than the regexp implementation.
If you are only ever going to have very small inputs, or otherwise don't care
about performance, the regexp implementation is probably the right answer
because it's *simpler*, so it will be easier to maintain.

If your input strings could be quite large, you will definitely want to use the
io.Reader implementation, since it is WAY more performant. For the 500
paragraph input, it was 16000x faster.

Feel free to check the [maxrepeats](./maxrepeats) directory for all the code!

    #go #performance #benchmarks
