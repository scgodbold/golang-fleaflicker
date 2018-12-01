package nfl

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *NflClient) parseWeekRecords(data io.Reader) ([]WeekRecord, error) {
	var records []WeekRecord
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return nil, err
	}
	doc.Find(".scoreboard").Each(func(i int, item *goquery.Selection) {
		link, ok := item.Find("a").Attr("href")
		if !ok {
			return
		}
		season_pid := strings.TrimPrefix(link, fmt.Sprintf("/nfl/leagues/%v/teams/", c.LeagueId))
		pid := strings.Split(season_pid, "?")[0]
		pidInt, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		scoreFloat, err := strconv.ParseFloat(item.Find(".text-right span").Text(), 32)
		if err != nil {
			return
		}

		record := WeekRecord{
			Name:  item.Find("a").Text(),
			Score: float32(scoreFloat),
			Id:    int(pidInt),
		}
		records = append(records, record)
	})

	return records, nil
}
