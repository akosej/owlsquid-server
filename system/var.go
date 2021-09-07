package system

import "time"

const IntervalPeriod = 24 * time.Hour

var (
	RedisServer       = Config("redis.ip") + ":" + Config("redis.port")
	RedisServerPass   = Config("redis.pass")
	OwlFolderLogs     = Config("folder.salve_logs")
	OwlInterface      = Config("interface.server")
	OwlQuotaDefault   = Config("default.quota")
	OwlAccesslog      = Config("path.AccessLog")
	OwlRestart        = Config("jobs.restart")
	OwlPortSquid      = Config("squid.port")
	OwlPortSquidSSL   = Config("squid.portssl")
	OwlRunScript      = Config("path.RunScript")
	OwlActivesRequest = Config("path.ActivesRequest")

	// Connects -----JOBS
	Connects     []string
	JobCheckUser = make(chan string)
	ConnectType  = []string{"TCP_MISS/200", "TCP_TUNNEL/200"}
	Jobs         []JobTicker
	JobsEnd      []string
	AllUser      []string
)
