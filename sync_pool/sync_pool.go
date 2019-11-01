package sync_pool

import (
	"fmt"
	"sync"
	"time"
)

// 每帧数据中包含多个数据项
type frameDataItem struct {
	SrcUserID uint32
	CpProto   string
	Timestamp uint64
}

// 每帧数据
type frameDataList struct {
	GameID     uint32
	RoomID     uint64
	FrameIndex uint32
	items      []*frameDataItem
}

// 每个房间中的帧数据
type roomFrameData struct {
	FrameData map[uint64]*frameDataList
}

// 房间中
type roomFrameDataPool struct {
	pool *sync.Pool
}

func NewRoomFrameDataPool() *roomFrameDataPool {
	return &roomFrameDataPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &roomFrameData{}
			},
		},
	}
}

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

func SampleForSyncPool() {
	standard()
	usePool()
}

func main() {
	SampleForSyncPool()
}
