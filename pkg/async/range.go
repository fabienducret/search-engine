package async

import "sync"

func Range[T interface{}](elements []T, f func(element T)) {
	wg := sync.WaitGroup{}

	for _, e := range elements {
		wg.Add(1)

		go func(e T) {
			defer wg.Done()
			f(e)
		}(e)
	}

	wg.Wait()
}
