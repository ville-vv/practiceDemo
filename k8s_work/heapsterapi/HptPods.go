package heapsterapi

import (
	"errors"
	"fmt"
)

type HptPod struct {
	NsName     string `json:"nsName"`
	PodName    string `json:"podName"`
	MetricName string `json:"metriceName"`
}

func NewHptPod(nsName, podName, metricName string) (pod *HptPod, err error) {
	pod = &HptPod{nsName, podName, metricName}
	return
}

func (this *HptPod) GetList(*HptRequest) (resp *HptResponse, err error) {
	return
}
func (this *HptPod) GetMetricList(*HptRequest) (resp *HptResponse, err error) {
	return
}
func (this *HptPod) GetMetricsInfo(req *HptRequest) (resp *HptResponse, err error) {
	if nil == req {
		err = errors.New("Request struct is nil")
		return
	}
	resp = &HptResponse{}
	// uri=/api/v1/model/namespaces/" + namespace + "/pods/" + podName + "/metrics/" + metricName
	resp.Url = fmt.Sprintf("%s/api/v1/model/namespaces/%s/pods/%s/metrics/%s%s", req.Url, this.NsName, this.PodName, this.MetricName, req.Param)
	resp.Body, err = Get(resp.Url)
	if err != nil {
		return
	}
	resp.NsName = this.NsName
	resp.PodName = this.PodName
	resp.MetricName = this.MetricName
	resp.Params = req.Param
	return
}
