package memorystorage

import (
	"testing"

	"github.com/Reversaidx/hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	// TODO
	t.Run("Complex", func(t *testing.T) {
		test := New()
		event := storage.Event{
			Title: "test",
		}
		test.Add(event)
		test.Change(0, storage.Event{
			Title: "test2",
		})
		require.Equal(t, 1, len(test.List()))
		require.Equal(t, test.List()[0].Title, "test2")
		test.Delete(0)
		require.Equal(t, 0, len(test.List()))
		err := test.Change(0, storage.Event{
			Title: "test3",
		})
		require.Error(t, err)
	})
}
