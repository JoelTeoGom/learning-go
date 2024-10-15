# Concurrency in Go - Notes

In my journey to mastering Go, I've been delving into the fascinating world of concurrency. Here are some key takeaways that I've found particularly enlightening:

## 1. Goroutines

Goroutines are lightweight threads managed by the Go runtime. They are initiated with the `go` keyword, allowing functions to run concurrently.

```go
go count("sheep")
count("fish")
```

This enables multiple functions to execute at the same time without blocking each other.


## 2. WaitGroups

To synchronize goroutines, we can use sync.WaitGroup. It allows us to wait for a collection of goroutines to finish executing before proceeding.

```go
var wg sync.WaitGroup
wg.Add(1)

go func() {
    count("sheep")
    wg.Done()
}()

wg.Wait()

```


## 3. Channels

Channels are used for communication between goroutines. They can be buffered or unbuffered, allowing safe sending and receiving of messages.

```go
c := make(chan string)
go count("sheep", c)

for msg := range c {
    fmt.Println(msg)
}
```
If a channel is not closed, the receiving function can result in a deadlock if it continues to wait for messages that will never come.


## 4. The Select Statement

The select statement provides another level of control when working with multiple channels. It allows us to wait on multiple communication operations.

```go
select {
case msg1 := <-c1:
    fmt.Println(msg1)
case msg2 := <-c2:
    fmt.Println(msg2)
}
```

Without select, using a simple if statement to receive from channels could lead to blocking, as it waits for the first message to arrive.


## 5. Closing Channels

Closing a channel indicates that no more values will be sent. As the sender, it's safe to close a channel when you’re done, but receivers should never close it to avoid panicking.

```go
close(c) // done sending
```

When receiving from a channel, we can check if it is open:
```go
msg, open := <-c
if !open {
    break
}
```
This prevents deadlocks and allows us to gracefully handle channel termination.

## 6. Practical Example

Here’s a simple example illustrating concurrency in Go:

```go
go func() {
    for {
        time.Sleep(time.Millisecond * 500)
        c1 <- "Every 500ms"
    }
}()

for {
    select {
    case msg1 := <-c1:
        fmt.Println(msg1)
    case msg2 := <-c2:
        fmt.Println(msg2)
    }
}

```

Using select prevents blocking on channels, ensuring that we can handle messages as they come in without unnecessary delays.


## Conclusion

Concurrency in Go is a powerful feature that enables developers to write efficient and responsive applications. By utilizing goroutines, channels, and select statements, we can build applications that perform well even under load.

