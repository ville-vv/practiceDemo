//
// @Time     : 19-9-16 下午4:37
// @Author   : Ville
// @File     : producer
// @Desc     :
// nsq_work
package svce

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"unsafe"
)

type Producer struct {
	addr string
	prd  *nsq.Producer
}

func NewProducer(addr string) (*Producer, error) {
	prd, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return nil, err
	}

	p := &Producer{
		addr: addr,
		prd:  prd,
	}
	return p, nil
}

func (p *Producer) publish(topic string, msg []byte) error {
	return p.prd.Publish(topic, msg)
}

func (p *Producer) PublishToJson(tp string, val interface{}) error {
	dt, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return p.publish(tp, dt)
}

func (p *Producer) PublishString(tp string, val string) error {
	x := (*[2]uintptr)(unsafe.Pointer(&val))
	h := [3]uintptr{x[0], x[1], x[1]}
	str := *(*[]byte)(unsafe.Pointer(&h))
	return p.publish(tp, str)
}
