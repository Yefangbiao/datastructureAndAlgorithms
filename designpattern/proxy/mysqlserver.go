package main

import "fmt"

type mysqlServer struct {
	*mysql
	max int //最大的允许访问数
	num int //当前数量
}

func NewMysqlServer(mysql *mysql, max int) *mysqlServer {
	return &mysqlServer{mysql: mysql, max: max, num: 0}
}

func (m *mysqlServer) connect() {
	m.num++
	if m.num > m.max {
		fmt.Println("connected is unsuccess!")
	} else {
		m.mysql.connect()
	}
}
