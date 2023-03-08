## Passing arguments to ... parameters

Variadic function means the func has the final parameter with form ...T

If function f has a final parameter p of type ...T, then within f the type of p is equivalent to type []T

For example:

```go
func Greeting(prefix string, who ...string)

Greeting("hello") //who = nil
Greeting("hello", "Kai", "Nguyen") //who = []string{"Kai", "Nguyen"}
```

If the final argument is assignable to a slice type []T and is followed by ..., it is passed unchanged as the value for a ...T parameter. In this case no new slice is created.

```go
s := []string{"Kai", "Nguyen"}

Greeting("hello", s...)

```
