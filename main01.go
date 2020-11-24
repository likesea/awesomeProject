package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

const timeoutSeconds = 5 * time.Second

// There are 4 channels that we are gonna use: 3 unbuffered and 1 buffered of 1.
var (
	// sigChan receives operating signals.
	// This will allow us to send a Ctrl-C to shut down our program cleanly.
	sigChan = make(chan os.Signal, 1)

	// timeout limits the amount of time the program has.
	// We really don't want to receive on this channel because if we do, that means something bad
	// happens, we are timing out and we need to kill the program.
	timeout = time.After(timeoutSeconds)

	// complete is used to report processing is done.
	// This is the channel we want to receive on. When the Goroutine finish the job, it will signal
	// to us on this complete channel and tell us any error that occurred.
	complete = make(chan error)

	// shutdown provides system wide notification.
	shutdown = make(chan struct{})
)

func main01() {
	log.Println("Starting process")
	signal.Notify(sigChan, os.Interrupt)
	log.Println("Launching processors")
	go processor(complete)
ControlLoop:
	for {
		select {
		case <-sigChan:
			log.Println("OS interrupt")
			close(shutdown)
			sigChan = nil
		case <-timeout:
			log.Println("Timeout - Killing Program")
			os.Exit(1)
		case err := <-complete:
			log.Printf("Task Completed: Error[%s]", err)
			// We are using a label break here.
			// We put one at the top of the for loop so the case has a break and the for has a
			// break.
			break ControlLoop
		}
	}
	log.Println("Process Ended")
}

func processor(complete chan <- error){
	log.Println("Processor - Starting")
	var err error
	defer func() {
		if r:=recover();r!=nil{
			log.Println("Processor - panic",r)
		}
		complete <-err
	}()
	err = doWork()
	log.Println("Processor -Completed")
}
func doWork() error {
	log.Println("Processor - Task 1")
	time.Sleep(2 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("Early Shutdown")
	}

	log.Println("Processor - Task 3")
	time.Sleep(1 * time.Second)

	return nil
}
func checkShutdown() bool {
	select {
	case <-shutdown:
		// We have been asked to shutdown cleanly.
		log.Println("checkShutdown - Shutdown Early")
		return true

	default:
		// If the shutdown channel was not closed, presume with normal processing.
		return false
	}
}