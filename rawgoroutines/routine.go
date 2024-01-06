package rawgoroutines

import (
	"fmt"
	"time"
)

func RunCountWithRoutines() {
	go Count("count 1")
	Count("count 2") // both run concurrently
	//running count 1 and 2 with go produces no results cos they are not joined back inot the main thread
}

func RunCountWithoutRoutines() {
	Count("count 1")
	Count("count 2") // count 2 never runs because 1 never finishes
}

func Count(somearg string) {
	// for i := 0; ; i++ {
	for i := 0; i < 5; i++ {
		fmt.Println(i, somearg)
		time.Sleep(time.Millisecond * 500)
	}
}
