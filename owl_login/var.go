package main

var (
	RedisServer     = Config("redis.ip") + ":" + Config("redis.port")
	RedisServerPass = Config("redis.pass")
)
