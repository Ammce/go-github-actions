package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumTwoNumbers(t *testing.T) {
	sum := SumTwoNumbers(5, 10)
	var expectedResult int32 = 115
	assert.Equal(t, expectedResult, sum)
}
