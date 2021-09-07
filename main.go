//	*********   ***         ***   ***
//	*********   ***         ***   ***
//	***   ***   ***   ***   ***   ***
//	***   ***   ***   ***   ***   ***
//	***   ***   ***   ***   ***   ***
//	*********   ***************   *********
//	*********   ***************   *********
//--Proxy Guard Owl
package main

import (
	"github.com/akosej9208/owlsquid/system"
	"github.com/hpcloud/tail"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//-------Check that the necessary files exist
	system.CheckFiles()
	//-------Connect to the redis server
	system.ConnectServerRedis()
	//-------JOBS RESTARTS SYSTEM
	go func() {
		system.RunJobsRestartCuota()
	}()
	//-------Inject log reading job
	go func() {
		for {
			select {
			//---Write to OwlActivesRequest all access_request
			case textRequestUser := <-system.JobCheckUser:
				orderJob := strings.Split(textRequestUser, "---")
				bytes, _ := strconv.ParseFloat(orderJob[1], 10)
				system.ActionRun(orderJob[0], bytes, orderJob[2], orderJob[3], orderJob[4])
			}
		}
	}()
	//-------Read real-time access.log from squid
	t, _ := tail.TailFile(system.OwlAccesslog, tail.Config{Follow: true})

	sort.Strings(system.ConnectType)

	for line := range t.Lines {
		//time duration client_address result_code bytes request_method url rfc931 hierarchy_code type
		// 0        1       2               3       4       5           6       7           8       9
		segment := strings.Fields(line.Text)
		//-- extract date data
		extractDateData := strings.Split(segment[0], ".")
		secondOfDifferences := system.SubtractDates(extractDateData[0])
		//-- If the log entry was executed in less than 10 seconds
		if secondOfDifferences < 10 {
			//--If it contains an @ it means that there is a user
			if strings.Contains(segment[7], "@") {
				//--Check the connection type from the connectType list
				if system.Contains(system.ConnectType, segment[3]) {
					//--If the request size is> 0
					if segment[4] > "0" {
						// -- Send request data to be processed in the job
						system.JobCheckUser <- segment[7] + "---" + segment[4] + "---" + segment[0] + "---" + segment[6] + "---" + segment[2]
					}
				}
			}
		}
	}
}
