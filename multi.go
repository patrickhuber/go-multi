package multi

import "fmt"

type Multi[T any] struct {
	Dimensions []int
	Values     []T
}

func New[T any](dimensions ...int) *Multi[T] {
	if len(dimensions) == 0 {
		return &Multi[T]{}
	}
	product := 1
	for _, dim := range dimensions {
		product *= dim
	}
	return &Multi[T]{
		Dimensions: dimensions,
		Values:     make([]T, product),
	}
}

func (m *Multi[T]) Get(indicies ...int) (T, error) {
	var zero T
	index, err := m.index(indicies...)
	if err != nil {
		return zero, err
	}
	return m.Values[index], nil
}

func (m *Multi[T]) Set(value T, indicies ...int) error {
	index, err := m.index(indicies...)
	if err != nil {
		return err
	}
	m.Values[index] = value
	return nil
}

func (m *Multi[T]) index(indicies ...int) (int, error) {
	if len(m.Dimensions) == 0 {
		return 0, fmt.Errorf("empty array")
	}
	if len(indicies) != len(m.Dimensions) {
		return 0, fmt.Errorf("given %d dimensions but given %d dimensions", len(indicies), len(m.Dimensions))
	}

	result := 0
	for i, index := range indicies {
		if index < 0 {
			return 0, fmt.Errorf("invalid index %d must be >= 0", index)
		}
		if index >= m.Dimensions[i] {
			return 0, fmt.Errorf("invalid index %d must be < %d", index, m.Dimensions[i])
		}
		result += i*m.Dimensions[i] + index
	}
	if result >= len(m.Values) {
		return 0, fmt.Errorf("index %d exceeds max %d", result, len(m.Values))
	}
	return result, nil
}
