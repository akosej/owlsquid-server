package system

import "time"

const IntervalPeriod = 24 * time.Hour

var (
	OwlServerDBType  = Config("server.DB.Type")
	RedisServer      = Config("ip.db") + ":" + Config("port.db")
	RedisServerPass  = Config("pass.db")
	MysqlCredentials = Config("user.db") + ":" + Config("pass.db") + "@tcp(" + Config("ip.db") + ":" + Config("port.db") + ")/" + Config("name.db") + "?charset=utf8&parseTime=True&loc=Local"
	OwlFolderAcls    = Config("folder.owl_acls")
	OwlFolderLogs    = Config("folder.salva_logs")
	OwlInterface     = Config("interface.server")
	OwlQuotaDefault  = Config("default.quota")
	OwlAccesslog     = Config("path.AccessLog")
	OwlRunOrder      = Config("restart.order")
	OwlRestart       = Config("jobs.restart")

	//-----JOBs
	Connects       []string
	JobCheckUser   = make(chan string)
	ConnecType     = []string{"TCP_MISS/200", "TCP_TUNNEL/200"}
	Jobs           []JobTicker
	JobsEnd        []string
	UsersBlockLast []string
	UsersBlock     []string
	AllUser        []string
)
