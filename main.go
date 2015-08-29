package main

import (
	"fmt"
	"os"
    "time"
	"os/signal"
    "github.com/cbess/filewatch/watcher"
)

func main() {
	fmt.Println("Watching...")

    fw := watcher.New("/Users/cbess/Documents/test.txt")
    fw.Watch()

    fw = watcher.New("/Users/cbess/Documents/test2.txt")
    fw.Watch()

	// watch until program exits
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
    <-sigChan // wait for interrupt

    fmt.Println("Done")
}
