# Slices

<p align="center">
    <span>The missing map/filter/reduce for <a href="https://pkg.go.dev/golang.org/x/exp/slices">golang.org/x/exp/slices</a></span>
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
    <!-- <a href="https://codecov.io/gh/felix-kaestner/slices">
        <img src="https://img.shields.io/codecov/c/github/felix-kaestner/slices?style=flat-square&token=KK7ZG7A90X"/>
    </a> -->
    <a href="https://twitter.com/kaestner_felix">
        <img alt="Twitter" src="https://img.shields.io/badge/twitter-@kaestner_felix-29b6f6?style=flat-square">
    </a>
</p>

## Quickstart

```go
package main

import "github.com/felix-kaestner/slices"

func main() {
    // Original slice
    s1 := []int{1, 2, 3, 4, 5}
    	
    // Filter only numbers greater than 2
    s2 := slices.Filter(s1, func(i int) bool { return i > 2 }) 
    // s2 == []int{3, 4, 5}
    
    // Filter only odd numbers in place
    // This will modify the underlying array of slice s2! Don't use s2 afterwards.
    s3 := slices.FilterInPlace(s2, func(i int) bool { return i > 2 }
    // s3 == []int{3, 5}
    
    // Map adds 1 to each element
    s4 := slices.Map(s3, func(i int) int { return i + 1 })
    // s4 == []int{4, 6}
    
    // Reduce sums all numbers in the slice
    sum := slices.Reduce(s4, func(sum, i int) int { return sum + i })
    // sum == 10
}
```

##  Installation

Install with the `go get` command:

```
$ go get -u github.com/felix-kaestner/slices
```

## Contribute

All contributions in any form are welcome! 🙌  
Just use the [Issue](.github/ISSUE_TEMPLATE) and [Pull Request](.github/PULL_REQUEST_TEMPLATE) templates and 
I will be happy to review your suggestions. 👍

## Support

Any kind of support is well appreciated! 👏  
If you want to tweet about the project, make sure to tag me [@kaestner_felix](https://twitter.com/kaestner_felix). You can also support my open source work on [GitHub Sponsors](https://github.com/sponsors/felix-kaestner). All income will be directly invested in Coffee ☕!

## Cheers ✌
