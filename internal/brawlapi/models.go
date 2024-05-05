package brawlapi

type PlayerInfo struct {
	Name                                 string        `json:"name"`
	Tag                                  string        `json:"tag"`
	Club                                 PlayerClub    `json:"club"`
	Icon                                 PlayerIcon    `json:"icon"`
	Trophies                             int           `json:"trophies"`
	ExpLevel                             int           `json:"expLevel"`
	HighestTrophies                      int           `json:"highestTrophies"`
	ThreeVSThreeVictories                int           `json:"3vs3Victories"`
	SoloVictories                        int           `json:"soloVictories"`
	DuoVictories                         int           `json:"duoVictories"`
}

type PlayerClub struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type PlayerIcon struct {
	Id int `json:"id"`
}

type BrawlerStat struct {
	Gadjets         []Accessory `json:"gadjets"`
	StarPowers      []StarPower `json:"starPowers"`
	Id              int         `json:"id"`
	Rank            int         `json:"rank"`
	Trophies        int         `json:"trophies"`
	HighestTrophies int         `json:"highestTrophies"`
	Power           int         `json:"power"`
	Gears           []GearStat  `json:"gears"`
	Name            string      `json:"name"`
}

type Accessory struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type StarPower struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type GearStat struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Level int    `json:"level"`
}

type ClubInfo struct {
	Tag              string       `json:"tag"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	Type             string       `json:"type"`
	BadgeID          int          `json:"badgeId"`
	RequiredTrophies int          `json:"requiredTrophies"`
	Trophies         int          `json:"trophies"`
	Members          []ClubMember `json:"members"`
}

type ClubMember struct {
	Tag       string     `json:"tag"`
	Name      string     `json:"name"`
	NameColor string     `json:"nameColor"`
	Role      string     `json:"role"`
	Trophies  int        `json:"trophies"`
	Icon      PlayerIcon `json:"icon"`
}

type ClientError struct {
	Reason  string `json:"reason"`
}

func (ce ClientError) Error() string {
	return ce.Reason
}
