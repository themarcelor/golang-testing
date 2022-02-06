package sort

import (
	"testing"
	"fmt"

	"github.com/tj/assert"
)

func TestBubbleSortOrderDESC( t *testing.T) {
	// Init
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	validateResult(elements, expected, t)
}

func TestBubbleSortAlreadySorted(t *testing.T) {
        // Init
	expected := []int{5, 4, 3, 2, 1}
        elements := []int{5, 4, 3, 2, 1}

        // Execution
        BubbleSort(elements)

        // Validation
        validateResult(elements, expected, t)
}

func validateResult(elements []int, expected []int, t *testing.T) {
        if elements[0] != expected[0] {
                t.Error(fmt.Sprintf("first element should be %d\n", expected[0]))
        }
        if elements[len(elements)-1] != expected[len(expected)-1] {
                t.Error(fmt.Sprintf("last element should be %d\n", expected[len(expected)-1]))
        }

        assert.Equal(t, expected, elements)
        fmt.Println(elements)
}
