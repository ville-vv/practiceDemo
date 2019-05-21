package heapsterapi

import (
	"time"
	//"strings"
	//"errors"
)

type HeapsterMoniterData struct {
	GameID  uint32
	NsName  string
	PodName string
	//CpuUsage int64
	//MemoryUsage int64
	//FileSystemUsage int64
	//NetworkTx int64
	//NetworkRx int64
	ValueMap  map[string]int64
	TimeStamp time.Time
}

/**
 * 请求heapster API  数据的接口
 */
type IHeapsterAPI interface {
	/**
	 * 获取模块列表
	 */
	GetList(*HptRequest) (*HptResponse, error)
	/**
	 * 获取Metric 列表
	 */
	GetMetricList(*HptRequest) (*HptResponse, error)
	/**
	 * 获取Metric详细信息
	 */
	GetMetricsInfo(*HptRequest) (*HptResponse, error)
}

/**
 * 调用Heapster API 的请求参数
 */
type HptRequest struct {
	Url   string
	Param string
}

/**
 * 调用Heapster API 的返回参数
 */
type HptResponse struct {
	GameID     uint32 `json:"gameID"`
	Url        string `json:"url"`
	Params     string `json:"params"`
	Body       []byte `json:"body"`
	NsName     string `json:"nsName"`
	PodName    string `json:"podName"`
	MetricName string `json:"metricName"`
}

func GetTimeParams(duration time.Duration) string {
	now := time.Now().UTC()
	end := now.Format(time.RFC3339)
	start := now.Add((-duration) * 1e9).Format(time.RFC3339)

	var query string

	if len(start) != 0 {
		query = "?start=" + start
	} else {
		query += "?start=" + "1970-00-00T00:00:00Z"
	}

	if len(end) != 0 {
		query += "&end=" + end
	}
	return query
}
