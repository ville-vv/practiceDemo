package heapsterapi

import "time"

const (
	// supported CPU metrics
	CPU_LIMIT            string = "cpu/limit"
	CPU_USAGE            string = "cpu/usage"
	CPU_REQUEST          string = "cpu/reqeust"
	CPU_USAGE_RATE       string = "cpu/usage_rate"
	CPU_NODE_CAPACITY    string = "cpu/node_capacity"
	CPU_NODE_ALLOCATABLE string = "cpu/node_allocatable"
	CPU_NODE_RESERVATION string = "cpu/node_reservation"
	CPU_NODE_UTILIZATION string = "cpu/node_utilization"

	// supported Memory metrics
	MEMORY_LIMIT                  string = "memory/limit"
	MEMORY_USAGE                  string = "memory/usage"
	MEMORY_CACHE                  string = "memory/cache"
	MEMORY_RSS                    string = "memory/rss"
	MEMORY_MAJOR_PAGE_FAULTS      string = "memory/major_page_faults"
	MEMORY_MAJOR_PAGE_FAULTS_RATE string = "memory/major_page_faults_rate"
	MEMORY_NODE_CAPACITY          string = "memory/node_capacity"
	MEMORY_NODE_ALLOCATABLE       string = "memory/node_allocatable"
	MEMORY_NODE_RESERVATION       string = "memory/node_reservation"
	MEMORY_NODE_UTILIZATION       string = "memory/node_utilization"
	MEMORY_WORKING_SET            string = "memory/working_set"
	MEMORY_PAGE_FAULTS            string = "memory/page_faults"
	MEMORY_PAGE_FAULTS_RATE       string = "memory/page_faults_rate"

	// supported Network metrics
	NETWORK_RX             string = "network/rx"
	NETWORK_RX_RATE        string = "network/rx_rate"
	NETWORK_RX_ERRORS      string = "network/rx_errors"
	NETWORK_RX_ERRORS_RATE string = "network/rx_errors_rate"
	NETWORK_TX             string = "network/tx"
	NETWORK_TX_RATE        string = "network/tx_rate"
	NETWORK_TX_ERRORS      string = "network/tx_errors"
	NETWORK_TX_ERRORS_RATE string = "network/tx_errors_rate"

	// supported Filesystem metrics
	FILESYSTEM_LIMIT       string = "filesystem/limit"
	FILESYSTEM_USAGE       string = "filesystem/usage"
	FILESYSTEM_AVAILABLE   string = "filesystem/available"
	FILESYSTEM_INODES      string = "filesystem/inodes"
	FILESYSTEM_INODES_FREE string = "filesystem/inodes_free"

	// others
	UPTIME string = "uptime"
)

type CPUMetrics struct {
	CPU string `json:"CPU"`
	Metrics
}

type MemoryMetrics struct {
	Memory string `json:"Memory"`
	Metrics
}

type NetworkMetrics struct {
	Network string `json:"Network"`
	Metrics
}

type FileSystemMetrics struct {
	FileSystem string `json:"Filesystem"`
	Metrics
}

type UptimeMetrics struct {
	Uptime string `json:"Uptime"`
	Metrics
}

type Metrics struct {
	Metrics         []Metric  `json:"metrics"`
	LatestTimestamp time.Time `json:"latestTimestamp"`
}

type Metric struct {
	Timestamp time.Time `json:"timestamp"`
	Value     int64     `json:"value"`
}
