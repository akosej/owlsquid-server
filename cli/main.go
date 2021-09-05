package main

import (
	"bufio"
	"cli/system"
	"cli/tools"
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"strings"
	"time"
)

func main() {
	system.VerifyConfiguration()
	tools.Clear()
	system.ConnectServerRedis()
	tools.IntroScreen()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(pterm.Cyan("Command: "))
		command, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
		cleanCommand := strings.TrimRight(command, "\r\n")
		cutCommand := strings.Split(cleanCommand, " ")
		switch arg := cutCommand[0]; arg {
		case "list":
			if len(cutCommand) < 2 {
				tools.PrintCommandList()
				continue
			}
			switch arg := cutCommand[1]; arg {
			case "all":
				tools.PrintUserTable(system.GetAllUserStoredRedis("all"), "List of users saved by owl", true)
			case "actives":
				tools.PrintUserTable(system.GetAllUserStoredRedis("actives"), "List actives users", true)
			case "blocked":
				tools.PrintUserTable(system.GetAllUserStoredRedis("blocked"), "List blocked users", true)
			default:
				tools.PrintCommandList()
			}
		case "user":
			if len(cutCommand) < 2 {
				tools.PrintCommandList()
				continue
			}
			fmt.Print(pterm.Cyan("Specify the user : "))
			switch arg := cutCommand[1]; arg {
			case "status":
				commandU, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
				cleanCommandU := strings.TrimRight(commandU, "\r\n")
				cutCommandU := strings.Split(cleanCommandU, " ")
				tools.PrintUserTable(system.GetUserStoredRedis(cutCommandU[0]), "User information", false)
			case "newquota":
				commandU, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
				cleanCommandU := strings.TrimRight(commandU, "\r\n")
				cutCommandU := strings.Split(cleanCommandU, " ")
				fmt.Print(pterm.Cyan("Specify the new quota in MB: "))
				commandC, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
				cleanCommandC := strings.TrimRight(commandC, "\r\n")
				cutCommandC := strings.Split(cleanCommandC, " ")
				result := system.OptUserStoredRedis(cutCommandU[0], "newquota", cutCommandC[0])
				if result {
					pterm.FgCyan.Println("User quota has been changed successfully ")
					tools.PrintUserTable(system.GetUserStoredRedis(cutCommandU[0]), "User information", false)
				} else {
					pterm.FgRed.Println("An error has occurred, verify that the user is correctly typed, remember that you should not put @domain")
				}
			case "reset":
				commandU, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
				cleanCommandU := strings.TrimRight(commandU, "\r\n")
				cutCommandU := strings.Split(cleanCommandU, " ")
				result := system.OptUserStoredRedis(cutCommandU[0], "reset", "")
				if result {
					pterm.FgCyan.Println("The user's browsing quota has been reset")
					tools.PrintUserTable(system.GetUserStoredRedis(cutCommandU[0]), "User information", false)
				} else {
					pterm.FgRed.Println("An error has occurred, verify that the user is correctly typed, remember that you should not put @domain")
				}
			case "blocked":
				commandU, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
				cleanCommandU := strings.TrimRight(commandU, "\r\n")
				cutCommandU := strings.Split(cleanCommandU, " ")
				result := system.OptUserStoredRedis(cutCommandU[0], "blocked", "")
				if result {
					pterm.FgCyan.Println("User has been successfully blocked ")
					tools.PrintUserTable(system.GetUserStoredRedis(cutCommandU[0]), "User information", false)
				} else {
					pterm.FgRed.Println("An error has occurred, verify that the user is correctly typed, remember that you should not put @domain")
				}
			default:
				tools.PrintCommandList()
			}
		case "clear":
			tools.Clear()
			tools.IntroScreen()
		case "help":
			tools.PrintCommandList()
		case "exit":
			pterm.FgCyan.Println("See you soon!!!")
			time.Sleep(2 * time.Second)
			tools.Clear()
			os.Exit(0)
		default:
			tools.PrintCommandList()
		}
	}

	//pseudoApplicaztionHeader()
	//time.Sleep(second)
	//installingPseudoList()
	//time.Sleep(second * 2)
	//pterm.DefaultSection.WithLevel(2).Println("Program Install Report")
	//installedProgramsSize()
	//time.Sleep(second * 4)
	//pterm.DefaultSection.Println("Tree Printer")
	//installedTree()
	//time.Sleep(second * 4)
	//pterm.DefaultSection.Println("TrueColor Support")
	//fadeText()
	//time.Sleep(second)
	//pterm.DefaultSection.Println("Bullet List Printer")
	//listPrinter()
}
