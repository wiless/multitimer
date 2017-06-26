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
	DoneFn    func(int)
	pause     chan bool
	paused    bool
	autoStart bool
	sync.Mutex
}

func NewTimer(id int) Timer {
	result := Timer{ID: id}
	result.pause = make(chan bool)
	result.paused = false
	result.autoStart = false
	return result
}

func (t *Timer) SetAutoStart(yes bool) {
	t.autoStart = yes
}

func (t *Timer) Start() {

	if t.maxCount == 0 {
		t.maxCount = 1
	} else {
		if t.autoStart {
			t.maxCount-- // Reduce the count as we immediately take first measurement @aba @ssk
		}
	}

	if t.d > 0 {
		t.Mutex.Lock()
		t.paused = false
		t.Mutex.Unlock()
		for i := 0; i < t.maxCount && t.paused == false; i++ {

			if !t.autoStart {
				//  sleep immediately..if not autostart function
				time.Sleep(t.d)
			}

			if t.paused {
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
			if t.autoStart {
				// if autostart.. sleep after executing first time
				time.Sleep(t.d)
			}

		}
		if t.DoneFn != nil {
			t.DoneFn(t.ID)
		}
		log.Println("Leaving timer ", t.ID)
		return
	} else {
		log.Println("Cannot Start timer  ", t.d)
	}

}

func (t *Timer) Stop() {
	t.Mutex.Lock()
	t.paused = true
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
