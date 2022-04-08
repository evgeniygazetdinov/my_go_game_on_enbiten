package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Message int

const (
    CacheLimit   = 100
    CacheTimeout = 5 * time.Second
)

func main() {
    input := make(chan Message, CacheLimit)

    go poll(input)
    generate(input)
}

// poll checks for incoming messages and caches them internally
// until either a maximum amount is reached, or a timeout occurs.
func poll(input <-chan Message) {
    cache := make([]Message, 0, CacheLimit)
    tick := time.NewTicker(CacheTimeout)

    for {
        select {
        // Check if a new messages is available.
        // If so, store it and check if the cache
        // has exceeded its size limit.
        case m := <-input:
            cache = append(cache, m)

            if len(cache) < CacheLimit {
                break
            }

            // Reset the timeout ticker.
            // Otherwise we will get too many sends.
            tick.Stop()

            // Send the cached messages and reset the cache.
            send(cache)
            cache = cache[:0]

            // Recreate the ticker, so the timeout trigger
            // remains consistent.
            tick = time.NewTicker(CacheTimeout)

        // If the timeout is reached, send the
        // current message cache, regardless of
        // its size.
        case <-tick.C:
            send(cache)
            cache = cache[:0]
        }
    }
}

// send sends cached messages to a remote server.
func send(cache []Message) {
    if len(cache) == 0 {
        return // Nothing to do here.
    }

    fmt.Printf("%d message(s) pending\n", len(cache))
}

// generate creates some random messages and pushes them into the given channel.
//
// Not part of the solution. This just simulates whatever you use to create
// the messages by creating a new message at random time intervals.
func generate(input chan<- Message) {
    for {
        select {
        case <-time.After(time.Duration(rand.Intn(100)) * time.Millisecond):
            input <- Message(rand.Int())
        }
    }
}