package system

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
	"time"
)

func ActionRun(userLog string, bytes float64, dateJoin string, urlLog string, ipRemote string) (bool, string) {
	var user NavigationUsersRedis
	IpRemote := ipRemote
	if err := RDB.HGetAll(CTX, userLog).Scan(&user); err != nil {
		fmt.Println(err)
	}
	if strings.Contains(ipRemote, ":") == true {
		IpRemote = user.IpRemote
	}
	if user.Bloquear == true {
		Run("tcpkill -i " + OwlInterface + " -9 host " + IpRemote + " and port " + OwlPortSquid + " > /dev/null 2>&1 &")
		Run("tcpkill -i " + OwlInterface + " -9 host " + IpRemote + " and port " + OwlPortSquidSSL + " > /dev/null 2>&1 &")
		if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
			rdb.HSet(CTX, user.Email, "activa", 0)
			return nil
		}); err != nil {
			fmt.Println(err)
		}
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
				rdb.HSet(CTX, userLog, "ipremote", IpRemote)
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
					rdb.HSet(CTX, userLog, "used", accumulated)
					rdb.HSet(CTX, userLog, "update", dateJoin)
					rdb.HSet(CTX, userLog, "last_url", urlLog)
					rdb.HSet(CTX, userLog, "last_size", bytes)
					rdb.HSet(CTX, userLog, "ipremote", IpRemote)
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
					rdb.HSet(CTX, userLog, "ipremote", IpRemote)
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

func KillUsersBlocked() {
	users := getAllUserStoredRedisBocked()
	fmt.Println(len(users))
	for _, user := range users {
		Run("tcpkill -i " + OwlInterface + " -9 host " + user.IpRemote + " and port " + OwlPortSquid + " > /dev/null 2>&1 &")
		Run("tcpkill -i " + OwlInterface + " -9 host " + user.IpRemote + " and port " + OwlPortSquidSSL + " > /dev/null 2>&1 &")
		if err := RDB.HGetAll(CTX, user.Email).Scan(&user); err != nil {
			fmt.Println(err)
		}
		if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
			rdb.HSet(CTX, user.Email, "activa", 0)
			return nil
		}); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Kill " + user.Email + " by ipaddress:" + user.IpRemote)
	}
	time.Sleep(10 * time.Second)
}
