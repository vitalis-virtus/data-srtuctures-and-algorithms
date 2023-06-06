package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibOne(t *testing.T) {
	actual := fibOne(6)
	require.Equal(t, actual, 8)
}

func TestFibTwo(t *testing.T) {
	actual := fibOne(6)
	require.Equal(t, actual, 8)
}

func TestFibThree(t *testing.T) {
	actual := fibOne(6)
	require.Equal(t, actual, 8)
}
