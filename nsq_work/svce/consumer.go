// @File     : consumer
// @Author   : Ville
// @Time     : 19-9-16 下午4:37
// @Desc     :
// nsq_work
package svce

import (
	"github.com/nsqio/go-nsq"
)

type Consumer struct {
	handler func([]byte) error
}

func NewConsumer(topic, channel, addr string, chFun func(msg []byte) error) (*Consumer, error) {
	cmer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	cm := &Consumer{handler: chFun}

	cmer.AddHandler(cm)
	err = cmer.ConnectToNSQD(addr)
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func (c *Consumer) HandleMessage(msg *nsq.Message) error {
	body := make([]byte, len(msg.Body))
	copy(body, msg.Body)
	return c.handler(body)
}
