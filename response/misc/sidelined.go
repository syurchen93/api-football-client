package misc

import "time"

type SidelinedEntry struct {
	Type  string     `json:"type"`
	Start time.Time  `json:"start"`
	End   *time.Time `json:"end"`
}
