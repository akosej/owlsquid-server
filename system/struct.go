package system

import (
	"gorm.io/gorm"
	"time"
)

type NavigationUsers struct {
	gorm.Model
	Name      string  `json:"name" form:"name"`
	Email     string  `gorm:"unique_index:user_email_index" json:"email" form:"email" binding:"required"`
	Quota     float64 `json:"quota" form:"quota" binding:"required"`
	Used      float64 `json:"used" form:"used"`
	Update    string  `json:"update" form:"update"`
	Last_size float64 `json:"last_size" form:"ast_size"`
	Last_url  string  `json:"last_url" form:"last_url"`
	Activa    bool    `json:"activa" form:"activa"`
	Ilimitada bool    `json:"ilimitada" form:"ilimitada"`
	Bloquear  bool    `json:"bloquear" form:"bloquear"`
}
type NavigationUsersRedis struct {
	Email     string  `json:"email" redis:"email" binding:"required"`
	Quota     float64 `json:"quota" redis:"quota" binding:"required"`
	Used      float64 `json:"used" redis:"used"`
	Update    string  `json:"update" redis:"update"`
	Last_size float64 `json:"last_size" redis:"ast_size"`
	Last_url  string  `json:"last_url" redis:"last_url"`
	IpRemote  string  `json:"ipremote" redis:"ipremote"`
	Activa    bool    `json:"activa" redis:"activa"`
	Ilimitada bool    `json:"ilimitada" redis:"ilimitada"`
	Bloquear  bool    `json:"bloquear" redis:"bloquear"`
}

type JobTicker struct {
	t *time.Timer
}
