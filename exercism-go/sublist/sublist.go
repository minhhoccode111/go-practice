package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	isL1Sublist := isSublist(l1, l2)
	isL2Sublist := isSublist(l2, l1)
	switch {
	case isL1Sublist && isL2Sublist:
		return RelationEqual
	case isL1Sublist:
		return RelationSublist
	case isL2Sublist:
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}

func isSublist(l1, l2 []int) bool {
	if len(l1) == 0 {
		return true
	}
	if len(l1) > len(l2) {
		return false
	}
	for i, v := range l2 {
		// if current value is the first value of l1 and still in bound
		if v == l1[0] && i+len(l1)-1 < len(l2) {
			if isEqual(l1, l2[i:i+len(l1)]) {
				return true
			}
		}
	}
	return false
}

// isEqual checks if two slices are equal (they must have the same length)
func isEqual(l1, l2 []int) bool {
	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}
