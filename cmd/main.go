package main

import (
	"log"
	"time"

	"github.com/wiless/multitimer"
)

func main() {

	t := make([]multitimer.Timer, 5)
	intervals := []int{2, 3, 7, 1, 10}
	for i := 0; i < 5; i++ {
		t[i] = NewTimer(i)
		t[i].SetMaxCount(5)
		t[i].SetInterval(time.Duration(intervals[i]) * time.Second)
		t[i].TimeOutFn = MeasureCurrent

	}
	for i := 0; i < 4; i++ {
		go t[i].Start()
	}
	t[4].Start()

}

func MeasureCurrent(indx int) {
	log.Printf("Some timer %d has passed out, I have to switch Relay %d and measure something ", indx, indx)
}
