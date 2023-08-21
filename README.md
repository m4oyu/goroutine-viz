# goroutine-viz

A simple goroutine visualization tool

## Usage

```go
package main

import (
	"time"

	viz "github.com/m4oyu/goroutine-viz"
)

func main() {
	viz.WatchGoroutine("BREAKPOINT1")
	go func() {
		viz.WatchGoroutine("BREAKPOINT2")
	}()

    <-time.After(time.Second * 1)
}
```

```bash
$ go run main.go
BREAKPOINT1
goroutine 1 (main goroutine)

BREAKPOINT2
goroutine 1 (main goroutine)
└── goroutine 6 (created by main.main)
```

## Installation

```bash
$ go install github.com/m4oyu/goroutine-viz@latest
```

