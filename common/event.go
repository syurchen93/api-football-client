package common

type EventType string

const (
	Goal         EventType = "Goal"
	Card         EventType = "Card"
	Substitution EventType = "subst"
	Var          EventType = "Var"
)

type EventTypeDetails string

const (
	GoalNormal        EventTypeDetails = "Normal Goal"
	GoalOwn           EventTypeDetails = "Own Goal"
	GoalPenalty       EventTypeDetails = "Penalty"
	GoalPenaltyMissed EventTypeDetails = "Missed Penalty"

	YellowCard EventTypeDetails = "Yellow Card"
	RedCard    EventTypeDetails = "Red Card"

	Substitution1 EventTypeDetails = "Substitution 1"
	Substitution2 EventTypeDetails = "Substitution 2"
	Substitution3 EventTypeDetails = "Substitution 3"
	Substitution4 EventTypeDetails = "Substitution 4"
	Substitution5 EventTypeDetails = "Substitution 5"
	Substitution6 EventTypeDetails = "Substitution 6"

	VarGoalCancelled    EventTypeDetails = "Goal cancelled"
	VarPenaltyConfirmed EventTypeDetails = "Penalty confirmed"
)

const (
	CommentPenaltyShootout = "Penalty Shootout"
	CommentFoul            = "Foul"
)
