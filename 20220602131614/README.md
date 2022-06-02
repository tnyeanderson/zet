# Idempotent operations

An operation or function is idempotent (`eye-dem-POE-tint`) if it can be called or run repeatedly,
over and over again, and return the same result (or have the same effect).

This function is **not** idempotent:
```js
let sum = 0

function increment(amount) {
	sum = sum + amount
  return sum
}
```
> `f(1)` will return `1` the first time, then `2`, then `3`...

This function **is** idempotent:
```js
function increment(sum) {
  return sum + 1
}
```
> `f(1)` will always be `2`

In other words, a function is idempotent if `f(f(x)) = f(x)`, or `f(x) = f(x); f(x)`.

    #vocab #learn
