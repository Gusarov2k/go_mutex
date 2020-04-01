package storage_test

import (
	"github.com/Gusarov2k/go_mutex/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetGetAndDelete(t *testing.T) {
	data := storage.DataCache{}
	data.New()
	data.Set("one", "value")
	data.Set("second", "value_2")
	assert.EqualValues(t, "value", data.Get("one"), "map value myst: ", data.Get("one"))
	assert.EqualValues(t, "value_2", data.Get("second"), "map value myst: ", data.Get("second"))
	assert.EqualValues(t, "key deleted", data.Delete("one"), "map deleted: ", data.Delete("one"))
	assert.Empty(t, "", data.Get("one"), "map get empty value: ", data.Get("one"))
}

func TestSetAndUpdate(t *testing.T) {
	data := storage.DataCache{}
	data.New()
	data.Set("3", "empty")
	assert.EqualValues(t, "empty", data.Get("3"), "map value myst: ", data.Get("3"))
	assert.EqualValues(t, "key update", data.Update("3", "full"), "map value myst updated: ", data.Update("3", "full"))
	assert.EqualValues(t, "full", data.Get("3"), "map value myst: ", data.Get("3"))
}
