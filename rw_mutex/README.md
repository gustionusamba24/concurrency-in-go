# RW Mutex
The standard library also exposes a [sync.RWMutex](https://golang.org/pkg/sync/#RWMutex)

In addition to these methods:
- [Lock()](https://pkg.go.dev/sync#RWMutex.Lock)
- [Unlock()](https://pkg.go.dev/sync#RWMutex.Unlock)

The `sync.RWMutex` also has these methods for concurrent reads:
- [RLock()](https://pkg.go.dev/sync#RWMutex.RLock)
- [RUnlock()](https://pkg.go.dev/sync#RWMutex.RUnlock)

The `sync.RWMutex` improves performance in read-intensive processes. Multiple goroutines can safely read from the map simultaneously, as many `RLock()` calls can occur at the same time. However, only one goroutine can hold a `Lock()`, and during this time, all `RLock()` operations are blocked.

## Assignment
Let's update our same code from the last assignment, but this time we can speed it up by allowing readers to read from the map concurrently.

Run the new test suite. You will notice that it hangs forever, and you will need to cancel it.

Update the `val` method to only lock the mutex for reading.