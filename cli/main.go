package main

import (
	"cli/system"
	"cli/tools"
	"github.com/chzyer/readline"
	"github.com/pterm/pterm"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// completer defines which commands the user can use
var completer = readline.NewPrefixCompleter()

// categories holding the initial default categories. The user can  add categories.
var listArg = []string{"all", "actives", "blocked"}
var l *readline.Instance

func main() {
	system.VerifyConfiguration()
	tools.Clear()
	system.ConnectServerRedis()
	tools.IntroScreen()
	// Initialize config
	config := readline.Config{
		Prompt:            "\033[31mcommand »\033[0m ",
		AutoComplete:      completer,
		HistoryFile:       system.FileHistory,
		HistoryLimit:      20,
		HistorySearchFold: true,
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
	}

	var err error
	// Create instance
	l, err = readline.NewEx(&config)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// Initial initialization of the completer
	updateCompleter()

	log.SetOutput(l.Stderr())

	// This loop watches for user input and process it
	for {
		line, err := l.Readline()

		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		// Checking which command the user typed
		args := strings.Split(line, " ")
		switch arg := args[0]; arg {
		case "list":
			if len(args) < 2 {
				tools.PrintCommandList()
				continue
			}
			switch arg := args[1]; arg {
			case "all":
				tools.PrintUserTable(system.GetAllUserStoredRedis("all"), "List of users saved by owl", true)
			case "actives":
				tools.PrintUserTable(system.GetAllUserStoredRedis("actives"), "List actives users", true)
			case "blocked":
				tools.PrintUserTable(system.GetAllUserStoredRedis("blocked"), "List blocked users", true)
			default:
				tools.PrintCommandList()
			}
		case "status":
			if len(args) < 2 {
				tools.PrintCommandList()
				continue
			}
			switch arg := args[1]; arg {
			case "user":
				if len(args) < 3 {
					tools.PrintCommandList()
					continue
				}
				tools.PrintUserTable(system.GetUserStoredRedis(args[2]), "User information", false)
			default:
				tools.PrintCommandList()
			}
		case "reset":
			if len(args) < 2 {
				tools.PrintCommandList()
				continue
			}
			switch arg := args[1]; arg {
			case "all":
				result := system.OptAllUserStoredRedis("reset", "")
				if result {
					pterm.FgCyan.Println("Quota has been successfully reset")
				} else {
					pterm.FgRed.Println("An error has occurred")
				}
			case "user":
				if len(args) < 3 {
					tools.PrintCommandList()
					continue
				}
				result := system.OptUserStoredRedis(args[2], "reset", "")
				if result {
					pterm.FgCyan.Println("The user's browsing quota has been reset")
					tools.PrintUserTable(system.GetUserStoredRedis(args[2]), "User information", false)
				} else {
					pterm.FgRed.Println("An error has occurred, verify that the user is correctly typed, remember that you should not put @domain")
				}
			default:
				tools.PrintCommandList()
			}
		case "blocked":
			if len(args) < 2 {
				tools.PrintCommandList()
				continue
			}
			switch arg := args[1]; arg {
			case "all":
				result := system.OptAllUserStoredRedis("blocked", "")
				if result {
					pterm.FgCyan.Println("All users have been blocked")
				} else {
					pterm.FgRed.Println("An error has occurred")
				}
			case "user":
				if len(args) < 3 {
					tools.PrintCommandList()
					continue
				}
				result := system.OptUserStoredRedis(args[2], "blocked", "")
				if result {
					pterm.FgCyan.Println("User has been successfully blocked ")
					tools.PrintUserTable(system.GetUserStoredRedis(args[2]), "User information", false)
				} else {
					pterm.FgRed.Println("An error has occurred, verify that the user is correctly typed, remember that you should not put @domain")
				}
			default:
				tools.PrintCommandList()
			}
		case "unblocked":
			if len(args) < 2 {
				tools.PrintCommandList()
				continue
			}
			switch arg := args[1]; arg {
			case "all":
				//todo: No funciona
				result := system.OptAllUserStoredRedis("unblocked", "")
				if result {
					pterm.FgCyan.Println("All users have been unblocked")
				} else {
					pterm.FgRed.Println("An error has occurred")
				}
			case "user":
				if len(args) < 3 {
					tools.PrintCommandList()
					continue
				}
				result := system.OptUserStoredRedis(args[2], "unblocked", "")
				if result {
					pterm.FgCyan.Println("User has been successfully unblocked ")
					tools.PrintUserTable(system.GetUserStoredRedis(args[2]), "User information", false)
				} else {
					pterm.FgRed.Println("An error has occurred, verify that the user is correctly typed, remember that you should not put @domain")
				}
			default:
				tools.PrintCommandList()
			}
		case "setquota":
			if len(args) < 2 {
				tools.PrintCommandList()
				continue
			}
			switch arg := args[1]; arg {
			case "all":
				result := system.OptAllUserStoredRedis("newquota", args[2])
				if result {
					pterm.FgCyan.Println("User quota has been changed successfully ")
				} else {
					pterm.FgRed.Println("An error has occurred")
				}
			case "user":
				if len(args) < 4 {
					tools.PrintCommandList()
					continue
				}
				result := system.OptUserStoredRedis(args[2], "newquota", args[3])
				if result {
					pterm.FgCyan.Println("User quota has been changed successfully ")
					tools.PrintUserTable(system.GetUserStoredRedis(args[2]), "User information", false)
				} else {
					pterm.FgRed.Println("An error has occurred, verify that the user is correctly typed, remember that you should not put @domain")
				}
			default:
				tools.PrintCommandList()
			}
		case "clear":
			tools.Clear()
			tools.IntroScreen()
		case "exit":
			pterm.FgCyan.Println("See you soon!!!")
			time.Sleep(2 * time.Second)
			tools.Clear()
			os.Exit(0)
		default:
			tools.PrintCommandList()
		}
	}

}

// updateCompleter is updates the completer allowing to add new command during runtime. The completer is recreated
// and the configuration of the instance update.
func updateCompleter() {

	var itemsList []readline.PrefixCompleterInterface
	var allUser []readline.PrefixCompleterInterface
	for _, all := range listArg {
		itemsList = append(itemsList, readline.PcItem(all))
	}
	users := system.GetAllUserStoredRedis("")
	for _, u := range users {
		if u.Email != "" {
			allUser = append(allUser, readline.PcItem(strings.Split(u.Email, "@")[0]))
		}
	}

	completer = readline.NewPrefixCompleter(
		readline.PcItem("list",
			itemsList...,
		),
		readline.PcItem("reset",
			readline.PcItem("all"),
			readline.PcItem("user", allUser...),
		),
		readline.PcItem("blocked",
			readline.PcItem("all"),
			readline.PcItem("user", allUser...),
		),
		readline.PcItem("unblocked",
			readline.PcItem("all"),
			readline.PcItem("user", allUser...),
		),
		readline.PcItem("setquota",
			readline.PcItem("all"),
			readline.PcItem("user", allUser...),
		),
		readline.PcItem("status",
			readline.PcItem("user", allUser...),
		),
		readline.PcItem("command"),
		readline.PcItem("clear"),
		readline.PcItem("exit"),
	)

	l.Config.AutoComplete = completer
}
