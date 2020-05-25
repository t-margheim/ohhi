package game

type Line []Cell

func (l *Line) isFinished() bool {
	counts := map[Cell]int{}
	for _, cell := range *l {
		if cell == Empty {
			return false
		}
		counts[cell]++
	}

	target := len(*l) / 2

	if counts[Red] != target || counts[Blue] != target {
		return false
	}
	return true
}

func tryFinish(l Line) (Line, []int) {
	counts := map[Cell]int{}
	for _, cell := range l {
		counts[cell]++
	}
	if counts[Red] == counts[Blue] {
		return l, nil
	}

	target := len(l) / 2

	if counts[Red] == target && counts[Blue] != target {
		return fillRow(l, Blue)

	}

	if counts[Blue] == target && counts[Red] != target {
		return fillRow(l, Red)
	}

	return l, nil
}

func fillRow(r Line, v Cell) (Line, []int) {
	var filled []int
	for i, cell := range r {
		if cell == Empty {
			r[i] = v
			filled = append(filled, i)
		}
	}
	return r, filled
}

func ruleOf3(r Line) (Line, []int) {
	var filled []int
	for i, cell := range r {
		if cell != Empty {
			continue
		}
		// check for checkerboard (red, empty, red)
		if i > 0 && i < len(r)-1 {
			if r[i-1] == r[i+1] && r[i-1] != Empty {
				r[i] = otherCell(r[i-1])
				filled = append(filled, i)
			}
		}

		// check for two-in-a-row (red, red, empty)
		if i > 1 {
			if r[i-2] == r[i-1] && r[i-1] != Empty {
				r[i] = otherCell(r[i-1])
				filled = append(filled, i)
			}
		}
		// check for two-in-a-row (empty, red, red)
		if i < len(r)-2 {
			if r[i+1] == r[i+2] && r[i+1] != Empty {
				r[i] = otherCell(r[i+1])
				filled = append(filled, i)
			}
		}
	}
	return r, filled
}

func ruleOfSame(index int, lines []Line) (Line, []int) {
	var filled []int
	l := lines[index]

	if l.containsTwoEmpties() {
		// fmt.Println("line contains two empties:", l)
		for i, o := range lines {
			if i == index {
				continue
			}
			if l.nearMatch(o) {
				// fmt.Println("NEAR MATCH", l, o)
				for i, c := range l {
					if c == Empty {
						l[i] = otherCell(o[i])
						filled = append(filled, i)
					}
				}
			}
		}
	}

	return l, filled
}

func (l Line) containsTwoEmpties() bool {
	var empties int
	for _, c := range l {
		if c == Empty {
			empties++
		}
	}
	return empties == 2
}

func (l Line) nearMatch(other Line) bool {

	for i, c := range l {
		if c == Empty {
			continue
		}
		if other[i] != c {
			return false
		}
	}
	return true
}
