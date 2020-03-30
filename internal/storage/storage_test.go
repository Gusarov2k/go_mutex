package storage_test

import (
	"github.com/Gusarov2k/go_mutex/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	date := storage.DataCache{}
	date.Set("one", "value")

	assert.EqualValues(t, "value", date.Get("one"), "map value myst: ", date.Get("one"))
}