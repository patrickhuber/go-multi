package multi_test

import (
	"testing"

	"github.com/patrickhuber/go-multi"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	m := multi.New[int](10, 10, 10)
	for i := 0; i < len(m.Values); i++ {
		m.Values[i] = 1
	}
	sum := multi.Sum(m)
	require.Equal(t, 1000, sum)
}
