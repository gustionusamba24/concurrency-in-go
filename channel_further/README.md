# Channels Further
In the previous lesson, we saw how we can receive values from channels like this:
```go
v := <-ch
```
However, sometimes we don't care what is passed through a channel. We only care when and if something is passed. In that situation, we can block and wait until something is sent on a channel using the following syntax:
```go
<-ch
```
This will block until it pops a single value off the channel, then continue, discarding the item.

In cases like this, [empty structs](https://dave.cheney.net/2014/03/25/the-empty-struct) are often used as a [unary](https://en.wikipedia.org/wiki/Unary_operation) value so that the sender communicates that this is only a "signal" and not data is meant to be used.

Here is an example:
```go
func downloadData() chan struct{} {
    downloadDoneCh := make(chan struct{})
    
    go func() {
        fmt.Println("Downloading data file...")
        time.Sleep(2 * time.Second) // simulate download time
    
        // after the download is done, send a "signal" to the channel
        downloadDoneCh <- struct{}{}
    }()
    
    return downloadDoneCh
}

func processData(downloadDoneCh chan struct{}) {
    // any code here can run normally
    fmt.Println("Preparing to process data...")
    
    // block until `downloadData` sends the signal that it's done
    <-downloadDoneCh
    
    // any code here can assume that data download is complete
    fmt.Println("Data download complete, starting data processing...")
}

processData(downloadData())
// Preparing to process data...
// Downloading data file...
// Data download complete, starting data processing...
```

## Assignment
Our Textio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code. It never exits! The channel passed to `waitForDBs` stays blocked, because it's only popping the first value off the channel.

Fix the `waitForDBs` function. It should pause execution until it receives a token for each of the `numDBs` databases from the `dbChan` channel. Each time `waitForDBs` reads a token, the `getDBsChannel` goroutine will print a message to the console for you. Look at the test code to see how these functions are used so you can understand the control flow.