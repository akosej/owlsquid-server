package system

import (
	"bufio"
	"os"
	"strings"
	"time"
)

//--Active monitoring of a file
func WatchFile(filePath string) bool {
	//-----------------------------------------------
	//job_msg <- "Waiting for requests to the proxy..."
	//-----------------------------------------------
	initialStat, _ := os.Stat(filePath)
	for {
		stat, _ := os.Stat(filePath)
		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			//	----------------------------------------
			file, _ := os.Open(filePath)
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var text []string
			for scanner.Scan() {
				text = append(text, scanner.Text())
			}
			file.Close()
			//	----------------------------------------
			allLines := len(text)
			startLine := 0
			linesPerBlock := 15
			endLine := startLine + linesPerBlock
			//	----------------------------------------
			for endLine < allLines {
				//-----------------------------------------------------------------
				//-Extraction of the data received in the actives request----------
				//-----------------------------------------------------------------
				line := text[startLine:endLine]
				//--Username
				lineUser := strings.Split(line[13], " ")
				username := lineUser[1]

				//--Request Size
				lineRequestSize := strings.Split(line[1], " ")
				reqSize := lineRequestSize[5]

				//--Ipaddress remote user
				lineIpaddressRemote := strings.Split(line[4], " ")
				ipRemoteUser := lineIpaddressRemote[1]

				//--Ipaddress Local
				//lineIpaddressLocal := strings.Split(line[5], " ")
				//ip_local_server := lineIpaddressLocal[1]

				//--Url request
				lineUrl := strings.Split(line[7], " ")
				urlRequest := lineUrl[1]

				//--Date request
				lineDate := strings.Split(line[12], " ")
				dateRequest := lineDate[1]

				//--Connection
				//lineConnection := strings.Split(line[0], " ")
				//nameConnection := lineConnection[1]
				//-----------------------------------------------------------------
				if len(username) > 1 {
					JobCheckUser <- username + "---" + reqSize + "---" + dateRequest + "---" + urlRequest + "---" + ipRemoteUser
				}
				startLine = endLine + 1
				endLine = startLine + linesPerBlock
			}
			//	----------------------------------------
			numRequest := allLines / 15
			cantConProccess := len(Connects)
			if cantConProccess > numRequest {
				dif := cantConProccess - numRequest
				Connects = append(Connects[:0], Connects[dif:]...)
			}
			return true
		}
		//--Delay the cycle for a second
		time.Sleep(1 * time.Second)
	}
}
