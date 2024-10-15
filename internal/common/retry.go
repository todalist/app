package common

import (
	"errors"
	"time"
)


func WaitFor(intervalMillis int, times int, completeFunc func () bool) error {
	for {
		if completeFunc() {
			return nil
		}
		if times <= 0 {
			return errors.New("wait for timed out")
		}
		times--
		time.Sleep(time.Millisecond * time.Duration(intervalMillis))
	}
}
