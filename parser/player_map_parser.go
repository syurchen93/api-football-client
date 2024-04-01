package parser

import (
	"github.com/syurchen93/api-football-client/response/fixtures"
	"strconv"
	"strings"
)

func ParseFormationStringIntoMap(formation string) map[string]interface{} {
	formationMap := make(map[string]interface{})
	formationParts := strings.Split(formation, "-")
	formationInts := make([]int, len(formationParts))
	for i, part := range formationParts {
		formationInts[i], _ = strconv.Atoi(part)
	}

	formationMap[string(fixtures.LineupPositionDefender)] = formationInts[0]
	last := len(formationParts) - 1
	formationMap[string(fixtures.LineupPositionForward)] = formationInts[last]

	var midfielderCount int
	switch len(formationInts) {
	case 3:
		midfielderCount = formationInts[1]
	case 4:
		if formationInts[1] == 2 && formationInts[2] == 3 {
			// edge case for 4-2-3-1
			formationMap[string(fixtures.LineupPositionDefensiveMidfielder)] = 2
			formationMap[string(fixtures.LineupPositionAttackingMidfielder)] = 3
		} else if formationParts[1] >= formationParts[2] {
			midfielderCount = formationInts[1]
			formationMap[string(fixtures.LineupPositionAttackingMidfielder)] = formationInts[2]
		} else {
			midfielderCount = formationInts[2]
			formationMap[string(fixtures.LineupPositionDefensiveMidfielder)] = formationInts[1]
		}
	case 5:
		midfielderCount = formationInts[2]
		formationMap[string(fixtures.LineupPositionDefensiveMidfielder)] = formationInts[1]
		formationMap[string(fixtures.LineupPositionAttackingMidfielder)] = formationInts[3]
	}
	if midfielderCount > 0 {
		formationMap[string(fixtures.LineupPositionMidfielder)] = midfielderCount
	}
	formationMap["original"] = formation

	return formationMap
}
