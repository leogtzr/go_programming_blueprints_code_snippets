package meander

import (
	"errors"
	"strings"
)

type Cost int8

const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

type CostRange struct {
	From Cost
	To   Cost
}

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (l Cost) String() string {
	for s, v := range costStrings {
		if l == v {
			return s
		}
	}
	return "invalid"
}

func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

func ParseCostRange(s string) (CostRange, error) {
	var r CostRange
	segs := strings.Split(s, "...")
	if len(segs) != 2 {
		return r, errors.New("invalid cost range")
	}
	r.From = ParseCost(segs[0])
	r.To = ParseCost(segs[1])
	return r, nil
}
