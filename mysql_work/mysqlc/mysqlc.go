package mysqlc

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

func(m *MysqlClient)Create(name string) (err error){
	res, err := m.db.Exec(fmt.Sprintf("create database %s", name))
	fmt.Println(res.RowsAffected())
	return
}
func(m *MysqlClient)Drop(name string)(err error){
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
