package blitz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkLengthUniquenessAndRange(t *testing.T, result []int, k int, expectedLength int) {
	assert.Len(t, result, expectedLength)
	uniqueMap := make(map[int]bool)
	for _, num := range result {
		if num < 0 || num >= k {
			t.Errorf("Number %d is out of range [0, %d)", num, k)
		}
		if uniqueMap[num] {
			t.Errorf("Number %d is not unique", num)
		}
		uniqueMap[num] = true
	}
}

func TestGenerateUniqueRandomSubset(t *testing.T) {
	// Test case 1: k = 5, n = 3
	// Should return 3 unique numbers between 0 and 4
	k1, n1 := 5, 3
	result1 := generateUniqueRandomSubset(k1, n1)
	checkLengthUniquenessAndRange(t, result1, k1, n1)

	// Test case 2: k = 10, n = 10
	// Should return 10 unique numbers between 0 and 9 (all of them)
	k2, n2 := 10, 10
	result2 := generateUniqueRandomSubset(k2, n2)
	if len(result2) != n2 {
		t.Errorf("Test Case 2 Failed: Expected length %d, got %d", n2, len(result2))
	}
	checkLengthUniquenessAndRange(t, result2, k2, n2)

	// Test case 3: k = 1, n = 1
	// Should return [0]
	k3, n3 := 1, 1
	result3 := generateUniqueRandomSubset(k3, n3)
	if len(result3) != n3 || result3[0] != 0 {
		t.Errorf("Test Case 3 Failed: Expected [0], got %v", result3)
	}
	assert.Equal(t, []int{0}, result3)

	// Test case 4: k = 2, n = 1
	// Should return 1 unique number between 0 and 1
	k4, n4 := 2, 1
	result4 := generateUniqueRandomSubset(k4, n4)
	if len(result4) != n4 {
		t.Errorf("Test Case 4 Failed: Expected length %d, got %d", n4, len(result4))
	}
	checkLengthUniquenessAndRange(t, result4, k4, n3)
}
