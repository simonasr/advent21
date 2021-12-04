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
