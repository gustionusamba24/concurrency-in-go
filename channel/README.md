# Channels
Channels are a typed, [thread-safe queue](https://en.wikipedia.org/wiki/Thread_safety). Channels allow different goroutines to communicate with each other.

## Create a Channel
Like maps and slices, channels must be created before use. They also use the same `make` keyword:
```go
ch := make(chan int)
```

## Send Data to a Channel
```go
ch <- 95 // Send 95 to channel ch
```
The `<-` operator is called the _channel operator_. Data flows in the direction of the arrow. This operation will [block](https://en.wikipedia.org/wiki/Blocking_(computing)) until another goroutine is ready to receive the value.

## Receive Data from a Channel
```go
v := <-ch // Receive from channel ch and assign to v
```
This reads and removes a value from the channel and saves it into the variable `v`. This operation will block until a value in the channel to be read. Channels are First-In-First-Out (FIFO), meaning values are received in the same order they are send.

## Reference Type
Like maps and slices, channels are reference types. Changes made inside a function affect the original.
```go
func send(ch chan int) {
    ch <- 99
}

func main() {
    ch := make(chan int)
    go send(ch)
    fmt.Println(<-ch) // Prints 99
}
```

## Blocking and Deadlocks
A [deadlock](https://yourbasic.org/golang/detect-deadlock/#:~:text=yourbasic.org%2Fgolang,look%20at%20this%20simple%20example.) is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

## Assignment
Run the program. You'll see that it deadlocks and never exits. The `sendIsOld` function is trying to send on a channel, but no other goroutines are running that can accept the value from the channel. 

Fix the deadlock by spawning a goroutine to send the "is old" values.