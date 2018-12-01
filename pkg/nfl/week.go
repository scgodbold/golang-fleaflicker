package nfl

import (
	"fmt"
)

type WeekRecord struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Score float32 `json:"score"`
}

type Match struct {
	Home WeekRecord `json:"home"`
	Away WeekRecord `json:"away"`
}

func (c *NflClient) GetScoresByWeek(week, year int) ([]*Match, error) {
	url := fmt.Sprintf("%v/scores?season=%v&week=%v", c.baseurl, year, week)
	resp, err := c.fetchUrl(url)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	records, err := c.parseWeekRecords(resp)
	if err != nil {
		return nil, err
	}
	homeRecords, awayRecords := c.splitRecords(records)

	var matchs []*Match
	for i, score := range homeRecords {
		newMatch := Match{
			Home: score,
			Away: awayRecords[i],
		}
		matchs = append(matchs, &newMatch)
	}

	return matchs, nil
}

func (c *NflClient) splitRecords(records []WeekRecord) ([]WeekRecord, []WeekRecord) {
	var home []WeekRecord
	var away []WeekRecord
	for i, record := range records {
		if i%2 == 0 {
			home = append(home, record)
		} else {
			away = append(away, record)
		}
	}
	return home, away
}
