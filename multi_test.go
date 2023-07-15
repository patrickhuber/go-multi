package multi_test

import (
	"testing"

	"github.com/patrickhuber/go-multi"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		arr := multi.New[int](10)
		require.Equal(t, 1, len(arr.Dimensions))

		err := arr.Set(5, 0)
		require.Nil(t, err)

		v, err := arr.Get(0)
		require.Nil(t, err)
		require.Equal(t, v, 5)

		err = arr.Set(5, 9)
		require.Nil(t, err)

		v, err = arr.Get(9)
		require.Nil(t, err)
		require.Equal(t, 5, v)

		// wrong index dimensions
		_, err = arr.Get(0, 0)
		require.NotNil(t, err)

		// index exceeds min
		_, err = arr.Get(-1)
		require.NotNil(t, err)

		// index exceeds max
		_, err = arr.Get(10)
		require.NotNil(t, err)
	})

	t.Run("two", func(t *testing.T) {
		arr := multi.New[int](10, 20)
		if len(arr.Dimensions) != 2 {
			t.Fatalf("expected 2 found %d", len(arr.Dimensions))
		}
		err := arr.Set(5, 0, 0)
		require.Nil(t, err)

		v, err := arr.Get(0, 0)
		require.Nil(t, err)
		require.Equal(t, 5, v)

		err = arr.Set(5, 9, 19)
		require.Nil(t, err)

		v, err = arr.Get(9, 19)
		require.Nil(t, err)
		require.Equal(t, 5, v)
	})
}
