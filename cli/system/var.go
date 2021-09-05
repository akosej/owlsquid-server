package system

import (
	"os"
	"time"
)

const IntervalPeriod = 24 * time.Hour

var (
	RedisServer     = Config("ip.db") + ":" + Config("port.db")
	RedisServerPass = Config("pass.db")
	EntityDomain    = Config("entity.domain")
	HOME            = os.Getenv("HOME")
	FolderConfig    = os.Getenv("HOME") + "/.owlcli"
	FileConfig      = os.Getenv("HOME") + "/.owlcli/config.owl"
	AllUser         []string
)
