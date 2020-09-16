package binlog

import (
	"fmt"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

type CanalEventHandler struct {
}

func (c *CanalEventHandler) OnRotate(roateEvent *replication.RotateEvent) error {
	fmt.Println("OnRotate\t", roateEvent)
	return nil
}

func (c *CanalEventHandler) OnTableChanged(schema string, table string) error {
	fmt.Println("OnTableChanged\t", schema, "\t", table)
	return nil
}

func (c *CanalEventHandler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	fmt.Println("OnDDL\t", nextPos, "\t", queryEvent)
	return nil
}

func (c *CanalEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Println("OnRow\t", e)
	return nil
}

func (c *CanalEventHandler) OnXID(nextPos mysql.Position) error {
	fmt.Println("OnDDL\t", nextPos)
	return nil
}

func (c *CanalEventHandler) OnGTID(gtid mysql.GTIDSet) error {
	fmt.Println("OnGTID\t", gtid.String())
	return nil
}

func (c *CanalEventHandler) OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error {
	fmt.Println("OnPosSynced\t", pos, "\t", set, "\t", force)
	return nil
}

func (c *CanalEventHandler) String() string {
	return "CanalEventHandler"
}

func Canal() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "localhost:3306"
	cfg.User = "repl"
	cfg.Password = "repl123"
	cfg.ServerID = 234
	cfg.Dump.ExecutionPath = ""
	sqlCanal, err := canal.NewCanal(cfg)

	if err != nil {
		panic(err)
	}
	sqlCanal.SetEventHandler(&CanalEventHandler{})

	if err := sqlCanal.Run(); err != nil {
		panic(err)
		return
	}

	return
}
