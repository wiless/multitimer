package multitimer

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type Timer struct {
	d         time.Duration
	maxCount  int
	ID        int
	TimeOutFn func(int)
	pause     chan bool
	paussed   bool
	sync.Mutex
}

func NewTimer(id int) Timer {
	result := Timer{ID: id}
	result.pause = make(chan bool)
	result.paussed = false
	return result
}

func (t *Timer) Start() {

	if t.maxCount == 0 {
		t.maxCount = 1
	}

	if t.d > 0 {
		t.Mutex.Lock()
		t.paussed = false
		t.Mutex.Unlock()
		for i := 0; i < t.maxCount && t.paussed == false; i++ {

			// select {
			// case _ = <-t.pause:
			// 	log.Println("[%d] : I was asked to PAUSE !!", t.ID)
			// 	i = t.maxCount
			// 	break
			// default:
			time.Sleep(t.d)
			if t.paussed {
				log.Printf("[T %d] I was asked to resign while sleeping...", t.ID)
			} else {
				if t.TimeOutFn != nil {
					// r := reflect.ValueOf(t.TimeOutFn)
					// log.Printf("[Timer -%d] (%d/%d) %s", t.ID, i, t.maxCount, runtime.FuncForPC(r.Pointer()).Name())
					t.TimeOutFn(t.ID)
				} else {
					log.Printf("[T %d] : %d of %d Calling DONT KNOW WHOM ", t.ID, i, t.maxCount)

				}
			}
			// }

		}
		log.Println("Leaving timer ", t.ID)
		return
	} else {
		log.Println("Cannot Start timer  ", t.d)
	}

}

func (t *Timer) Stop() {
	t.Mutex.Lock()
	t.paussed = true
	t.Mutex.Unlock()
	runtime.Gosched()
	// t.pause <- true
}

func (t *Timer) SetInterval(d time.Duration) {
	t.d = d
}

func (t *Timer) SetMaxCount(max int) {
	t.maxCount = max
}
