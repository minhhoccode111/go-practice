package kindergarten

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	gardenMap map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//     diagram := `
//     VVCCGG
//     VVCCGG`

var dict = map[byte]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !isValidDiagram(diagram) || !isValidChildren(children) {
		return nil, errors.New("invalid diagram")
	}
	sortedChildren := append([]string(nil), children...)
	sort.Strings(sortedChildren)
	diagrams := strings.Fields(diagram)
	gardenMap := map[string][]string{}
	for index, child := range sortedChildren {
		plants := []string{}
		colIndex := index * 2
		plant1 := dict[diagrams[0][colIndex]]
		plant2 := dict[diagrams[0][colIndex+1]]
		plant3 := dict[diagrams[1][colIndex]]
		plant4 := dict[diagrams[1][colIndex+1]]
		plants = append(plants, plant1, plant2, plant3, plant4)
		gardenMap[child] = plants
	}
	return &Garden{gardenMap: gardenMap}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	val, ok := g.gardenMap[child]
	if !ok {
		return nil, false
	}
	return val, true
}

func isValidDiagram(diagram string) bool {
	re := regexp.MustCompile(`\n(\w+)\n(\w+)`)
	rows := re.FindStringSubmatch(diagram)
	if rows == nil {
		return false
	}
	if l1, l2 := len(rows[1]), len(rows[2]); l1 != l2 || l1%2 != 0 {
		return false
	}
	for i, v1 := range rows[1] {
		_, ok1 := dict[byte(v1)]
		_, ok2 := dict[rows[2][i]]
		if !ok1 || !ok2 {
			return false
		}
	}
	return true
}

func isValidChildren(children []string) bool {
	table := map[string]bool{}
	for _, v := range children {
		_, ok := table[v]
		if ok {
			return false
		}
		table[v] = true
	}
	return true
}
