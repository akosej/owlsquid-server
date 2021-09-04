package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type NavigationUsersRedis struct {
	Email     string  `json:"email" redis:"email" binding:"required"`
	Quota     float64 `json:"quota" redis:"quota" binding:"required"`
	Used      float64 `json:"used" redis:"used"`
	Update    string  `json:"update" redis:"update"`
	Last_size float64 `json:"last_size" redis:"ast_size"`
	Last_url  string  `json:"last_url" redis:"last_url"`
	Activa    bool    `json:"activa" redis:"activa"`
	Ilimitada bool    `json:"ilimitada" redis:"ilimitada"`
	Bloquear  bool    `json:"bloquear" redis:"bloquear"`
}

func main() {
	ConnectServerRedis()
	reader := bufio.NewReader(os.Stdin)
	req, _ := reader.ReadString('\n')
	userLogin := strings.Split(req, " ")
	var user NavigationUsersRedis
	if err := RDB.HGetAll(CTX, userLogin[0]).Scan(&user); err != nil {
		fmt.Println(err)
	}
	if user.Bloquear == false {
		fmt.Println("OK")
	} else {
		fmt.Println("ERR")
	}
}
