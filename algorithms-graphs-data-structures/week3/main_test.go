package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSequence(t *testing.T) {
	s := newSequence()
	s.Insert(50)
	assert.Equal(t, 50, s.Median())

	s.Insert(5)
	assert.Equal(t, 5, s.Median())

	s.Insert(15)
	assert.Equal(t, 15, s.Median())

	s.Insert(1)
	assert.Equal(t, 5, s.Median())

	s.Insert(3)
	assert.Equal(t, 5, s.Median())

	s.Insert(20)
	assert.Equal(t, 5, s.Median())

	s.Insert(6)
	assert.Equal(t, 6, s.Median())

	s.Insert(4)
	assert.Equal(t, 5, s.Median())

	s.Insert(7)
	assert.Equal(t, 6, s.Median())
}
