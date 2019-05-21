package main

import (
	"practiceDemo/sync_pool/sample"
	"sync"
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

func main() {
	// // 禁用GC，并保证在main函数执行结束前恢复GC
	// defer debug.SetGCPercent(debug.SetGCPercent(-1))
	// var count int32

	// newFunc := func() interface{} {
	// 	return atomic.AddInt32(&count, 1)
	// }

	// pool := sync.Pool{New: newFunc}

	// // New 字段值的作用
	// v1 := pool.Get()
	// fmt.Printf("v1: %v\n", v1)

	// // 临时对象池的存取
	// pool.Put(newFunc())
	// pool.Put(newFunc())
	// pool.Put(newFunc())
	// v2 := pool.Get()
	// fmt.Printf("v2: %v\n", v2)

	// // 垃圾回收对临时对象池的影响
	// debug.SetGCPercent(100)
	// runtime.GC()
	// v3 := pool.Get()
	// fmt.Printf("v3: %v\n", v3)
	// pool.New = nil
	// v4 := pool.Get()
	// fmt.Printf("v4: %v\n", v4)

	////////////////----------------------
	//pool := sync.Pool{New: func() interface{} {
	//	return "empty string"
	//}}
	//s := "Hello World"
	//pool.Put(s)
	//pool.Put("bbbb")
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//fmt.Println(pool.Get())
	//
	//type testA = struct {
	//	AA int
	//}
	//type testS = struct {
	//	ab int
	//	aa *testA
	//}
	//
	//var ab testS
	//ab.ab = 0
	//
	//fmt.Printf("struct :%v", ab.aa)

	sample.SampleForSyncPool()
}
