package p3_test

import (
	"testing"

	"github.com/simonasr/advent21/internal/input"
	"github.com/simonasr/advent21/internal/p3"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := input.ReadReport("p3_test.txt")
	result := p3.Part1(input)

	require.Equal(t, 198, result)
}

func TestPart2(t *testing.T) {
	input := input.ReadReport("p3_test.txt")
	oxygen := p3.GetRating(input, "oxygen")
	co2 := p3.GetRating(input, "co2")
	result := p3.Part2(input)

	require.Equal(t, 230, result)
	require.Equal(t, 23, oxygen)
	require.Equal(t, 10, co2)
}
