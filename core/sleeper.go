package core

import "time"

type Sleeper interface {
	Sleep(duration time.Duration)
}

type RealSleeper struct{}

func (rs *RealSleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}
