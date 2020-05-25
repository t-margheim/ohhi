package game

import (
	"fmt"
	"reflect"
)

var (
	rows = []Line{
		{0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 2, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 2, 2},
		{2, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0, 0, 1},
		// {Red, Empty, Red, Empty, Empty, Empty},
		// {Empty, Blue, Empty, Empty, Empty, Empty},
		// {Red, Empty, Empty, Empty, Blue, Empty},
		// {Empty, Blue, Empty, Empty, Empty, Empty},
		// {Empty, Empty, Red, Empty, Empty, Red},
		// {Empty, Empty, Empty, Empty, Empty, Empty},
	}
)

func New() Board {
	var b Board
	b.Rows = rows
	for range rows {
		c := make(Line, len(rows))
		b.Columns = append(b.Columns, c)
	}
	for x, r := range rows {
		for y, v := range r {
			b.Columns[y][x] = v
		}
	}
	return b
}

type Board struct {
	Rows    []Line
	Columns []Line
}

func (b Board) Print() {
	for _, row := range b.Rows {
		for _, cell := range row {
			cell.Print()
		}
		fmt.Println()
	}

	fmt.Println()
}

func (b Board) copy() Board {
	var newBoard Board

	for _, row := range b.Rows {
		var newRow Line
		for _, c := range row {
			newRow = append(newRow, c)
		}
		newBoard.Rows = append(newBoard.Rows, newRow)
	}
	for _, col := range b.Columns {
		var newCol Line
		for _, c := range col {
			newCol = append(newCol, c)
		}
		newBoard.Columns = append(newBoard.Columns, newCol)
	}
	return newBoard
}

func (b Board) Solve() {
	var iter int
	logIter := func() { fmt.Println("Iteration", iter) }
	for {
		logIter()
		b.Print()
		start := b.copy()

		// check to see if puzzle is solved
		fmt.Println("check to see if puzzle is solved")

		if b.isSolved() {
			fmt.Println("SOLVED!")
			break
		}

		// check to see if each row is solvable
		for i, r := range b.Rows {
			var filled []int
			b.Rows[i], filled = tryFinish(r)
			for _, f := range filled {
				b.Columns[f][i] = b.Rows[i][f]
			}
			// apply rule of 3 to rows
			b.Rows[i], filled = ruleOf3(r)
			for _, f := range filled {
				b.Columns[f][i] = b.Rows[i][f]
			}

			// apply rule of same
			b.Rows[i], filled = ruleOfSame(i, b.Rows)
			for _, f := range filled {
				b.Columns[f][i] = b.Rows[i][f]
			}
		}

		// check to see if each column is solvable
		for i, r := range b.Columns {
			var filled []int
			b.Columns[i], filled = tryFinish(r)
			for _, f := range filled {
				b.Rows[f][i] = b.Columns[i][f]
			}
			// apply rule of 3 to columns
			b.Columns[i], filled = ruleOf3(r)
			for _, f := range filled {
				b.Rows[f][i] = b.Columns[i][f]
			}

			// apply rule of same
			b.Columns[i], filled = ruleOfSame(i, b.Columns)
			for _, f := range filled {
				b.Rows[f][i] = b.Columns[i][f]
			}
		}

		logIter()
		b.Print()

		if reflect.DeepEqual(start, b) {
			fmt.Println("STUCK!")
			break
		}
		iter++
	}
	logIter()
	b.Print()
}

func (b Board) isSolved() bool {
	solved := true

	for _, r := range b.Rows {
		solved = solved && r.isFinished()
		if !solved {
			return false
		}
	}

	for _, c := range b.Columns {
		solved = solved && c.isFinished()
		if !solved {
			fmt.Println("col not finished:", c)
			return false
		}
	}

	return solved
}
