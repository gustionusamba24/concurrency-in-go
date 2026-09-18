# Buffered Channels
Channels can optionally be buffered.

## Creating a Channel with a Buffer
You can provide a buffer length as the second argument to `make()` to create a buffered channel.
```go
ch := make(chan int, 100)
```
A buffer allows the channel to hold a fixed number of values before sending blocks. This means sending on a buffered channel only blocks when the buffer is full, and receiving blocks only when the buffer is empty.

## Assignment
We want to be able to send emails in batches. A writing goroutine will write an entire batch of email messages to a buffered channel. Later, once the channel is full, a reading goroutine will read all the messages from the channel and send them out to our clients.

Complete the `addEmailsToQueue` function. It should create a buffered channel with a buffer large enough to store all of the `emails` it's given. It should then write the emails to the channel in order and finally return the channel.