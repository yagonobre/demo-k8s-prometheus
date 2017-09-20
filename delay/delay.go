package delay

import (
	"math/rand"
	"time"
)

func Delay1() {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func Delay2() {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func Delay3() {
	r := rand.Intn(500)
	time.Sleep(time.Duration(r) * time.Millisecond)
}
