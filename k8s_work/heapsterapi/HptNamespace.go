package heapsterapi

import (
	"errors"
	"fmt"
)

/**
 *
 */
type HptNamespace struct {
	NsName     string `json:"nsName"`
	MetricName string `json:"metricName"`
}

func NewHptNamespace(nsName, metricName string) (ns *HptNamespace, err error) {
	ns = &HptNamespace{nsName, metricName}
	return
}

func (this *HptNamespace) GetList(*HptRequest) (resp *HptResponse, err error) {
	return
}
func (this *HptNamespace) GetMetricList(*HptRequest) (resp *HptResponse, err error) {
	return
}
func (this *HptNamespace) GetMetricsInfo(req *HptRequest) (resp *HptResponse, err error) {

	if nil == req {
		err = errors.New("Request struct is nil")
		return
	}
	resp = &HptResponse{}
	resp.Url = fmt.Sprintf("%s/api/v1/model/namespaces/%s/metrics/%s%s", req.Url, this.NsName, this.MetricName, req.Param)
	resp.Body, err = Get(resp.Url)
	if err != nil {
		return
	}
	resp.NsName = this.NsName
	resp.MetricName = this.MetricName
	resp.Params = req.Param
	return
}
