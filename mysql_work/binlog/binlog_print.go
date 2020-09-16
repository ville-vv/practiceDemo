package binlog

import (
	"context"
	"fmt"
	"github.com/siddontang/go-mysql/replication"
	"os"
)

func PrintSample() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 19,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "repl",
		Password: "repl123",
	}
	syncer := replication.NewBinlogSyncer(cfg)
	point := syncer.GetNextPosition()

	stremer, err := syncer.StartSync(point)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		ev, _ := stremer.GetEvent(context.Background())
		// Dump event
		ev.Dump(os.Stdout)
	}
}
