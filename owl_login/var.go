package main

var (
	RedisServer     = Config("ip.db") + ":" + Config("port.db")
	RedisServerPass = Config("pass.db")
	AllUser         []string
)
