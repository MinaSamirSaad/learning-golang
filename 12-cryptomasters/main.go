package main

import (
	"sync"

	"test.com/go/cryptomasters/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	// wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	for _, currency := range currencies {
		// increment the wait group counter
		wg.Add(1)
		// start a goroutine to get the currency data
		go func(currency string) {
			// decrement the wait group counter when the goroutine finishes
			defer wg.Done()
			api.GetCurrencyData(currency)
		}(currency)
	}
	// wait for all goroutines to finish
	wg.Wait()
}
