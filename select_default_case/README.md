# Select Default Case
The `default` case in a `select` statement executes immediately if no other channel has a value ready. A `default` case stops the `select` statement from blocking.
```go
select {
case v := <-ch:
    // use v
default:
    // receiving from ch would block
    // so do something else
}
```

## Ignoring Channels
Sometimes you want to ignore a channel's value. You can do this by not binding it to a variable:
```go
select {
case <-ch:
    // event received; value ignored
default:
    // so do something else
}
```
Alternatively, you can use blank identifier `_` to ignore the value:
```go
select {
case _ = <-ch:
    // event received; value ignored
default:
    // so do something else
}
```

## Tickers
- [time.Tick()](https://golang.org/pkg/time/#Tick) is a standard library function that returns a channel that sends a value on a given interval. 
- [time.After()](https://golang.org/pkg/time/#After) sends a value once after the duration has passed. 
- [time.Sleep()](https://golang.org/pkg/time/#Sleep) blocks the current goroutine for the specified duration of time.

The functions take a [time.Duration](https://pkg.go.dev/time#Duration) as an argument. For example:
```go
time.Tick(500 * time.Millisecond)
```
If you don't add `time.Millisecond` (or another unit), it will default to nanoseconds. That's – taking a wild guess here – probably faster than you want it to be.

## Read-Only Channels
A channel can be marked as ready-only by casting it from `chan` to a `<-chan` type. This is useful when you want to ensure that a channel is only used for receiving values and not for sending them. For example:
```go
var ch chan int = make(chan int)
var readOnlyCh <-chan int = ch // readOnlyCh can only receive values from ch

// Another example
func main() {
    ch := make(chan int)
    readCh(ch)
}

func readCh(ch <-chan int) {
    // ch can only be read from
    // in this function
}
```

## Write-Only Channels
The same goes for write-only channels, but the arrow's position moves.
```go
var writeOnlyCh chan<- int = make(chan int)

// Another example
func writeCh(ch chan<- int) {
    // ch can only be written to
    // in this function
}
```

## Assignment
Like all good back-end engineers, we frequently save backup snapshots of the Textio database.

Complete the `saveBackups` function.

It should use a loop to read values from the `snapshotTicker` and `saveAfter` channels simultaneously and continuously.
1. If a value is received from `snapshotTicker`, call `takeSnapshot()`.
2. If a value is received from `saveAfter`, call `saveSnapshot()` and `return` from the function: you're done.
3. If neither channel has a value ready, call `waitForData()` and then [time.Sleep()](https://pkg.go.dev/time#example-Sleep) for 500 milliseconds. After all, we want to show in the logs that the snapshot service is running.