# Digital root

A digital root is calculated by summing the digits of a number, then summing the digits of that sum,
and continuing to sum the digits of the resulting sums, until the sum is only one digit.

For instance:
```
Number: 195

1 + 9 + 5 = 14
1 + 4 = 5

Answer: 5
```

Or:
```
Number: 8675309 

8 + 6 + 7 + 5 + 3 + 0 + 9 = 38
3 + 8 = 11
1 + 1 = 2

Answer: 2
```

In a base 10 system, the digital root of `x` can be represented as:
```
f(x) = (x - 1) % 9 + 1
```

But why? Here's some examples to see if there is a pattern:
```
f(12) => 11 % 9 + 1 => 2 + 1 => 3

f(16) => 15 % 9 + 1 => 6 + 1 => 7

f(195) => 194 % 9 + 1 => 5 + 1 => 6
1 + 9 + 5 = 15
1 + 5 = 6

f(8675309) => 8675308 % 9 + 1 => 1 + 1 => 2
8 + 6 + 7 + 5 + 3 + 0 + 9 = 38
3 + 8 = 11
1 + 1 = 2
```

`0` must be a special case because you can't sum positive integers and get `0`

For each group of 9 possible answers (first one includes special case `0` for 10 possible answers):
* `f(x = [0,9]) => x`
* `f(x = [10,18]) => x - 9`
* `f(x = [19,27]) => x - 18`

So `f(x)` is essentially the difference between `x` and its closest *non-inclusive* multiple of 9. 
To make it non-inclusive, we first do `x - 1` before taking `mod 9`. Then simply add that `1` back
afterward!
```)
f(x) => (x - 1) % 9 + 1
```

Put another way, since no base 10 number (except zero) may have a digital root of `0`, and since
`f(10) = 1` in base 10, the digital root answers just cycle:
```
f(1)  = 1
f(2)  = 2
f(3)  = 3
f(4)  = 4
f(5)  = 5
f(6)  = 6
f(7)  = 7
f(8)  = 8
f(9)  = 9
f(10) = 1
f(11) = 2
f(12) = 3
f(13) = 4
f(14) = 5
f(15) = 6
f(16) = 7
f(17) = 8
f(18) = 9
f(19) = 1
f(20) = 2
...
```

    #learn #vocab #tips #algorithms
