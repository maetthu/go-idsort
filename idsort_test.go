package idsort_test

import (
	"testing"

	"github.com/maetthu/go-idsort"
	"github.com/stretchr/testify/assert"
)

func TestIDsort(t *testing.T) {
	assert.Equal(t, []int{2, 1, 3}, idsort.IDsort([]int{2, 1, 3}), "Should be equal")
	assert.NotEqual(t, []int{1, 2, 3}, idsort.IDsort([]int{1, 3, 2}), "Should not be equal")
}
