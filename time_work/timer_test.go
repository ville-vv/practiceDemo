package time_work

import (
	"fmt"
	"testing"
	"time"
)

/*
这里推荐使用定时器  NewTicker ，
定时 time.After 每次都会创建一个新的变量
*/
// 定时器使用方法 1
func Test_Timer(t *testing.T) {
	timer := time.NewTicker(time.Second * 1)
	fmt.Println(time.Now())
	for {
		select {
		case <-timer.C:
			fmt.Println(time.Now())
		}
	}
}

// 定时器使用方法 2
func Test_timeAfter(t *testing.T) {
	fmt.Println(time.Now())
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println(time.Now())
		}
	}
}

func TestTimeSubDates(t *testing.T) {
	tm := time.Now().UTC()
	t1, err := time.Parse("2006-01-02", "2017-08-07")
	fmt.Println("时间：", t1)
	if err != nil {
		fmt.Println("时间转换错误：", err)
	}
	fmt.Println("时间差：", TimeSubDates(tm, t1))
	fmt.Println("时间差：", TimeSub(tm, t1))
	fmt.Println("时间差：", TimeSub2(tm, t1))

	fmt.Println("时间差：", TimeSubDates(tm, tm.Add(time.Hour*-34)))
	fmt.Println("时间差：", TimeSub(tm, tm.Add(time.Hour*-24)))
	fmt.Println("时间差：", TimeSub2(tm, tm.Add(time.Hour*-24)))
}
