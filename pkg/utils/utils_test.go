package utils

import "testing"

func TestExtractNums(t *testing.T) {
	nums := ExtractNums("hello5 bobby9486 (6,5)")
	expectedList := []int{5, 9486, 6, 5}

	for idx := range nums {
		if nums[idx] != expectedList[idx] {
			t.Logf("Expceted: %d, Got: %d", expectedList[idx], nums[idx])
			t.Fail()
		}
	}
}
