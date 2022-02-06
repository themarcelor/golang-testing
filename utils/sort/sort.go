package sort

import (
	"fmt"
)

func BubbleSort(elements []int) {
	keepWorking := true

        for keepWorking {
		keepWorking = false

		for i:= 0; i < len(elements)-1; i++ {
			// fmt.Printf("comparing %d and %d\n", elements[i], elements[i+1])
			if elements[i] < elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
		// fmt.Printf("bubble sorting so far: %v\n", elements)
	}
	fmt.Println("done")
}
