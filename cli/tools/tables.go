package tools

import (
	"cli/system"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pterm/pterm"
	"os"
	"strings"
)

// PrintCommandList --Print command list
func PrintCommandList() {
	d := pterm.TableData{{"#", "Command", "Description"}}
	d = append(d, []string{pterm.Gray("1"), pterm.LightBlue("list ") + "all", "List all registered users"})
	d = append(d, []string{pterm.Gray("2"), pterm.LightBlue("list ") + pterm.LightGreen("actives"), "List actives users"})
	d = append(d, []string{pterm.Gray("3"), pterm.LightBlue("list ") + pterm.LightRed("blocked"), "List blocked users"})
	d = append(d, []string{pterm.Gray("4"), pterm.LightBlue("status ") + pterm.LightGreen("user ") + pterm.LightRed("(username)"), "Status of a user"})
	d = append(d, []string{pterm.Gray("5"), pterm.LightBlue("reset ") + pterm.LightGreen("all"), "Reset the quota of all users"})
	d = append(d, []string{pterm.Gray("6"), pterm.LightBlue("reset ") + pterm.LightGreen("user ") + pterm.LightRed("(username)"), "Reset quota to a user"})

	d = append(d, []string{pterm.Gray("7"), pterm.LightBlue("blocked ") + pterm.LightGreen("all"), "Block all users"})
	d = append(d, []string{pterm.Gray("8"), pterm.LightBlue("blocked ") + pterm.LightGreen("user ") + pterm.LightRed("(username)"), "Block a user"})

	d = append(d, []string{pterm.Gray("9"), pterm.LightBlue("unblocked ") + pterm.LightGreen("all"), "Unblock all users"})
	d = append(d, []string{pterm.Gray("10"), pterm.LightBlue("unblocked ") + pterm.LightGreen("user ") + pterm.LightRed("(username)"), "Unlock a user"})

	d = append(d, []string{pterm.Gray("11"), pterm.LightBlue("setquota ") + pterm.LightGreen("all"), "Assign new quota to all users"})
	d = append(d, []string{pterm.Gray("12"), pterm.LightBlue("setquota ") + pterm.LightGreen("user ") + pterm.LightRed("(username)") + pterm.LightYellow("(mb)"), "Assign new quota to a user"})

	d = append(d, []string{pterm.Gray("13"), pterm.LightBlue("command"), "command list "})
	d = append(d, []string{pterm.Gray("14"), pterm.LightBlue("clear"), "Clear terminal owl "})
	d = append(d, []string{pterm.Gray("15"), pterm.LightBlue("exit"), "Exit OWL"})

	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
}

// PrintUserTable --Print user table
func PrintUserTable(users []system.Model, title string, footer bool) {
	if len(users) == 0 {
		pterm.FgYellow.Println("No results found.")
	} else {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle(title)
		t.AppendHeader(table.Row{"#", "Email", "Quota MB", "Used MB", "Ipaddress", "Lock", "Last url"})
		for num, user := range users {
			t.AppendSeparator()
			if user.Bloquear == true {
				t.AppendRow([]interface{}{
					pterm.FgRed.Sprint(num),
					pterm.FgRed.Sprint(strings.Split(user.Email, "@")[0]),
					pterm.FgRed.Sprint(user.ConvertBitsMB("quota")),
					pterm.FgRed.Sprint(user.ConvertBitsMB("used")),
					pterm.FgRed.Sprint(user.IpRemote),
					pterm.FgRed.Sprint(user.Bloquear),
					pterm.FgRed.Sprint(user.Last_url),
				})
			} else {
				t.AppendRow([]interface{}{
					pterm.FgCyan.Sprint(num),
					pterm.FgCyan.Sprint(strings.Split(user.Email, "@")[0]),
					pterm.FgCyan.Sprint(user.ConvertBitsMB("quota")),
					pterm.FgCyan.Sprint(user.ConvertBitsMB("used")),
					pterm.FgCyan.Sprint(user.IpRemote),
					pterm.FgCyan.Sprint(user.Bloquear),
					pterm.FgCyan.Sprint(user.Last_url),
				})
			}
		}
		if footer {
			t.AppendFooter(table.Row{"#", "Email", "Quota MB", "Used MB", "Ipaddress", "Lock", "Last url"})
		}
		t.Render()
	}
}
