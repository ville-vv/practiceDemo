package sample


import (
	"sync"
	"time"
	"fmt"
)

type structR6 struct {
	B1 [100000]int
}
var r6Pool = sync.Pool{
	New: func() interface{} {
		return new(structR6)
	},
}
func usePool() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		sr6 := r6Pool.Get().(*structR6)
		sr6.B1[0] = 0
		r6Pool.Put(sr6)
	}
	fmt.Println("pool Used:", time.Since(startTime))
}
func standard() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		var sr6 structR6
		sr6.B1[0] = 0
	}
	fmt.Println("standard Used:", time.Since(startTime))
}
func main() {

}

func SampleForSyncPool() {
	standard()
	usePool()
}