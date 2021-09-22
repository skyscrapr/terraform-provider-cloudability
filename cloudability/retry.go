package cloudability

import (
	"fmt"
	"log"
	"time"
)

// Need to document this function and add tests
func retry(attempts int, sleep time.Duration, f func() (bool, error)) (err error) {
	var exit bool
	for i := 0; ; i++ {
		exit, err = f()
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
