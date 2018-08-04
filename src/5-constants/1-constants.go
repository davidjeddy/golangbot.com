package main

import (  
    "fmt"
    "math"
)

func main() {  
    fmt.Println("Hello, playground")
    var a = math.Sqrt(4)//allowed
    const b = math.Sqrt(4)//not allowed
}
