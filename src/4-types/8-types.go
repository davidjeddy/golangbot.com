package main

import (  
    "fmt"
)

func main() {  
    i := 10
    var j float64 = float64(i) //this statement will not work without explicit conversion
    fmt.Println("j", j)
}
