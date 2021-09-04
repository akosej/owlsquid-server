package system

import "fmt"

func UnBlockLast() {
	AllKeyRedis()
	for _, user := range AllUser {
		UsersBlockLast = append(UsersBlockLast, user)
	}

}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func UnBlock() {
	equal := Equal(UsersBlock, UsersBlockLast)
	if equal == false {
		AllKeyRedis()
		UsersBlock = nil
		for _, user := range AllUser {
			var modelUserRedis NavigationUsersRedis
			// Scan all fields into the model.
			if err := RDB.HGetAll(CTX, user).Scan(&modelUserRedis); err != nil {
				fmt.Println(err)
			}
			if modelUserRedis.Bloquear == true {
				UsersBlock = append(UsersBlock, user)
			}
		}
		UsersBlockLast = nil
		ResetFile(OwlFolderAcls + "/owl_acl_user_denied")
		//_, _ = RunString("chmod -R 777 " + OwlFolderAcls + "/owl_acl_user_denied")
		for _, email := range UsersBlock {
			_ = AppendStrFile(OwlFolderAcls+"/owl_acl_user_denied", email+"\n")
			UsersBlockLast = append(UsersBlockLast, email)
		}
		//--Reload Squid
		_, _ = RunString("/etc/init.d/squid reload")
	} else {
		UsersBlock = nil
		AllKeyRedis()
		for _, user := range AllUser {
			var modelUserRedis NavigationUsersRedis
			// Scan all fields into the model.
			if err := RDB.HGetAll(CTX, user).Scan(&modelUserRedis); err != nil {
				fmt.Println(err)
			}
			if modelUserRedis.Bloquear == true {
				UsersBlock = append(UsersBlock, user)
			}
		}
	}
}
