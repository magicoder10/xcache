# xCache

Convenient Golang caching library

## Getting Started

### Installation

```
$ go get github.com/byliuyang/xcache
```

### Usage

```go
import "github.com/byliuyang/xcache"

// Initialize buffer with length of 3
// LRU represents Least Recently Used replacement policy
buffer := NewLRUBuffer(3)
cache := NewMemCache(buffer)

cache.Set(1, 4)
value, err := cache.Get(1)
```


## Author
- **Harry Liu** - *Initial work* - [byliuyang](https://github.com/byliuyang)

## License
This library is maintained under the MIT License.