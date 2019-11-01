package work_pool

import (
	"strconv"
	"testing"
)

func TestRunPool(t *testing.T) {
	go ChanFor()
	RunPool()
}

type AD struct {
	Cx []string
}

func (sel *AD) GetCxIndexRange() []string {
	rs := make([]string, len(sel.Cx))
	for i, v := range sel.Cx {
		rs[i] = v
	}
	return rs
}

func (sel *AD) GetCxIndexNoRange() []string {
	l := len(sel.Cx)
	rs := make([]string, l)
	for i := 0; i < l; i++ {
		rs[i] = sel.Cx[i]
	}
	return rs
}

func (sel *AD) GetCxAppendRange() []string {
	rs := make([]string, 0, len(sel.Cx))
	for _, v := range sel.Cx {
		rs = append(rs, v)
	}
	return rs
}

func (sel *AD) GetCxAppendNoRange() []string {
	l := len(sel.Cx)
	rs := make([]string, 0, l)
	for i := 0; i < l; i++ {
		rs = append(rs, sel.Cx[i])
	}
	return rs
}

func BenchmarkGetCxAppendRange(b *testing.B) {
	b.StopTimer()
	l := int64(100000000)
	ad := &AD{}
	ad.Cx = make([]string, l)
	for i := int64(0); i < l; i++ {
		ad.Cx[i] = strconv.Itoa(int(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ad.GetCxAppendRange()
	}
}

func BenchmarkGetIndexRange(b *testing.B) {
	b.StopTimer()
	l := int64(100000000)
	ad := &AD{}
	ad.Cx = make([]string, l)
	for i := int64(0); i < l; i++ {
		ad.Cx[i] = strconv.Itoa(int(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ad.GetCxIndexRange()
	}
}

func BenchmarkGetCxIndexNoRange(b *testing.B) {
	b.StopTimer()
	l := int64(100000000)
	ad := &AD{}
	ad.Cx = make([]string, l)
	for i := int64(0); i < l; i++ {
		ad.Cx[i] = strconv.Itoa(int(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ad.GetCxIndexNoRange()
	}
}

func BenchmarkGetCxAppendNoRange(b *testing.B) {
	b.StopTimer()
	l := int64(100000000)
	ad := &AD{}
	ad.Cx = make([]string, l)
	for i := int64(0); i < l; i++ {
		ad.Cx[i] = strconv.Itoa(int(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ad.GetCxAppendNoRange()
	}
}
