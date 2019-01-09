# Part 4 Results
The expected output from this calculation is exactly 0, since we decrement `i` just as many times as we increment it. However, due to the incrementation operation not being atomic, we have a race condition, resulting in all kinds of strange outputs, raraly ever equal to 0. 

## C implementation
The C implementation, using regular threads results in strange numbers both negative and positive. Some example executions:
```bash
$ ./foo
The magic number is: 41869
$ ./foo
The magic number is: -12898
$ ./foo
The magic number is: 0
$ ./foo
The magic number is: 105929

```

## Python implementation
The python implementation also uses regular threads and the same results as before occur.

## Go implementation
The go implementation on the other hand uses coroutines (which in go are called goroutines). These are not threads per say, but are executed in different threads as handled by the runtime. Interestingly enough, this results in the expected `0` each time the program is run. However, when running the program with the race detector (with `go run -race foo.go`), a data race is indeed detected, the strange numbers _do_ appear! Why the results depend on weather we're using the race detector or not, I don't know. I'm not that familiar with the race detector. But I hypothesize that the race condition is indeed present when running the program normally, it just doesn't show itself very often. The race detector however tries provoke this exact behavior and thus results in the result we see here.

