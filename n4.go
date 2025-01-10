package main

import (
	"fmt"
	"time"

	"github.com/MyTempoESP/serial"
)

type SerialForth struct {
	config *serial.Config
	port   *serial.Port
}

func (SerialForth) getBytes(in string) (fixed string) {

	for _, c := range in {

		fixed = fmt.Sprintf("%s %d", fixed, c)
	}

	fixed = fmt.Sprintf("%s %d", fixed, len(in))

	return
}

func NewSerialForth() (forth SerialForth, err error) {

	// Configure the serial port
	forth.config = &serial.Config{
		Name: "/dev/ttyUSB0", // Replace with your Arduino port (e.g., COM3 on Windows)
		Baud: 115200,         // Match the baud rate of your Arduino sketch
	}

	// Open the serial port
	forth.port, err = serial.OpenPort(forth.config)

	// Allow time for Arduino to reset
	time.Sleep(2 * time.Second)

	// defining the API functions
	err = forth.Run("BYE")
	err = forth.Run(": a 1 API ;")
	err = forth.Run(": m 2 API ;")
	err = forth.Run(": d 3 API ;")

	// shorthand for end of line
	err = forth.Run(": $ 15 ;")

	return
}

func (forth *SerialForth) Close() {

	forth.port.Close()
}

func (forth *SerialForth) Run(msg string) (err error) {

	_, err = forth.port.Write(append([]byte(msg), '\n'))

	time.Sleep(500 * time.Millisecond)

	return
}
