package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]interface{}),
	}
}

func (m *ConcurrentMap) Set(key string, val interface{}) {
	m.Lock() // Lock для захвата на запись, во время этого захвата чтение недопустимо
	defer m.Unlock()
	m.data[key] = val
}

func (m *ConcurrentMap) Get(key string) (val interface{}, ok bool) {
	m.RLock() // RLock - допускаем параллельное чтение, не допускаем захват на запись
	defer m.RUnlock()
	val, ok = m.data[key]
	return
}

func (m *ConcurrentMap) Delete(key string) {
	m.Lock() // Изменение мапы - лок на запись и чтение
	defer m.Unlock()
	delete(m.data, key)
}

func (m *ConcurrentMap) Len() int {
	m.RLock() // не можем позволять запись (новых элементов), но можем делать чтение
	defer m.RUnlock()
	return len(m.data)
}

func main() {
	m := NewConcurrentMap()
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Set(fmt.Sprintf("%d", i), fmt.Sprintf("%d%d", i, i+1))
		}()
	}
	//wg.Wait() // раскомментировать, для того чтоб все элементы были записаны. На данный момент конкуретное чтение-запись
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := fmt.Sprintf("%d", i)
			if value, ok := m.Get(key); ok {
				fmt.Printf("ключ: %s, значение: %v\n", key, value)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Total elements:", m.Len())
}
