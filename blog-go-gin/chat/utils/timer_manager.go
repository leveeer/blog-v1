package utils

import "time"

type Event struct {
	id   int
	Exec func()
}

var timerIndex = 1
var timers map[int]*time.Timer
var eventChan chan *Event

func AddTimer(t time.Duration, f func()) int {
	idx := timerIndex
	timerIndex++
	timers[idx] = time.AfterFunc(t, func() {
		eventChan <- &Event{id: idx, Exec: f}
	})
	return idx
}

func init() {
	eventChan = make(chan *Event, 100)
	timers = make(map[int]*time.Timer)
}

func RemoveTimer(idx int) {
	if t, ok := timers[idx]; ok {
		delete(timers, idx)
		t.Stop()
	}
}

func TimerExe(e *Event) {
	if _, ok := timers[e.id]; ok {
		delete(timers, e.id)
		e.Exec()
	}
}
func TimerTick() chan *Event {
	return eventChan
}

func GetTimerNum() int {
	return len(timers)
}
