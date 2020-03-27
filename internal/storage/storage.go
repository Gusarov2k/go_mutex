package storage

import (
	"fmt"
	"sync"
	)

type Storage struct {
	Data int
}

func (c *Storage) Set(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()    // блокируем доступ к переменной counter
	c.Data = 0
	for k := 1; k <= 5; k++{
		c.Data++
		fmt.Println("Goroutine", number, "-", c.Data)
	}
	mutex.Unlock()  // деблокируем доступ
	ch <- true
}
