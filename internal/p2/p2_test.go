package p2_test

import (
	"testing"

	"github.com/simonasr/advent21/internal/input"
	"github.com/simonasr/advent21/internal/p2"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := input.ReadReport("p2_test.txt")
	result := p2.Part1(input)

	require.Equal(t, 150, result)
}

func TestPart2(t *testing.T) {
	input := input.ReadReport("p2_test.txt")
	result := p2.Part2(input)

	require.Equal(t, 900, result)
}
