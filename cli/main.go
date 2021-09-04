package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	RedisServer     = Config("ip.db") + ":" + Config("port.db")
	RedisServerPass = Config("pass.db")
	AllUser         []string
	CTX             context.Context
	RDB             *redis.Client
)

type Model struct {
	Email    string  `json:"email" redis:"email" binding:"required"`
	Quota    float64 `json:"quota" redis:"quota" binding:"required"`
	Used     float64 `json:"used" redis:"used"`
	Update   string  `json:"update" redis:"update"`
	Bloquear bool    `json:"bloquear" redis:"bloquear"`
}

// ConnectServerRedis -- Connect to the redis server
func ConnectServerRedis() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisServer,
		Password: RedisServerPass,
	})
	CTX = ctx
	RDB = rdb
}

// AllKeyRedis -- Get all keys stored in redis
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
	return users
}

// PrintCommandList --Print command list
func PrintCommandList() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Command", "Description"})
	t.AppendRow([]interface{}{1, "owlcli list all", "List all registered users"})
	t.AppendRow([]interface{}{2, "owlcli list actives", "List active users"})
	t.AppendRow([]interface{}{3, "owlcli list blocked", "List blocked users"})
	t.AppendSeparator()
	//t.AppendRow([]interface{}{4, "owlcli search user", "Search for a user"})
	t.AppendFooter(table.Row{"#", "Command", "Description"})
	t.Render()
}

// PrintUserTable --Print user table
func PrintUserTable(users []Model) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Email", "Quota MB", "Used MB", "Update", "Lock"})
	for num, user := range users {
		t.AppendSeparator()
		date := strings.Split(SubtractDates(user.Update).String(), ".")
		t.AppendRow([]interface{}{num, user.Email, user.ConvertBitsMB("quota"), user.ConvertBitsMB("used"), date[0], user.Bloquear})
	}
	t.AppendFooter(table.Row{"#", "Email", "Quota MB", "Used MB", "Update", "Lock"})
	t.Render()
}

func main() {
	if len(os.Args) < 2 {
		color.Red("Command list")
		PrintCommandList()
		return
	}
	ConnectServerRedis()
	switch arg := os.Args[1]; arg {
	case "list":
		if len(os.Args) < 3 {
			color.Red("Command list")
			PrintCommandList()
			return
		}
		switch arg := os.Args[2]; arg {
		case "all":
			color.Cyan("List of users saved by owl")
			PrintUserTable(GetAllUserStoredRedis("all"))
		case "actives":
			color.Cyan("List actives users")
			PrintUserTable(GetAllUserStoredRedis("actives"))
		case "blocked":
			color.Cyan("List blocked users")
			PrintUserTable(GetAllUserStoredRedis("blocked"))
		default:
			color.Red("Command list")
			PrintCommandList()
		}

	default:
		color.Red("Command list")
		PrintCommandList()
	}
}

func Config(data string) string {
	lines, err := File2lines("./config.owl")
	value := ""
	if err != nil {
		fmt.Println("The configuration file could not be found")
	} else {
		// --- Extract the variables from the configuration file
		for _, line := range lines {
			if strings.Contains(line, data) {
				cut := strings.Split(line, "=")
				value += cut[1]
			}
		}
	}
	// -------------------------------------------
	return value
}

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func (u Model) ConvertBitsMB(str string) float64 {
	total := 0.0
	if str == "used" {
		total = u.Used / 1024 / 1024
	} else {
		total = u.Quota / 1024 / 1024
	}

	return Round(total, 2)
}

func Round(result float64, places int) float64 {
	var round float64

	if places == 1 {
		round = math.Round(result)
	} else if places == 2 {
		Rounding, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
		round = Rounding
	} else if places == 3 {
		Rounding, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", result), 64)
		round = Rounding
	} else {
		Rounding, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		round = Rounding
	}

	return round
}

func SubtractDates(date string) time.Duration {
	a, _ := strconv.ParseInt(date, 10, 64)
	ta := time.Unix(a, 0)
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
	parse, _ := dateparse.ParseLocal(ta.String())
	return time.Now().Sub(parse)
}
