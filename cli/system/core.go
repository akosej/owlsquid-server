package system

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/pterm/pterm"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func VerifyConfiguration() {
	_, errFolderConfig := os.Stat(FolderConfig)
	if os.IsNotExist(errFolderConfig) {
		_ = os.Mkdir(FolderConfig, 0755)
	}
	_, err := File2lines(FileConfig)
	if err != nil {
		createFile(FileConfig)
		_ = AppendStrFile(FileConfig, "\n")
		_ = AppendStrFile(FileConfig, "  *********   ***         ***   ***\n")
		_ = AppendStrFile(FileConfig, "  *********   ***         ***   ***\n")
		_ = AppendStrFile(FileConfig, "  ***   ***   ***   ***   ***   ***\n")
		_ = AppendStrFile(FileConfig, "  ***   ***   ***   ***   ***   ***\n")
		_ = AppendStrFile(FileConfig, "  ***   ***   ***   ***   ***   ***\n")
		_ = AppendStrFile(FileConfig, "  *********   ***************   *********\n")
		_ = AppendStrFile(FileConfig, "  *********   ***************   *********\n")
		_ = AppendStrFile(FileConfig, "--Proxy Guard Owl\n")
		_ = AppendStrFile(FileConfig, "--  Created by Edgar Javier akosej9208@gmail.com  --\n")
		_ = AppendStrFile(FileConfig, "--  Created by Manuel Cabrera mc@infomed.sld.cu   --\n")
		_ = AppendStrFile(FileConfig, "--  System configuration file CLI  --\n")
		_ = AppendStrFile(FileConfig, "\n")
		_ = AppendStrFile(FileConfig, "#--Ipaddress server REDIS\n")
		_ = AppendStrFile(FileConfig, "ip.db=127.0.0.1\n")
		_ = AppendStrFile(FileConfig, "\n")
		_ = AppendStrFile(FileConfig, "#-- Password server REDIS\n")
		_ = AppendStrFile(FileConfig, "pass.db=test\n")
		_ = AppendStrFile(FileConfig, "\n")
		_ = AppendStrFile(FileConfig, "#-- Port server REDIS\n")
		_ = AppendStrFile(FileConfig, "port.db=6379\n")
		_ = AppendStrFile(FileConfig, "\n")
		_ = AppendStrFile(FileConfig, "#--Entity domain\n")
		_ = AppendStrFile(FileConfig, "entity.domain=infomed.sld.cu\n")
		_, errLocalBin := os.Stat(HOME + "/.local/bin")
		if os.IsNotExist(errLocalBin) {
			_ = os.Mkdir(HOME+"/.local/bin", 0755)
		}
		_, _ = copyFile("./owlcli", HOME+"/.local/bin/owlcli")
		Run("chmod +x " + HOME + "/.local/bin/owlcli")
		pterm.FgRed.Println("The system did not find the configuration file, so it generated one by default, \nyou must configure it with its particularities in: " + FileConfig)
		pterm.FgYellow.Println("You can run the cli directly using the command:")
		pterm.FgGreen.Println("owlcli")
		os.Exit(0)
	}
}

func Config(data string) string {
	lines, err := File2lines(FileConfig)
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

func SubtractDates(date string) int {
	a, _ := strconv.ParseInt(date, 10, 64)
	ta := time.Unix(a, 0)
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
	parse, _ := dateparse.ParseLocal(ta.String())
	return int(time.Now().Sub(parse).Seconds())
}

func Contains(s []string, searcher string) bool {
	i := sort.SearchStrings(s, searcher)
	return i < len(s) && s[i] == searcher
}

// Ct --Current time
func Ct() time.Time {
	return time.Now()
}

func ResetFile(path string) {
	_ = os.Remove(path)
	createFile(path)
	_, _ = RunString("chmod -R 777 " + path)
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
