package multi

import (
	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Ordered](m *Multi[T]) T {
	var sum T
	for _, v := range m.Values {
		sum += v
	}
	return sum
}
