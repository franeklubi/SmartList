# SmartList #

[![GoDoc](https://godoc.org/github.com/franeklubi/SmartList?status.svg)](https://godoc.org/github.com/franeklubi/SmartList)

SmartList provides an easy way to manage slices of vars of any type.

The most prominent feature of SmartList is a combination of Remove() and Execute().

    * Remove() enables user to specify indexes to-be-removed in the future
    * Execute() removes items of the indexes specified earlier

# Usage example: #
```go
package main

import (
    . "github.com/franeklubi/SmartList"
    "fmt"
)

func main() {
    // define a list
    var sl SList

    // populate the list
    sl.Put(1, 5, 4, "hi!", "multitype stuff!")
    fmt.Println(sl.ToList())
    // output: [1 5 4 hi! multitype stuff!]

    // remove and execute
    sl.Remove(2, 1, 4)
    sl.Execute()

    fmt.Println(sl.ToList())
    // output: [1 hi!]
    fmt.Println(sl.Get(1))
    // output: hi!
}
```
