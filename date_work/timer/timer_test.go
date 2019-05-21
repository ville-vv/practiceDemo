package timer

import (
	"fmt"
	"testing"
	"time"
)

func Test_Timer(t *testing.T) {
	timer := time.NewTicker(time.Second * 1)
	fmt.Println(time.Now())
	for {
		select {
		case <-timer.C:
			fmt.Println(time.Now())
		}
	}
}
func Test_timeAfter(t *testing.T) {
	fmt.Println(time.Now())
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println(time.Now())
		}
	}
}
