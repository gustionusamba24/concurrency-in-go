# Range
Similar to slices and maps, channels can be ranged over.
```go
for item := range ch {
    // item is the next value received from the channel
}
```
This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.

## Assignment
It's that time again, Textio is hiring, and we've been assigned to do the interview. The [Fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_number) is Textio's interview problem at choice. We've been tasked with building a small toy program we can use in the interview.

Complete the `concurrentFib` function. It should:
- Create a new channel of `int`s.
- Call `fibonacci` concurrently.
- Use a `range` loop to read from the channel and append the values to a slice.
- Return the slice.

> The `fibonacci` function is the **`producer`**. It runs in its own goroutine and sends numbers into the channel. The **`concurrentFib`** function is the `consumer`. It reads all the numbers from that channel using a `range` loop and aggregates them.