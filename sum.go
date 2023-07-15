package multi

import (
	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Integer](m *Multi[T]) T {
	var sum T = 0
	for _, v := range m.Values {
		sum += v
	}
	return sum
}
