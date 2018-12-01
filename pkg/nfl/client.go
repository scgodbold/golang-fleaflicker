package nfl

import (
	"fmt"
)

const fleaflickerURL = "https://fleaflicker.com/nfl"

type NflClient struct {
	LeagueId int
	baseurl  string
}

func New(id int) *NflClient {
	client := new(NflClient)
	client.LeagueId = id
	client.baseurl = fmt.Sprintf("%v/leagues/%v", fleaflickerURL, id)
	return client
}
