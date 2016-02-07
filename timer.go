package main

import (
	"log"
	"time"
)

type Timer struct {
	d         time.Duration
	maxCount  int
	ID        int
	TimeOutFn func(int)
}

func NewTimer(id int) Timer {
	result := Timer{ID: id}
	return result
}

func (t *Timer) Start() {
	if t.maxCount == 0 {
		t.maxCount = 1
	}
	if t.d > 0 {
		for i := 0; i < t.maxCount; i++ {

			time.Sleep(t.d)

			if t.TimeOutFn != nil {
				log.Printf("I am out %d...., Calling timeoutFn()", t.ID)
				t.TimeOutFn(t.ID)
			} else {
				log.Printf("I am out %d...., but dont know what to do !! ", t.ID)

			}

		}
	}
}

func (t *Timer) Stop() {

}

func (t *Timer) SetInterval(d time.Duration) {
	t.d = d
}

func (t *Timer) SetMaxCount(max int) {
	t.maxCount = max
}
