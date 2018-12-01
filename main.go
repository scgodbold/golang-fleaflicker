package main

import (
	"github.com/scgodbold/golang-fleaflicker/pkg/nfl"
	"log"
	"os"
	"strconv"
)

func getIdEnv() int {
	strId := os.Getenv("FLEAFLICKER_LEAGUE_ID")
	intVal, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(intVal)
}

func main() {
	nflClient := nfl.New(getIdEnv())
	records := nfl.NewUsers()
	for i := 1; i < 16; i++ {
		results, err := nflClient.GetScoresByWeek(i, 2016)
		if err != nil {
			log.Fatal(err)
		}
		records.AddWeek(results)
	}
	for i := 1; i < 17; i++ {
		results, err := nflClient.GetScoresByWeek(i, 2017)
		if err != nil {
			log.Fatal(err)
		}
		records.AddWeek(results)
	}
	for i := 1; i < 13; i++ {
		results, err := nflClient.GetScoresByWeek(i, 2018)
		if err != nil {
			log.Fatal(err)
		}
		records.AddWeek(results)
	}
	for _, u := range records.Records {
		u.PrintRecords()
	}
}
