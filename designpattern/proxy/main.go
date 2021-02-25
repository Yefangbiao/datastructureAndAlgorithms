package main

func main() {
	mysql := &mysql{}
	proxy := NewMysqlServer(mysql, 3)
	proxy.connect()
	proxy.connect()
	proxy.connect()
	proxy.connect()
}
