package main

import (
	"fmt"
	"time"

	"github.com/MyTempoESP/serial"
)

type Arduino struct {
	config *serial.Config
	port   *serial.Port
}

func FixString(in string) (fixed string) {

	for _, c := range in {

		fixed = fmt.Sprintf("%s %d", fixed, c)
	}

	return
}

func NewArduino() (ino Arduino, err error) {

	// Configure the serial port
	ino.config = &serial.Config{
		Name: "/dev/ttyUSB0", // Replace with your Arduino port (e.g., COM3 on Windows)
		Baud: 115200,         // Match the baud rate of your Arduino sketch
	}

	// Open the serial port
	ino.port, err = serial.OpenPort(ino.config)

	// Allow time for Arduino to reset
	time.Sleep(2 * time.Second)

	// defining the API functions
	err = ino.Send("BYE\n")
	err = ino.Send(": a 1 API ;\n")
	err = ino.Send(": m 2 API ;\n")
	err = ino.Send(": d 3 API ;\n")

	return
}

func (ino *Arduino) Close() {

	ino.port.Close()
}

func (ino *Arduino) Append(msg string) (err error) {

	err = ino.Send(fmt.Sprintf("%s %d a\n", FixString(msg), len(msg)))

	return
}

func (ino *Arduino) Move(x uint8, y uint8) (err error) {

	err = ino.Send(fmt.Sprintf("%d %d m\n", y, x))

	return
}

func (ino *Arduino) Clear(upto uint8) (err error) {

	err = ino.Send(fmt.Sprintf("%d d\n", upto))

	return
}

func (ino *Arduino) Send(msg string) (err error) {

	_, err = ino.port.Write([]byte(msg))

	time.Sleep(500 * time.Millisecond)

	return
}
