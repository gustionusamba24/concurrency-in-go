# Channels Review
Here are a few extra things you should understand about channels from [Dave Cheney's awesome article](https://dave.cheney.net/2014/03/19/channel-axioms).

## A Declared but Uninitialized Channel is Nil Just Like a Slice
```go
var s []int       // s is nil
var c chan string // c is nil

var s = make([]int, 5) // s is initialized and not nil
var c = make(chan int) // c is initialized and not nil
```

## A Send to a Nil Channel Blocks Forever
```go
var c chan string        // c is nil
c <- "let's get started" // blocks
```

## A Receive from a Nil Channel Blocks Forever
```go
var c chan string // c is nil
fmt.Println(<-c)  // blocks
```

## A Send to a Closed Channel Panics
```go
var c = make(chan int, 100)
close(c)
c <- 1 // panic: send on closed channel
```

## A Receive from a Closed Channel Returns the Zero Value
```go
var c = make(chan int, 100)
close(c)
fmt.Println(<-c) // 0
```
Any buffered values are received first. Once the buffer is empty, receives return the zero value.