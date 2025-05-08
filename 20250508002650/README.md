# Dice rolls, probability histograms, and recursion

If I have a number of dice, each with some number of sides, and I roll them all
at once and sum the values, how likely am I to get each potential sum?

This data can be represented as a histogram, which correlates a potential sum
to the number of ways that sum can be reached using the provided dice.

We'll build a CLI where each argument represents a die, and the value of the
argument is the number of sides that die has. The program will calculate and
print the histogram. We'll also include some benchmarks for the different
implementations to see which is faster!

Feel free to check out the [complete source code](./dice)!

If we knew exactly how many dice to support (for example, only ever supporting
two dice), we could just nest a couple loops. However, since the number of dice
is arbitrary, we need a way to loop an unknown number of times, with each
possible combination of dice values. To accomplish this, one might immediately
jump to recursion: take an array of dice, and for each side of the first die,
call the function with the remaining dice in the array. For each recursion, the
sides of the first die will be added to these recursed sums, and a histogram
(map of sums to frequency) will be returned, and when the recursion bottoms
out, the results will all bubble up!

```go
func createHistogramRecursion(dice []int) map[int]int {
	if len(dice) == 0 {
		return map[int]int{0: 1}
	}

	histogram := map[int]int{}
	for i := 1; i <= dice[0]; i++ {
		results := createHistogramRecursion(dice[1:])
		for sum, occurrences := range results {
			histogram[sum+i] += occurrences
		}
	}

	return histogram
}
```

This is fine, but many shops don't like recursion (for good reason!), so what
does a functional implementation of this look like?

```go
func createHistogram(dice []int) map[int]int {
	histogram := map[int]int{0: 1}

	for _, sides := range dice {
		results := map[int]int{}
		for i := 1; i <= sides; i++ {
			for sum, occurrences := range histogram {
				results[sum+i] += occurrences
			}
		}
		histogram = results
	}

	return histogram
}
```

Here, we start with the "bottomed out" histogram, then for each of the dice, we
assemble a `result` where we take that histogram, and add each value in it to
the possible sides on the current die, then set the new histogram value to that
result. It's essentially the recursion flipped upside-down!

But where is this "bottomed out" map coming from? It means "the sum of 0 has
occurred 1 time"... but the sum of zero is impossible! Why don't we say that
the sum of 0 occurred 0 times? The answer lies in the fact that, for each new
sum, we are adding the *previous number of encountered occurrences of that sum*
to the `results`.

Let's run through an example, keeping it simple! Two dice, 2 sides each. We
will end up adding the following possible combinations:

```txt
1 + 1 = 2
1 + 2 = 3
2 + 1 = 3
2 + 2 = 4
```

In the first iteration of the outer loop (the first die), we create an empty
`results` map, then iterate through each of the values (1, 2) in the inner
loop. On the first inner loop iteration, we loop through the `histogram`
members, which are the "bottomed out" `{0: 1}`, and for each key in the
histogram (in this case, just 0), we add it to the value of the current side
(1). We set that sum as a key in the `results` array, then we add the count
already in the histogram (1) to the existing count for the sum in the `results`
map (0), and set this as the new value for that key in the results map. Once we
have done this for all existing possible sums in the histogram, and for all
sides, we override the `histogram` map with our new `results`, then move on to
the next die. This can be represented in a table to visualize the iterations:

| Dice | Side | Histogram | Sum | Previous occurrences in `histogram`* | Previous occurrences in `results` | New `results`       | New `histogram`    |
| ---- | ---- | --------- | --- | ------------------------------------ | --------------------------------- | ------------------- | ------------------ |
| 0    | 0    | 0         | -   | -                                    | -                                 | {}                  | {0: 1}             |
| 1    | 1    | 0         | 1   | 1                                    | 0                                 | {1: 1}              | {0: 1}             |
| 1    | 2    | 0         | 2   | 1                                    | 0                                 | {1: 1, 2: 1}        | {0: 1}             |
| -    | -    | -         | -   | -                                    | -                                 | {}                  | {1: 1, 2: 1}       |
| 2    | 1    | 1         | 2   | 1                                    | 0                                 | {2: 1}              | {1: 1, 2: 1}       |
| 2    | 1    | 2         | 3   | 1                                    | 0                                 | {2: 1, 3: 1}        | {1: 1, 2: 1}       |
| 2    | 2    | 1         | 3   | 1                                    | 1                                 | {2: 1, 3: 2}        | {1: 1, 2: 1}       |
| 2    | 2    | 2         | 4   | 1                                    | 0                                 | {2: 1, 3: 2, 4: 1}  | {1: 1, 2: 1}       |
| -    | -    | -         | -   | -                                    | -                                 | {}                  | {2: 1, 3: 2, 4: 1} |

\* More accurately (and VERY importantly), this is the previous occurrences of
current key in the `histogram` (column 3). This is what will be added to the
occurrences found in the current `results` key.

Another way to think about column 5 above: let's say we've already calculated
six possible combinations of dice rolls which sum to a value of 10. We could
imagine that for each of those six existing combinations, we need to consider
what adding a given side of another die (say a value of 2) would produce. Well,
it will always produce a sum of 12, so there are six instances where the new
sum could be 12! Of course, the next sum in the existing histogram might be 11,
encountered three times in previous combinations. And if we roll a value of 1
on our "current" die and add it, then we'd have three *more* possible sums
which add to 12, so we'd need to add these with our previous results to get a
new number of possible occurrences (nine possible occurrences of a sum of 12).
Then replace the histogram with our new results once we've done this for all
sides of the current die.

**This is exactly why we start with 1 "occurrence" of a sum of 0.** Starting
with 0 occurrences would mean adding 0 to the count every time we encountered
that new sum. This is true whether we start with that *before* iterating, or if
it is the value returned when the recursion "bottoms out".

Maybe that's even less clear, or maybe your brain works like mine :) either way
that's enough of that!

On to the important stuff: which implementation is easier to understand? Even
though my first instinct was to use recursion, I'd argue the non-recursion
implementation is easier to understand. It works in similar ways, but is easier
to conceptualize (it iterates forward, and you don't have to "unwind" the
recursion as it bubbles up). But which is faster?!

Turns out, the non-recursion is also ~7455x faster!

```txt
goos: linux
goarch: amd64
pkg: dice
BenchmarkDiceRollFunctional-12           226     5084751 ns/op
BenchmarkDiceRollRecursive-12              1  37906991525 ns/op
PASS
ok    dice  39.067s
```

This makes it the clear winner, in my opinion. This is not to say that there
aren't cases where recursion is better--if a recursive function is in fact
easier to understand than its non-recursive equivalent, I may be more likely to
pick it. This is because code is more often read than it is written, and most
of the time there is very little need to nitpick over performance differences
which are usually miniscule and don't matter in real world usage.  Of course in
this case the difference is dramatic: it's the difference between the CLI
taking 0.005 seconds to complete, or taking 37.9 seconds! But that's also for a
pretty wild scenario (including a 1000 sided die!) that is rather unlikely for
the intended use of the application. Understanding the use case is important!

Generally, I think recursion should be avoided, or at the *very* least you
should evaluate the non-recursive equivalent before deciding.

There is absolutely no reason I should have gone so deep here. I don't even
play board games like that. But the thought popped into my head and I got
curious, and I just can't help myself when that happens!

    #recursion #histogram #golang #performance
