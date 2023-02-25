# Slices

<p align="center">
    <span>The missing map/filter/reduce (+more) for <a href="https://pkg.go.dev/golang.org/x/exp/slices">golang.org/x/exp/slices</a> (requires Go v1.18+).</span>
    <br><br>
    <a href="https://github.com/felix-kaestner/slices/issues">
        <img alt="Issues" src="https://img.shields.io/github/issues/felix-kaestner/slices?color=29b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/slices/stargazers">
        <img alt="Stars" src="https://img.shields.io/github/stars/felix-kaestner/slices?color=29b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/slices/blob/main/LICENSE">
        <img alt="License" src="https://img.shields.io/github/license/felix-kaestner/slices?color=29b6f6&style=flat-square">
    </a>
    <a href="https://pkg.go.dev/github.com/felix-kaestner/slices">
        <img alt="Stars" src="https://img.shields.io/badge/go-documentation-blue?color=29b6f6&style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/felix-kaestner/slices">
        <img alt="Issues" src="https://goreportcard.com/badge/github.com/felix-kaestner/slices?style=flat-square">
    </a>
    <a href="https://codecov.io/gh/felix-kaestner/slices">
        <img src="https://img.shields.io/codecov/c/github/felix-kaestner/slices?style=flat-square&token=R3OVLPMFB9"/>
    </a>
    <a href="https://twitter.com/kaestner_felix">
        <img alt="Twitter" src="https://img.shields.io/badge/twitter-@kaestner_felix-29b6f6?style=flat-square">
    </a>
</p>

## Quickstart

```go
package main

import (
    "fmt"

    "github.com/felix-kaestner/slices"
)

func main() {
    // Original slice of integers
    numbers := []int{1, 2, 3, 4, 5}
    	
    // Filter only numbers greater than 2
    greaterThanTwo := slices.Filter(s1, func(i int) bool { return i > 2 })
    
    fmt.Println("Numbers > 2:", greaterThanTwo)
    // Prints "Numbers > 2: []int{3, 4, 5}"
}
```

##  Installation

Install with the `go get` command:

```
$ go get -u github.com/felix-kaestner/slices
```

## Contribute

All contributions in any form are welcome! 🙌🏻  
Just use the [Issue](.github/ISSUE_TEMPLATE) and [Pull Request](.github/PULL_REQUEST_TEMPLATE) templates and I'll be happy to review your suggestions. 👍

---

Released under the [MIT License](LICENSE).
