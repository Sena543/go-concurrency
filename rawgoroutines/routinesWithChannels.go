package rawgoroutines

import "time"

func CountWithChannel(arg string, c chan string) {

	for i := 0; i < 5; i++ {
		c <- arg
		time.Sleep(time.Millisecond * 500)
	}
}
