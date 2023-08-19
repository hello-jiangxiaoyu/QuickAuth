package test

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestSingle(t *testing.T) {
	var wg sync.WaitGroup
	client := &http.Client{Timeout: 5 * time.Second}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			if _, err := client.Get("http://localhost/api/quick/apps"); err != nil {
				fmt.Println("err: ", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
