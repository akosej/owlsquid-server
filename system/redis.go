package system

import (
	"fmt"
	"strings"
)

func AllKeyRedis() {
	AllUser = nil
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = RDB.Scan(CTX, cursor, "*", 0).Result()
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			AllUser = append(AllUser, key)
		}
		if cursor == 0 { // no more keys
			break
		}
	}
}

// GetAllUserStoredRedis --Get all users stored in redis
func getAllUserStoredRedisBocked() []NavigationUsersRedis {
	AllKeyRedis()
	//--------------
	var users []NavigationUsersRedis
	for _, user := range AllUser {
		var userR NavigationUsersRedis
		if err := RDB.HGetAll(CTX, user).Scan(&userR); err != nil {
			fmt.Println(err)
		}
		if userR.Email != "" {
			if userR.Bloquear {
				if strings.Contains(userR.IpRemote, ":") == false {
					users = append(users,
						NavigationUsersRedis{
							Email:     userR.Email,
							Quota:     userR.Quota,
							Used:      userR.Used,
							Update:    userR.Update,
							Last_size: userR.Last_size,
							Last_url:  userR.Last_url,
							IpRemote:  userR.IpRemote,
							Activa:    userR.Activa,
							Ilimitada: userR.Ilimitada,
							Bloquear:  userR.Bloquear,
						})
				}

			}
		}

	}
	return users
}
