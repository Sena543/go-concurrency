package rawgoroutines

import "sync"

func Wait(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		Count("someArg")
		wg.Done()
	}()

}
