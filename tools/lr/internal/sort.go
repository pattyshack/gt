package lr

type ByDeclLoc []*Term

func (list ByDeclLoc) Len() int {
	return len(list)
}

func (list ByDeclLoc) Less(i int, j int) bool {
	if list[i].Location.Line == list[j].Location.Line {
		return list[i].Location.Column < list[j].Location.Column
	}

	return list[i].Location.Line < list[j].Location.Line
}

func (list ByDeclLoc) Swap(i int, j int) {
	list[i], list[j] = list[j], list[i]
}

type ByRuleLoc []*Term

func (list ByRuleLoc) Len() int {
	return len(list)
}

func (list ByRuleLoc) Less(i int, j int) bool {
	loc1 := list[i].RuleLocation
	loc2 := list[j].RuleLocation

	if loc1.Line == loc2.Line {
		return loc1.Column < loc2.Column
	}

	return loc1.Line < loc2.Line
}

func (list ByRuleLoc) Swap(i int, j int) {
	list[i], list[j] = list[j], list[i]
}
