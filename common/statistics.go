package common

type StatsType string

const (
	ShotsOnGoal        StatsType = "Shots on Goal"
	ShotsOffGoal       StatsType = "Shots off Goal"
	ShotsInsideBox     StatsType = "Shots insidebox"
	ShotsOutsideBox    StatsType = "Shots outsidebox"
	TotalShots         StatsType = "Total Shots"
	BlockedShots       StatsType = "Blocked Shots"
	Fouls              StatsType = "Fouls"
	CornerKicks        StatsType = "Corner Kicks"
	Offsides           StatsType = "Offsides"
	BallPossession     StatsType = "Ball Possession"
	YellowCards        StatsType = "Yellow Cards"
	RedCards           StatsType = "Red Cards"
	GoalkeeperSaves    StatsType = "Goalkeeper Saves"
	TotalPasses        StatsType = "Total passes"
	PassesAccurate     StatsType = "Passes accurate"
	PassesPercentage   StatsType = "Passes %"
)

func (st StatsType) IsPercentage() bool {
	return st == BallPossession || st == PassesPercentage
}