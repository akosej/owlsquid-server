package system

import "time"

const IntervalPeriod = 24 * time.Hour

var (
	RedisServer     = Config("redis.ip") + ":" + Config("redis.port")
	RedisServerPass = Config("redis.pass")
	OwlFolderLogs   = Config("folder.salve_logs")
	OwlInterface    = Config("interface.server")
	OwlQuotaDefault = Config("default.quota")
	OwlAccesslog    = Config("path.AccessLog")
	OwlRestart      = Config("jobs.restart")
	OwlPortSquid    = Config("squid.port")
	OwlPortSquidSSL = Config("squid.portssl")

	// Connects -----JOBS
	Connects     []string
	JobCheckUser = make(chan string)
	ConnectType  = []string{"TCP_CLIENT_REFRESH_MISS/200", "TCP_MISS/200", "TCP_MISS_ABORTED/200", "TCP_TUNNEL/200", "TCP_TUNNEL_ABORTED/200", "TCP_REFRESH_UNMODIFIED/200", "TCP_NEGATIVE_HIT/200", "TCP_REFRESH_MISS/200", "TCP_SWAPFAIL_MISS/200"}
	Jobs         []JobTicker
	JobsEnd      []string
	AllUser      []string
)
