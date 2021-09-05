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
	d = append(d, []string{pterm.Gray("4"), pterm.LightGreen("user ") + pterm.LightBlue("status"), "Status of a user"})
	d = append(d, []string{pterm.Gray("4"), pterm.LightGreen("user ") + pterm.LightMagenta("newquota"), "Assign new quota to a user"})
	d = append(d, []string{pterm.Gray("5"), pterm.LightGreen("user ") + pterm.LightYellow("reset"), "Reset quota to a user"})
	d = append(d, []string{pterm.Gray("6"), pterm.LightGreen("user ") + pterm.LightRed("blocked"), "Block a user"})
	d = append(d, []string{pterm.Gray("9"), pterm.LightYellow("clear"), "Clear terminal owl "})
	d = append(d, []string{pterm.Gray("10"), pterm.LightRed("exit"), "Exit OWL"})

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
		t.AppendHeader(table.Row{"#", "Email", "Quota MB", "Used MB", "Lock"})
		for num, user := range users {
			t.AppendSeparator()
			if user.Bloquear == true {
				t.AppendRow([]interface{}{pterm.FgRed.Sprint(num), pterm.FgRed.Sprint(strings.Split(user.Email, "@")[0]), pterm.FgRed.Sprint(user.ConvertBitsMB("quota")), pterm.FgRed.Sprint(user.ConvertBitsMB("used")), pterm.FgRed.Sprint(user.Bloquear)})
			} else {
				t.AppendRow([]interface{}{num, pterm.FgCyan.Sprint(strings.Split(user.Email, "@")[0]), user.ConvertBitsMB("quota"), user.ConvertBitsMB("used"), user.Bloquear})
			}
		}
		if footer {
			t.AppendFooter(table.Row{"#", "Email", "Quota MB", "Used MB", "Lock"})
		}
		t.Render()
	}

}
