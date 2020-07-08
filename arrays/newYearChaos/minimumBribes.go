package newyearchaos

import (
	"fmt"
)

func MinimumBribes(q []int32) {
	var bribes int
	for i, initial := range q {
		current := i + 1
		distance := int(initial) - current
		if distance > 2 {
			fmt.Println("Too chaotic")
			return
		}
		if distance > 0 {
			bribes += distance
		}
	}
	fmt.Println(bribes)
}
