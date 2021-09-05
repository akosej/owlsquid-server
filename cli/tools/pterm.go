package tools

import (
	"github.com/pterm/pterm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Change this to time.Millisecond*200 to speed up the demo.
// Useful when debugging.
const second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func IntroScreen() {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("owl", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("cli", pterm.NewStyle(pterm.FgLightRed))).
		Srender()
	pterm.Print(ptermLogo)
	pterm.Println()

	//introSpinner, _ := pterm.DefaultSpinner.WithShowTimer(false).WithRemoveWhenDone(true).Start("Waiting for 3 seconds...")
	//time.Sleep(second)
	//for i := 2; i > 0; i-- {
	//	if i > 1 {
	//		introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
	//	} else {
	//		introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
	//	}
	//	time.Sleep(second)
	//}
	//introSpinner.Stop()
}

func InstalledTree() {
	leveledList := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "C:"},
		pterm.LeveledListItem{Level: 1, Text: "Go"},
		pterm.LeveledListItem{Level: 1, Text: "Windows"},
		pterm.LeveledListItem{Level: 1, Text: "Programs"},
	}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 2, Text: s})
		}
		if s == "pseudo-chrome" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Tabs"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Extensions"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "Refined GitHub"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "GitHub Dark Theme"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Bookmarks"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "OWL"})
		}
	}

	pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList(leveledList)).Render()
}

func InstallingPseudoList() {
	pterm.DefaultSection.Println("Installing pseudo programs")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.Title = "Installing " + pseudoProgramList[i]
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(second / 2)
	}
	p.Stop()
}

func ListPrinter() {
	pterm.NewBulletListFromString(`Good bye
 Have a nice day!`, " ").Render()
}

func FadeText() {
	from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
	to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients first point.

	str := "If your terminal has TrueColor support, you can use RGB colors!\nYou can even fade them :)"
	strs := strings.Split(str, "")
	var fadeInfo string // String which will be used to print info.
	// For loop over the range of the string length.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}
	pterm.Info.Println(fadeInfo)
}

func InstalledProgramsSize() {
	d := pterm.TableData{{"Program Name", "Status", "Size"}}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			d = append(d, []string{s, pterm.LightGreen("pass"), strconv.Itoa(RandomInt(7, 200)) + "mb"})
		} else {
			d = append(d, []string{pterm.LightRed(s), pterm.LightRed("fail"), "0mb"})
		}
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
}

func PseudoApplicationHeader() *pterm.TextPrinter {
	return pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"Pseudo Application created with PTerm")
}

func Clear() {
	print("\033[H\033[2J")
	print("\n")
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
