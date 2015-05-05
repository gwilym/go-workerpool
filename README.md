# go-workerpool

A simple, stoppable, concurrent worker pool interface including a function-calling implementation.

You provide the workerpool with a buffered bool channel when started. Workers loop until stopped by calling `Stop`, which will prevent the next loop iteration. When all workers stop, `true` is sent over the channel.

Specifically, the function-calling workerpool requires a `func() bool` worker. The worker will be looped by the pool until it either returns `false` or `Stop` is called on the pool.

I'm still kinda learning Go, so I'm writing little libraries as I figure things out. This was something I used on another project that was easily extracted. Though I wouldn't be the least bit surprised if this repeats someone else's work, or even something I don't know about in the Go standard library.

## Installation

    go get github.com/gwilym/go-workerpool

## Usage by Example

See `example/main.go` for a working example.

