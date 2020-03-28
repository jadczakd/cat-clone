package main

import (
	"os"
	"io"
	"os/signal"
	"syscall"
)

func main() {
	signal.Ignore(syscall.SIGPIPE)
	io.Copy(os.Stdout, os.Stdin)
}
