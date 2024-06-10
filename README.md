# ⚠️ This Version is no longer maintained.

# gocache

[![Go Tests](https://github.com/grqphical/gocache/actions/workflows/tests.yml/badge.svg)](https://github.com/grqphical/gocache/actions/workflows/tests.yml)

A simple, threadsafe cache that can be used within your Golang apps as a replacement for external caches.

## Installation

Simply run

```bash
$ go get github.com/grqphical/gocache
```

## Usage

GoCache has two types of caches: `GoCache` and `GenericCache`.

### GoCache

`GoCache` is a cache with `string` keys and `interface{}` values.

```go
package main

import (
    "github.com/grqphical/gocache"
    "fmt"
)

func main() {
    cache := gocache.NewGoCache()

    cache.Set("foo", "bar")

    fmt.Println(cache.GetString("foo"))
}
```

### GenericCache

`GenericCache` uses generics in order to give the user more control over the data that goes in the store. When you create a `GenericCache`, you pass in the
types as `K` and `V` in `NewGenericCache`

```go
package main

import (
    "github.com/grqphical/gocache/generic"
    "fmt"
)

func main() {
    cache := gocache.NewGenericCache[int, string]()

    cache.Set(0, "bar")

    fmt.Println(cache.Get(0))
}
```

## License

gocache is licensed under the MIT License
