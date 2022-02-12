package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestArray(t *testing.T) {
	array := TransferKey([]string{"I", "am", "stupid", "and", "weak"})
	assert.NotEqual(t, array[2], "stupid")
}

func TestDeleteItem(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := DeleteItem(numbers, 1)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, 1, result[0])
	assert.Equal(t, 3, result[1])
}
