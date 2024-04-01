package common

type FixtureStatus string

const (
	TimeToBeDefined      FixtureStatus = "TBD"
	NotStarted           FixtureStatus = "NS"
	FirstHalf            FixtureStatus = "1H"
	Halftime             FixtureStatus = "HT"
	SecondHalf           FixtureStatus = "2H"
	ExtraTime            FixtureStatus = "ET"
	BreakTime            FixtureStatus = "BT"
	Penalty              FixtureStatus = "P"
	Suspended            FixtureStatus = "SUSP"
	Interrupted          FixtureStatus = "INT"
	Finished             FixtureStatus = "FT"
	FinishedAfterExtra   FixtureStatus = "AET"
	FinishedAfterPenalty FixtureStatus = "PEN"
	Postponed            FixtureStatus = "PST"
	Cancelled            FixtureStatus = "CANC"
	Abandoned            FixtureStatus = "ABD"
	TechnicalLoss        FixtureStatus = "AWD"
	WalkOver             FixtureStatus = "WO"
	InProgress           FixtureStatus = "LIVE"
)

func (fs FixtureStatus) IsLive() bool {
	return fs != Finished &&
		fs != FinishedAfterExtra &&
		fs != FinishedAfterPenalty &&
		fs != Postponed &&
		fs != Cancelled &&
		fs != Abandoned &&
		fs != TechnicalLoss &&
		fs != WalkOver
}

func (fs FixtureStatus) IsGameplayHappening() bool {
	return fs == FirstHalf ||
		fs == SecondHalf ||
		fs == ExtraTime ||
		fs == Penalty
}
