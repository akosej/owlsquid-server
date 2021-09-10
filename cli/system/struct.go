package system

type Model struct {
	Email    string  `json:"email" redis:"email" binding:"required"`
	Quota    float64 `json:"quota" redis:"quota" binding:"required"`
	Used     float64 `json:"used" redis:"used"`
	Update   string  `json:"update" redis:"update"`
	IpRemote string  `json:"ipremote" redis:"ipremote"`
	Bloquear bool    `json:"bloquear" redis:"bloquear"`
}
