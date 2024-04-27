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
	return fs == InProgress || fs == FirstHalf ||
		fs == SecondHalf || fs == ExtraTime || fs == Penalty
}

func (fs FixtureStatus) IsFinished() bool {
	return fs == Finished ||
		fs == FinishedAfterExtra ||
		fs == FinishedAfterPenalty ||
		fs == TechnicalLoss || fs == WalkOver ||
		fs == Cancelled
}

func (fs FixtureStatus) IsInFuture() bool {
	return fs == NotStarted || fs == TimeToBeDefined
}
