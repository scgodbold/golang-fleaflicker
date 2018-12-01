package nfl

type Users struct {
	Records map[int]*User
}

func NewUsers() *Users {
	return &Users{
		Records: make(map[int]*User),
	}
}

func (us *Users) AddWeek(matches []*Match) error {
	for _, m := range matches {
		if _, ok := us.Records[m.Home.Id]; !ok {
			us.Records[m.Home.Id] = NewUser(m.Home.Name, m.Home.Id)
		}
		if _, ok := us.Records[m.Away.Id]; !ok {
			us.Records[m.Away.Id] = NewUser(m.Away.Name, m.Away.Id)
		}
		us.Records[m.Home.Id].AddMatch(m)
		us.Records[m.Away.Id].AddMatch(m)
	}
	return nil
}
