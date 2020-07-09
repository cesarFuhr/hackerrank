package newyearchaos

import (
	"fmt"
)

func minimumBribes(q []int32) {
	var bribes int
	for i, initial := range q {
		distance := int(initial) - (i + 1)
		if distance > 2 {
			fmt.Println("Too chaotic")
			return
		}
		var j int
		if (q[i] - 2) >= 0 {
			j = int(q[i] - 2)
		} else if (q[i] - 1) >= 0 {
			j = int(q[i] - 1)
		}
		for ; j < i; j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}
	fmt.Println(bribes)
}
