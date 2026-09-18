# Select in Golang
Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A `select` statement is used to listen to multiple channels at the same time. It is similar to a `switch` statement but for channels.
```go
select {
case i, ok := <-chInts:
    if ok {
        fmt.Println(i)
    }
case s, ok := <-chStrings:
    if ok {
        fmt.Println(s)
    }
}
```
The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time, the one is chosen randomly. In the example above, the `ok` variable is a boolean. It is `true` while the channel is open, and `false` when the channel is closed by the sender.

## Assignment
Complete the `logMessages` function. It should:
- Use an infinite `for` loop.
- Use a `select` statement to read from `chEmails` and `chSMS` in the order messages arrive.
- Log messages with `logEmail` and `logSMS`.
- `return` when one of the two channels closes, whichever is first.