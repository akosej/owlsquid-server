package system

import (
	"fmt"
	"github.com/araddon/dateparse"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

func CheckFiles() {
	_, err := File2lines("./config.owl")
	if err != nil {
		AddConfigDefault()
		os.Exit(0)
	}
	//--ACLs and Logs File
	_, err1 := os.Stat(OwlFolderAcls)
	if os.IsNotExist(err1) {
		Run("mkdir " + OwlFolderAcls)
	}
	_, err2 := os.Stat(OwlFolderLogs)
	if os.IsNotExist(err2) {
		Run("mkdir " + OwlFolderLogs)
	}
	_, err3 := File2lines(OwlFolderAcls + "/owl_acl_user_denied")
	if err3 != nil {
		Run("touch " + OwlFolderAcls + "/owl_acl_user_denied")
		Run("chmod -R 777 " + OwlFolderAcls + "/owl_acl_user_denied")
	}
	_, err4 := File2lines(OwlAccesslog)
	if err4 != nil {
		Run("touch " + OwlAccesslog)
		Run("chmod -R 777 " + OwlAccesslog)
	}
}

func Run(arg string) {
	cmdArgs := strings.Fields(arg)
	run := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	_ = run.Run()
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

func AddConfigDefault() {
	Run("touch ./config.owl")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "  *********   ***         ***   ***\n")
	_ = AppendStrFile("./config.owl", "  *********   ***         ***   ***\n")
	_ = AppendStrFile("./config.owl", "  ***   ***   ***   ***   ***   ***\n")
	_ = AppendStrFile("./config.owl", "  ***   ***   ***   ***   ***   ***\n")
	_ = AppendStrFile("./config.owl", "  ***   ***   ***   ***   ***   ***\n")
	_ = AppendStrFile("./config.owl", "  *********   ***************   *********\n")
	_ = AppendStrFile("./config.owl", "  *********   ***************   *********\n")
	_ = AppendStrFile("./config.owl", "--Proxy Guard Owl\n")
	_ = AppendStrFile("./config.owl", "--  Created by Edgar Javier akosej9208@gmail.com  --\n")
	_ = AppendStrFile("./config.owl", "--  Created by Manuel Cabrera mc@infomed.sld.cu   --\n")
	_ = AppendStrFile("./config.owl", "--  System configuration file   --\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#-- Path access.log file\n")
	_ = AppendStrFile("./config.owl", "path.AccessLog=./access.log\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#-- Default quota for users 50 mb 1mb=1048576 Bytes\n")
	_ = AppendStrFile("./config.owl", "default.quota=52428800\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#-- Interface server\n")
	_ = AppendStrFile("./config.owl", "interface.server=eth0\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#--Folder Acl\n")
	_ = AppendStrFile("./config.owl", "folder.owl_acls=./owl_acls\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#------------------------CONNECT MYSQL SERVER-----------------------------\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#--Username\n")
	_ = AppendStrFile("./config.owl", "user.db=userdb\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#--Password\n")
	_ = AppendStrFile("./config.owl", "pass.db=passdb\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#--Ipaddress server\n")
	_ = AppendStrFile("./config.owl", "ip.db=127.0.0.1\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#--Port server\n")
	_ = AppendStrFile("./config.owl", "port.db=3306\n")
	_ = AppendStrFile("./config.owl", "\n")
	_ = AppendStrFile("./config.owl", "#--Name DB\n")
	_ = AppendStrFile("./config.owl", "name.db=owl\n")
	_ = AppendStrFile("./config.owl", "#-------Hours in which the system will restart the quota and rotate the logs  ---------------\n")
	_ = AppendStrFile("./config.owl", "jobs.restart=08:00:00 10:00:00 12:00:00 13:00:00 14:00:00 15:00:00 16:00:00 17:00:00 18:00:00 19:00:00\n")
	fmt.Println("Restart the OWL system, the necessary files have been created.")
}

// Ct --Current time
func Ct() time.Time {
	return time.Now()
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func createFile(path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	//fmt.Println("File Created Successfully", path)
}

func ResetFile(path string) {
	_ = os.Remove(path)
	createFile(path)
	Run("chmod -R 777 " + path)
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
