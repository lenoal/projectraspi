package main

import (
        "fmt"
        "github.com/stianeikeland/go-rpio"
        "os"
        "time"
)

var (
        pin = rpio.Pin(17)
	xpin = rpio.Pin(22)
)

func main() {
        // Open and map memory to access gpio, check for errors
        if err := rpio.Open(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        // Unmap gpio memory when done
        defer rpio.Close()

        // Set pin to output mode
        pin.Output()
	xpin.Input()


        // Toggle - infinite loop
	for xpin.High(){
                pin.Toggle()
                time.Sleep(500*time.Millisecond)
	}
	for xpin.Low() {
		pin.Low()
	}

}
