// демонстарция остановки по контексту
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancelCtx := context.WithCancel(context.Background())

	wg.Add(1)
	go func(wg *sync.WaitGroup, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(wg, ctx)

	time.Sleep(3 * time.Second)
	cancelCtx()
	wg.Wait()
}
