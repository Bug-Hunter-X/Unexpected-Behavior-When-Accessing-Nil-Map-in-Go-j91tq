```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, but return the zero value of int (0)
    // However, if you try to access a value from a nil map using a method like this:
    //fmt.Println(m["a"])
    // it will panic with a runtime error: "runtime error: map index out of range"
}
```