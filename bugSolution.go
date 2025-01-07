```go
package main

import "fmt"

func main() {
    var m map[string]int
    // Safe way 1: Check if the map is nil
    if m == nil {
        fmt.Println("Map is nil")
    } else {
        fmt.Println(m["a"])
    }
    
    // Safe way 2: Comma ok idiom
    value, ok := m["a"]
    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("Key not found")
    }
    
    // Safe way 3: Initialize the map
    m = make(map[string]int)
    m["a"] = 10
    fmt.Println(m["a"])
}
```