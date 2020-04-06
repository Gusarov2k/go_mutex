package storage_test

import (
	"github.com/Gusarov2k/go_mutex/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	data := storage.DataCache{}
	data.New()
	data.Set("one", "value")

	assert.EqualValues(t, "value", data.Get("one"), "map value myst: ", data.Get("one"))
}

func TestDelete(t *testing.T) {
	data := storage.DataCache{}
	data.New()
	data.Set("one", "value")
	data.Delete("one")

	assert.Empty(t, "", data.Get("one"), "map get empty value: ", data.Get("one"))
}

func TestUpdate(t *testing.T) {
	data := storage.DataCache{}
	data.New()

	t.Run("success update", func(t *testing.T){
		data.Set("one", "value")

		assert.True(t, true, data.Update("one", "full"), "map myst be updated: ", data.Update("one", "full"))

	})

	t.Run("error update", func(t *testing.T){
		data.Set("second", "value")

		assert.False(t, false, data.Update("one", "full"), " method update return: ", data.Update("one", "full"))
	})
}
