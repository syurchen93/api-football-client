package fixtures

type Lineup struct {
	Team        TeamLineup     `json:"team"`
	Formation   Formation      `json:"formation"`
	StartXI     []PlayerLineup `mapstructure:"startXI" json:"startXI"`
	Substitutes []PlayerLineup `json:"substitutes"`
	Coach       Coach          `json:"coach"`
}

type TeamLineup struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	KitColors struct {
		FieldPlayer KitColor `mapstructure:"player" json:"player"`
		Goalkeeper  KitColor `json:"goalkeeper"`
	} `mapstructure:"colors" json:"colors"`
}

type PlayerLineup struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Number   int            `json:"number"`
	Position LineupPosition `mapstructure:"pos" json:"pos"`
	// For some reason this is always null
	Grid GridCell `json:"grid"`
}

type Coach struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type KitColor struct {
	Primary string `json:"primary"`
	Number  string `json:"number"`
	Border  string `json:"border"`
}

type GridCell struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type Formation struct {
	Defenders            int    `mapstructure:"D" json:"D"`
	DefenciveMidfielders int    `mapstructure:"DM" json:"DM"`
	Midfielders          int    `mapstructure:"M" json:"M"`
	AttackingMidfielders int    `mapstructure:"AM" json:"AM"`
	Forwards             int    `mapstructure:"F" json:"F"`
	Original             string `mapstructre:"original" json:"original"`
}

type LineupPosition string

const (
	LineupPositionGoalkeeper          LineupPosition = "G"
	LineupPositionDefender            LineupPosition = "D"
	LineupPositionDefensiveMidfielder LineupPosition = "DM"
	LineupPositionMidfielder          LineupPosition = "M"
	LineupPositionAttackingMidfielder LineupPosition = "AM"
	LineupPositionForward             LineupPosition = "F"
)
