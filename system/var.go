package system

import "time"

const IntervalPeriod = 24 * time.Hour

var (
	RedisServer     = Config("ip.db") + ":" + Config("port.db")
	RedisServerPass = Config("pass.db")
	OwlFolderAcls   = Config("folder.owl_acls")
	OwlFolderLogs   = Config("folder.salva_logs")
	OwlInterface    = Config("interface.server")
	OwlQuotaDefault = Config("default.quota")
	OwlAccesslog    = Config("path.AccessLog")
	OwlRestart      = Config("jobs.restart")
	OwlPortSquid    = Config("squid.port")

	// Connects -----JOBS
	Connects     []string
	JobCheckUser = make(chan string)
	ConnectType  = []string{"TCP_MISS/200", "TCP_TUNNEL/200"}
	Jobs         []JobTicker
	JobsEnd      []string
	AllUser      []string
)
