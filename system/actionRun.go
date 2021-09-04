package system

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
)

func ActionRun(userLog string, bytes float64, dateJoin string, urlLog string, ipRemote string) (bool, string) {
	var user NavigationUsersRedis
	if err := RDB.HGetAll(CTX, userLog).Scan(&user); err != nil {
		fmt.Println(err)
	}
	if user.Bloquear == true {
		_, _ = RunString("tcpkill -i " + OwlInterface + " -9 host " + ipRemote + " and port " + OwlPortSquid + " > /dev/null 2>&1 &")
		fmt.Println("Kill request: " + urlLog + "  User: " + userLog + " from ip:" + ipRemote)
	} else {
		if user.Email == "" {
			bytesQuotaDefault, _ := strconv.ParseFloat(OwlQuotaDefault, 10)
			if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
				rdb.HSet(CTX, userLog, "email", userLog)
				rdb.HSet(CTX, userLog, "quota", bytesQuotaDefault)
				rdb.HSet(CTX, userLog, "used", bytes)
				rdb.HSet(CTX, userLog, "update", dateJoin)
				rdb.HSet(CTX, userLog, "last_url", urlLog)
				rdb.HSet(CTX, userLog, "last_size", bytes)
				rdb.HSet(CTX, userLog, "activa", 1)
				rdb.HSet(CTX, userLog, "ilimitada", 0)
				rdb.HSet(CTX, userLog, "bloquear", 0)
				return nil
			}); err != nil {
				fmt.Println(err)
			}
		} else {
			quota := user.Quota
			used := user.Used
			accumulated := used + bytes
			if quota > used && quota >= accumulated {
				if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
					rdb.HSet(CTX, userLog, "used", bytes)
					rdb.HSet(CTX, userLog, "update", dateJoin)
					rdb.HSet(CTX, userLog, "last_url", urlLog)
					rdb.HSet(CTX, userLog, "last_size", bytes)
					rdb.HSet(CTX, userLog, "activa", 1)
					return nil
				}); err != nil {
					fmt.Println(err)
				}
				ip := strings.Split(ipRemote, ":")
				fmt.Println("Update Quota used request:" + urlLog + " User:" + userLog + " url:" + ip[0])
				return true, "Update Quota used " + userLog
			} else {
				accumulated := used + bytes
				if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
					rdb.HSet(CTX, userLog, "used", accumulated)
					rdb.HSet(CTX, userLog, "activa", 0)
					rdb.HSet(CTX, userLog, "bloquear", 1)
					return nil
				}); err != nil {
					fmt.Println(err)
				}
				fmt.Println("User " + userLog + " lock -> send kill request")
				return true, "Kill request " + userLog
			}
		}
	}
	return false, "No new entries"
}
