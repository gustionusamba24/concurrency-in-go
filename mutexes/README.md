# Mutexes in Golang
Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.

Go's standard library provides a built-in implementation of a mutex with the [sync.Mutex](https://pkg.go.dev/sync#Mutex) type and its two methods:
- [Lock()](https://golang.org/pkg/sync/#Mutex.Lock)
- [Unlock()](https://golang.org/pkg/sync/#Mutex.Unlock)

We can protect a block of code by surrounding it with a call to `Lock()` and `Unlock()` as shown on the `protected()` function below.

It's good practice to structure the protected code within a function so that `defer` can be used to ensure that we never forget to unlock the mutex.
```go
func protected(){
    mu.Lock()
    defer mu.Unlock()
    // the rest of the function is protected
    // any other calls to `mu.Lock()` will block
}
```
In the example above, `mu` is a mutex variable. In your assignment, the mutex is a field on your `safeCounter` struct (`sc.mu`), so you will call `sc.mu.Lock()` and use `sc.mu.Unlock()` to protect the critical sections.

Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.

## Maps are not [Thread-Safe](https://en.wikipedia.org/wiki/Thread_safety)
Maps are not safe for concurrent use. If you have multiple goroutines accessing the same map, and at least one of them is writing to the map, you must lock your maps with a mutex.

## Assignment
We send emails across many different goroutines at Textio. To keep track of how many we've sent to a given email address, we use an in-memory map.

Our `safeCounter` struct is unsafe. Update the `inc()` and `val()` methods so that they use the safeCounter's mutex to ensure that the map is not accessed by multiple goroutines at the same time.

## Note: Wasm is Single-Threaded
Now, it's worth pointing out that our execution engine on Boot.dev uses web assembly to run the code you write in your browser. Web assembly is [single-threaded](https://en.wikipedia.org/wiki/Thread_(computing)), which awkwardly means that maps are thread-safe in web assembly. I've simulated a [multithreaded](https://en.wikipedia.org/wiki/Multithreading_(computer_architecture)) environment with the `slowIncrement` and `slowVal` functions.

In reality, any Go code you write may or may not run on a single-core machine, so it's always best to write your code so that it is safe no matter which hardware it runs on.