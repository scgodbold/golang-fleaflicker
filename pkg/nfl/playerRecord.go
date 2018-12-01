package nfl

type PlayerRecord struct {
	Name      string
	HomeWins  int
	RoadWins  int
	HomeLoses int
	RoadLoses int
	HomeTies  int
	RoadTies  int
}

func (pr *PlayerRecord) Wins() int {
	return pr.HomeWins + pr.RoadWins
}

func (pr *PlayerRecord) Loses() int {
	return pr.HomeLoses + pr.RoadLoses
}

func (pr *PlayerRecord) Ties() int {
	return pr.HomeTies + pr.RoadTies
}

func newPlayerRecord() *PlayerRecord {
	return &PlayerRecord{
		HomeWins:  0,
		HomeLoses: 0,
		RoadWins:  0,
		RoadLoses: 0,
		HomeTies:  0,
		RoadTies:  0,
	}
}
