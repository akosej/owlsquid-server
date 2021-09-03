package system

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
