package multitimer

import (
	"log"
	"time"
)

type Timer struct {
	d         time.Duration
	maxCount  int
	ID        int
	TimeOutFn func(int)
	pause     chan bool
}

func NewTimer(id int) Timer {
	result := Timer{ID: id}
	return result
}

func (t *Timer) Start() {

	if t.maxCount == 0 {
		t.maxCount = 1
	}
	t.pause = make(chan bool)
	if t.d > 0 {
		for i := 0; i < t.maxCount; i++ {

			select {
			case _ = <-t.pause:
				log.Println("[%d] : I was asked to PAUSE !!", t.ID)
				return
			default:
				time.Sleep(t.d)

				if t.TimeOutFn != nil {
					log.Printf("[%d] : %d of %d Calling timeoutFn()", t.ID, i, t.maxCount)
					t.TimeOutFn(t.ID)
				} else {
					log.Printf("[%d] : %d of %d Calling DONT KNOW WHOM ", t.ID, i, t.maxCount)

				}
			}

		}
	} else {
		log.Println("Cannot Start timer  ", t.d)
	}
}

func (t *Timer) Stop() {
	t.pause <- true
}

func (t *Timer) SetInterval(d time.Duration) {
	t.d = d
}

func (t *Timer) SetMaxCount(max int) {
	t.maxCount = max
}
