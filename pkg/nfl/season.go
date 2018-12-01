package nfl

type SeasonRecord struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Wins             int     `json:"wins"`
	Loses            int     `json:"loses"`
	PointsFor        float32 `json:"points_for"`
	PointsForAvg     float32 `json:"points_for_avg"`
	PointsAgainst    float32 `json:"points_against"`
	PointsAgainstAvg float32 `json:"points_against_avg"`
}

// func (c *NflClient) GetSeason(year int) ([]*SeasonRecord, error) {
// 	url := fmt.Sprintf("%v?season=%v", c.baseurl, year, week)
// 	resp, err := c.fetchUrl(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Close()
// 	seasonRecords, err := c.parseSeasonRecords(resp)
// }
