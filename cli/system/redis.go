package system

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
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
func GetAllUserStoredRedis(opt string) []Model {
	AllKeyRedis()
	//--------------
	var users []Model
	for _, user := range AllUser {
		var userR Model
		if err := RDB.HGetAll(CTX, user).Scan(&userR); err != nil {
			fmt.Println(err)
		}
		if userR.Email != "" {
			if opt == "actives" {
				if userR.Bloquear == false {
					users = append(users, Model{Email: userR.Email, Quota: userR.Quota, Used: userR.Used, Update: userR.Update, Bloquear: userR.Bloquear})
				}
			} else if opt == "blocked" {
				if userR.Bloquear == true {
					users = append(users, Model{Email: userR.Email, Quota: userR.Quota, Used: userR.Used, Update: userR.Update, Bloquear: userR.Bloquear})
				}
			} else {

				users = append(users, Model{Email: userR.Email, Quota: userR.Quota, Used: userR.Used, Update: userR.Update, Bloquear: userR.Bloquear})
			}
		}

	}
	return users
}

// GetUserStoredRedis --Get all users stored in redis
func GetUserStoredRedis(user string) []Model {
	AllKeyRedis()
	//--------------
	var users []Model
	var userR Model
	if err := RDB.HGetAll(CTX, user+"@"+EntityDomain).Scan(&userR); err != nil {
		fmt.Println(err)
	}
	users = append(users, Model{Email: userR.Email, Quota: userR.Quota, Used: userR.Used, Update: userR.Update, Bloquear: userR.Bloquear})
	return users
}

// ResetUserStoredRedis --Get all users stored in redis
func OptUserStoredRedis(user string, opt string, value string) bool {
	var userR Model
	if err := RDB.HGetAll(CTX, user+"@"+EntityDomain).Scan(&userR); err != nil {
		return false
	}
	if userR.Email != "" {
		if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
			if opt == "reset" {
				rdb.HSet(CTX, user+"@"+EntityDomain, "used", 0)
				rdb.HSet(CTX, user+"@"+EntityDomain, "activa", 1)
				rdb.HSet(CTX, user+"@"+EntityDomain, "bloquear", 0)
			} else if opt == "newquota" {
				quotaMB, _ := strconv.Atoi(value)
				rdb.HSet(CTX, user+"@"+EntityDomain, "quota", quotaMB*1024*1024)
			} else {
				rdb.HSet(CTX, user+"@"+EntityDomain, "activa", 0)
				rdb.HSet(CTX, user+"@"+EntityDomain, "bloquear", 1)
			}
			return nil
		}); err != nil {
			return false
		}
		return true
	} else {
		return false
	}
}

// OptAllUserStoredRedis --Get all users stored in redis
func OptAllUserStoredRedis(opt string, value string) bool {
	AllKeyRedis()
	//--------------
	for _, user := range AllUser {
		var userR Model
		if err := RDB.HGetAll(CTX, user).Scan(&userR); err != nil {
			fmt.Println(err)
		}
		if userR.Email != "" {
			if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
				if opt == "reset" {
					rdb.HSet(CTX, userR.Email, "used", 0)
					rdb.HSet(CTX, userR.Email, "activa", 1)
					rdb.HSet(CTX, userR.Email, "bloquear", 0)
				} else if opt == "newquota" {
					quotaMB, _ := strconv.Atoi(value)
					rdb.HSet(CTX, userR.Email, "quota", quotaMB*1024*1024)
				} else {
					rdb.HSet(CTX, userR.Email, "activa", 0)
					rdb.HSet(CTX, userR.Email, "bloquear", 1)
				}
				return nil
			}); err != nil {
				return false
			}
		}
	}
	return true
}
