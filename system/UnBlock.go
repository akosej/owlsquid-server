package system

import "fmt"

func UnBlockLast() {
	if OwlServerDBType == "redis" {
		AllKeyRedis()
		for _, user := range AllUser {
			UsersBlockLast = append(UsersBlockLast, user)
		}
	} else {
		var users []NavigationUsers
		DB.Where("Bloquear = ?", true).Find(&users)
		for _, user := range users {
			UsersBlockLast = append(UsersBlockLast, user.Email)
		}
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
	if OwlServerDBType == "redis" {
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

			for _, email := range UsersBlock {
				_ = AppendStrFile(OwlFolderAcls+"/owl_acl_user_denied", email+"\n")
				UsersBlockLast = append(UsersBlockLast, email)
			}
			//--Reload Squid
			_, _ = RunString(OwlRunOrder)
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
		//	-------------------------------------MYSQL
	} else {
		var users []NavigationUsers
		DB.Where("Bloquear = ?", true).Find(&users)
		if equal == false {
			UsersBlock = nil
			for _, user := range users {
				UsersBlock = append(UsersBlock, user.Email)
			}
			UsersBlockLast = nil
			ResetFile(OwlFolderAcls + "/owl_acl_user_denied")

			for _, email := range UsersBlock {
				_ = AppendStrFile(OwlFolderAcls+"/owl_acl_user_denied", email+"\n")
				UsersBlockLast = append(UsersBlockLast, email)
			}
			//--Reload Squid
			_, _ = RunString(OwlRunOrder)
		} else {
			UsersBlock = nil
			for _, user := range users {
				UsersBlock = append(UsersBlock, user.Email)
			}
		}
	}

}
