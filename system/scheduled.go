package system

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"sort"
	"strconv"
	"strings"
	"time"
)

func newJobTicker(now time.Time, HourToTick int, MinuteToTick int, SecondToTick int) JobTicker {
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), HourToTick, MinuteToTick, SecondToTick, 0, now.Location())
	if nextTick.Before(now) {
		nextTick = nextTick.Add(IntervalPeriod)
	}
	fmt.Println(nextTick.Sub(now))
	return JobTicker{time.NewTimer(nextTick.Sub(now))}
}

func addJobs(now time.Time, HourToTick int, MinuteToTick int, SecondToTick int) {
	fmt.Printf("Run job at  %v:%v:%v missing ->", HourToTick, MinuteToTick, SecondToTick)
	Jobs = append(Jobs, newJobTicker(now, HourToTick, MinuteToTick, SecondToTick))
}

func RunJobsRestartCuota() {
	//--Time extract
	timesAll := strings.Split(OwlRestart, " ")
	sort.Strings(timesAll)
	ct := Ct()

	for _, times := range timesAll {
		part := strings.Split(times, ":")
		//--Convert to integers
		hour, _ := strconv.Atoi(part[0])
		minute, _ := strconv.Atoi(part[1])
		second, _ := strconv.Atoi(part[2])

		//--That the first job is when the execution time is after the current time
		if ct.Hour() < hour {
			addJobs(ct, hour, minute, second)
		} else if ct.Hour() == hour && ct.Minute() < minute {
			addJobs(ct, hour, minute, second)
		} else if ct.Hour() == hour && ct.Minute() == minute && ct.Second() < second {
			addJobs(ct, hour, minute, second)
		} else {
			JobsEnd = append(JobsEnd, times)
		}
	}

	for _, times := range JobsEnd {
		part := strings.Split(times, ":")
		hour, _ := strconv.Atoi(part[0])
		minute, _ := strconv.Atoi(part[1])
		second, _ := strconv.Atoi(part[2])
		addJobs(ct, hour, minute, second)
	}

	//--Check the jobs and execute the task
	for {
		for _, v := range Jobs {
			a := <-v.t.C
			fmt.Println(a.String())
			time.Local = ct.Location()
			ctn := Ct()
			cutName := strings.Split(ctn.String(), " ")
			cutName2 := strings.Split(cutName[1], ".")
			_, _ = copyFile(OwlAccesslog, OwlFolderLogs+"/access_"+cutName[0]+"_"+cutName2[0]+".log")
			_, _ = RunString("echo '' >" + OwlAccesslog)
			allKeyRedis()
			for _, user := range AllUser {
				if _, err := RDB.Pipelined(CTX, func(rdb redis.Pipeliner) error {
					rdb.HSet(CTX, user, "used", 0)
					rdb.HSet(CTX, user, "activa", 1)
					rdb.HSet(CTX, user, "bloquear", 0)
					return nil
				}); err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println("Scheduled task executed, user quota has been reset ")
			//--Restart order
			addJobs(ctn, ctn.Hour(), ctn.Minute(), ctn.Second())
		}
	}
}
