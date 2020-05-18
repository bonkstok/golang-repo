package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTSTP, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)

	go func() {
		for {
			s := <-sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println("Got a SIGINT")
				done <- true
			case syscall.SIGTSTP:
				fmt.Println("Someone pressed ctrl+Z")
				done <- true
			case syscall.SIGKILL:
				fmt.Println("F, got killed bro")
			case syscall.SIGHUP:
				fmt.Println("SIGUP received, reloading config..")
				done <- true
			case syscall.SIGTERM:
				fmt.Println("Will try to stop nicely")
			}
		}
	}()
	fmt.Println("Program blocked until a signal is caught")
	fmt.Println("PID is:", os.Getpid())
	<-done
	fmt.Println("Out of here")
}
