package cloudability

import (
	"fmt"
	"log"
	"time"
)

// Need to document this function and add tests
func retry(attempts int, sleep time.Duration, f func() (error, bool)) (err error) {
	var exit bool
	for i := 0; ; i++ {
		err, exit = f()
		if exit {
			return err
		}
		if i >= (attempts - 1) {
			break
		}
		time.Sleep(sleep)
		log.Println("retrying after error:", err.Error())
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err.Error())
}
