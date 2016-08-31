package input 

import (
	"fmt"
	"os"
	"reflect"
)

func ReadInput(prompt string, readInto interface{}) {

	satisfied := false

	for !satisfied {
		fmt.Print(prompt)
		var intVal uint
		_, err := fmt.Scanf("%d", &intVal)
		if err != nil {
			fmt.Printf("Error reading input (%s)\n", err)
			continue
		}

		satisfied = true

		switch readInto := readInto.(type) {
		case *byte:
			if intVal > 0xff {
				fmt.Printf("Input too large!")	
				satisfied = false	
			}
			*readInto = byte(intVal)
		case *uint16:
			if intVal > 0xffff {
				fmt.Printf("Input too large!")
				satisfied = false
			}
			*readInto = uint16(intVal)
		default:
			fmt.Fprintf(os.Stderr, "Can only read into byte or uint16 pointers. Got: %s!\n",
				reflect.TypeOf(readInto))
			os.Exit(1)
		}

	}
}
