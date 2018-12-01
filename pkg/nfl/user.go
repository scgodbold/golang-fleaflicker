package nfl

import (
	"fmt"
)

type User struct {
	Id        int
	Name      string
	HomeWins  int
	HomeLoses int
	HomeTies  int
	RoadWins  int
	RoadLoses int
	RoadTies  int
	Records   map[int]*PlayerRecord
}

func (u *User) Wins() int {
	return u.HomeWins + u.RoadWins
}

func (u *User) Loses() int {
	return u.HomeLoses + u.RoadLoses
}

func (u *User) Ties() int {
	return u.HomeTies + u.RoadTies
}

func (u *User) WinPercentage() float32 {
	return (float32(u.Wins()) + (float32(u.Ties()) / 2.0)) / (float32(u.Wins()) + float32(u.Loses()) + float32(u.Ties()))
}

func (u *User) AddMatch(m *Match) error {
	var oppId int
	var oppName string
	var oppScore float32
	if m.Home.Id == u.Id {
		// Home Matchup
		oppScore = m.Away.Score
		oppId = m.Away.Id
		oppName = m.Away.Name
		if _, ok := u.Records[oppId]; !ok {
			u.Records[oppId] = newPlayerRecord()
		}

		// Update Names
		u.Records[oppId].Name = oppName
		u.Name = m.Home.Name

		if oppScore < m.Home.Score {
			u.Records[oppId].HomeWins += 1
			u.HomeWins += 1
		} else if oppScore > m.Home.Score {
			u.Records[oppId].HomeLoses += 1
			u.HomeLoses += 1
		} else {
			u.Records[oppId].HomeTies += 1
			u.HomeTies += 1
		}
	} else {
		// Away Matchup
		oppScore = m.Home.Score
		oppId = m.Home.Id
		oppName = m.Home.Name
		if _, ok := u.Records[oppId]; !ok {
			u.Records[oppId] = newPlayerRecord()
		}

		// Update Names
		u.Records[oppId].Name = oppName
		u.Name = m.Away.Name

		if oppScore < m.Away.Score {
			u.Records[oppId].RoadWins += 1
			u.RoadWins += 1
		} else if oppScore > m.Away.Score {
			u.Records[oppId].RoadLoses += 1
			u.RoadLoses += 1
		} else {
			u.Records[oppId].RoadTies += 1
			u.RoadTies += 1
		}
	}
	return nil
}

func (u *User) PrintRecords() {
	fmt.Printf("Head to Head Records for: %v\n", u.Name)
	for _, record := range u.Records {
		total := record.Wins() + record.Loses()
		if total == 0 {
			continue
		}
		fmt.Printf("\t%v: Total Matches %v, Wins %v, Loses %v\n", record.Name,
			total, record.Wins(), record.Loses())
	}
	fmt.Printf("Total Record: %v-%v-%v\n", u.Wins(), u.Loses(), u.Ties())
	fmt.Printf("--------------------------------------------------------\n")
}

func NewUser(name string, id int) *User {
	return &User{
		Id:        id,
		Name:      name,
		HomeWins:  0,
		HomeLoses: 0,
		HomeTies:  0,
		RoadWins:  0,
		RoadLoses: 0,
		RoadTies:  0,
		Records:   make(map[int]*PlayerRecord),
	}
}
