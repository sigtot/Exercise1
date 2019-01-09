// Use `go run foo.go` to run your program

package main

import (
    "fmt"
    "runtime"
    "sync"
)

var i = 0

func incrementing() {
    for j := 0; j < 1000000; j++ {
        i++
    }
}

func decrementing() {
    for j := 0; j < 1000000; j++ {
        i--
    }
}

func main() {
    var wg sync.WaitGroup
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
	                                    // Try doing the exercise both with and without it!
    wg.Add(2) // Add 2 since we're spawning 2 goroutines
    go func() {
        defer wg.Done()
        incrementing()
    }()
    go func() {
        defer wg.Done()
        decrementing()
    }()

    fmt.Println("The magic number is:", i)
}
