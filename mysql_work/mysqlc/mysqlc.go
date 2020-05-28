package mysqlc

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vilsongwei/vilgo/vuid"
	"os"
	"time"
)

type MysqlClient struct {
	db *sql.DB
}

// 建立mysql 连接
func NewMysqlClient(user string, password string, server_addr string, dbname string) (mc *MysqlClient, err error) {
	mc = &MysqlClient{}
	dataSourceName := fmt.Sprintf("%v:%v@(%v)/%v", user, password, server_addr, dbname)
	mc.db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	// 开启链接池，SetMaxOpenConns 设置最大链接数， SetMaxIdleConns 用于设置闲置的连接数。
	mc.db.SetMaxOpenConns(1000)
	mc.db.SetMaxIdleConns(100)

	err = mc.db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return mc, nil
}

func (m *MysqlClient) Create(name string) (err error) {
	res, err := m.db.Exec(fmt.Sprintf("create database %s", name))
	fmt.Println(res.RowsAffected())
	return
}
func (m *MysqlClient) Drop(name string) (err error) {
	res, err := m.db.Exec(fmt.Sprintf("drop database %s", name))
	fmt.Println(res.RowsAffected())
	return
}

// 关闭 mysql 连接
func (m *MysqlClient) Close() {
	m.db.Close()
}
func (m *MysqlClient) SelectExample(userid uint32) (nickname string, avatar string, err error) {
	var sqlstr string
	sqlstr = fmt.Sprintf("SELECT nickname,avatar FROM user.user WHERE userID = %d", userid)
	rows, err := m.db.Query(sqlstr)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&nickname, &avatar); err != nil {
			return
		}
	}
	return
}
func (m *MysqlClient) InsertExample(id int, name, remark string) (err error) {
	var elementStr string
	elementStr = fmt.Sprintf("(%d,'%s','%s')", id, name, remark)
	//ON DUPLICATE KEY UPDATE name=VALUES(name) 判断是否存在该字段，存在就更新
	_, err = m.db.Exec("INSERT INTO user.user (id, name, remark) VALUES " + elementStr + " ON DUPLICATE KEY UPDATE id=VALUES(id);")
	if err != nil {
		return
	}
	return
}
func (m *MysqlClient) UpdateExample(id int, name string) (err error) {
	sql := fmt.Sprintf("UPDATE user.user SET name='%s' WHERE id=%d  ;", name, id)
	_, err = m.db.Exec(sql)
	if err != nil {
		return
	}
	return
}
func (m *MysqlClient) DeleteExample(id int) (err error) {
	sql := fmt.Sprintf("DELETE FROM user.user  WHERE id=%d  ;", id)
	_, err = m.db.Exec(sql)
	if err != nil {
		return
	}
	return
}

func (m *MysqlClient) InsertTestData() {
	file, err := os.OpenFile("mysql_test_data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("open file fail %s", err.Error())
		return
	}
	sqlstr := "INSERT INTO vil_user.account(user_no, phone, email)values"
	file.WriteString(sqlstr + "\n")
	for i := 0; i < 1000000; i++ {
		phone := fmt.Sprintf("190878%d", i)
		sqlCmd := fmt.Sprintf("(%d, '%s', '%s'),", vuid.GenUUid(), phone, phone+"@test.com")

		//fmt.Printf(sqlCmd)
		//_, err := mysql.Db().Exec(sqlCmd)
		//if err != nil {
		//	fmt.Printf("mysql exec fail %s", err.Error())
		//	return
		//}

		file.WriteString(sqlCmd + "\n")
	}
	file.WriteString(";")
}

func (m *MysqlClient) SelectForDulip() {
	sqlbase := "SELECT * FROM vil_user.account WHERE user_no = %s;"
	strList := []string{"73836420374208512", "73836419799588864", "73836419636011008", "73836419585679360", "73836419610845184", "73836424564318208", "73836424669175808", "73836424740478976", "73836424945999872",
		"73836425277349888", "73836425378013184", "73836425625477120",
		"73836425696780288", "73836425730334720", "73836425797443584", "73836425835192320", "73836425885523968", "73836426019741696", "73836426422394880", "73836426787299328", "73836427181563904"}
	start := time.Now().UnixNano()
	for _, v := range strList {
		sqlcmd := fmt.Sprintf(sqlbase, v)
		if _, err := m.db.Exec(sqlcmd); err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("SelectForDulip:查询实际时间%d", end-start)
}

func CompStr(slist []string, cstr string) (out string) {
	for _, v := range slist {
		if out == "" {
			out = v
		}
		out = fmt.Sprintf("%s%s%s", out, cstr, v)
	}
	return out
}
func CompOr(slist []string, cstr string) (out string) {
	for _, v := range slist {
		if out == "" {
			out = cstr + "=" + v
		}
		out = fmt.Sprintf("%s OR %s", out, cstr+"="+v)
	}
	return out
}

func (m *MysqlClient) SelectWithIn() {
	sqlbase := "SELECT * FROM vil_user.account WHERE user_no in(%s);"
	strList := []string{"73836420374208512", "73836419799588864", "73836419636011008", "73836419585679360", "73836419610845184", "73836424564318208", "73836424669175808", "73836424740478976", "73836424945999872",
		"73836425277349888", "73836425378013184", "73836425625477120",
		"73836425696780288", "73836425730334720", "73836425797443584", "73836425835192320", "73836425885523968", "73836426019741696", "73836426422394880", "73836426787299328", "73836427181563904"}

	sqlcmd := fmt.Sprintf(sqlbase, CompStr(strList, ","))
	fmt.Printf(sqlcmd)
	start := time.Now().UnixNano()
	if _, err := m.db.Exec(sqlcmd); err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	end := time.Now().UnixNano()
	fmt.Printf("SelectWithIn:查询实际时间%d", end-start)

}

func (m *MysqlClient) SelectWithOr() {
	sqlbase := "SELECT * FROM vil_user.account WHERE %s;"
	strList := []string{"73836420374208512", "73836419799588864", "73836419636011008", "73836419585679360", "73836419610845184", "73836424564318208", "73836424669175808", "73836424740478976", "73836424945999872",
		"73836425277349888", "73836425378013184", "73836425625477120",
		"73836425696780288", "73836425730334720", "73836425797443584", "73836425835192320", "73836425885523968", "73836426019741696", "73836426422394880", "73836426787299328", "73836427181563904"}

	sqlcmd := fmt.Sprintf(sqlbase, CompOr(strList, "user_no"))
	fmt.Printf(sqlcmd)
	start := time.Now().UnixNano()
	if _, err := m.db.Exec(sqlcmd); err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	end := time.Now().UnixNano()
	fmt.Printf("SelectWithOr:查询实际时间%d", end-start)

}
