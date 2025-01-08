package main

import (
	"fmt"
	"log"
)

func main() {

	ino, err := NewArduino()

	if err != nil {
		log.Fatalf("Error opening arduino: %v", err)
	}

	defer ino.Close()

	//	ino.Append("tags: ")
	//
	//	for i := range uint8(30) {
	//		ino.Move(0, 0)
	//		ino.Clear((i % 16) - 1)
	//		ino.Move(i%16, 0)
	//		ino.Append("O")
	//	}

	/* Annotated code:

	: SC1                    # Scene1
		30 FOR I               # FOR I = 30 to 1
			0 0 m                #   LCD.move(0, 0)
			I ASC 16 MOD 1 - d   #   LCD.clear([I (in ASCending order) MOD 16] - 1)
			I ASC 16 MOD 0 SWP m #   LCD.move(I (in ASCending order) MOD 16, 0)
			65 1 a               #   LCD.print('A')
			500 DLY              #   delay(500)
		NXT                    # NEXT I
	;

	*/

	//ino.Send(`: a 1 API ; : m 2 API ; : d 3 API ;\n`)

	//ino.Send(": ASC 30 SWP - ;\n") // trick to invert the order of the FOR to ASCending
	//ino.Send(": SC1 30 FOR I ASC 0 0 2 API I ASC 16 MOD 1 - 3 API I ASC 16 MOD 0 SWP 2 API 65 1 1 API 500 DLY NXT ;\n")

	/* Display 'tags: 55555':

	: SC1
		$FixString("tags: 55555")
		$len("tags: 55555")

		0 0 m
		a
	;

	*/

	ino.Send(fmt.Sprintf(": SC1 %s %d 0 0 m a ;\n", FixString("tags: 55555"), len("tags: 55555")))

	// running the scene
	ino.Send("SC1\n")
}
