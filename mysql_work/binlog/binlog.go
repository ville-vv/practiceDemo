package binlog

import (
	"context"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"os"
)

func Main() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 19,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "Root123",
	}
	syncer := replication.NewBinlogSyncer(cfg)
	syncer.GetNextPosition()

	stremer, _ := syncer.StartSync(mysql.Position{})

	for {
		ev, _ := stremer.GetEvent(context.Background())
		// Dump event
		ev.Dump(os.Stdout)
	}
}
