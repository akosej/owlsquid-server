package system

type Model struct {
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
