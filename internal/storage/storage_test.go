package storage_test

import (
	"github.com/Gusarov2k/go_mutex/internal/storage"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

func TestSet(t *testing.T) {
	store := storage.Storage{}
	log.Println("store in tests :", store.Data)
	ch := make(chan bool)       // канал
	var mutex sync.Mutex        // определяем мьютекс
	for i := 1; i < 5; i++{
		go store.Set(i, ch, &mutex)
	}
	log.Println("store after gorutines :", store.Data)
	assert.EqualValues(t, 5, store.Data, "store.Data - got: ", store.Data)
}