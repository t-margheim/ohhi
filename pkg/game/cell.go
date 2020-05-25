package game

import (
	"fmt"

	"github.com/fatih/color"
)

type Cell int

func (c Cell) Print() {
	switch c {
	case Empty:
		fmt.Print(color.BlackString("[■]"))
	case Red:
		fmt.Print(color.RedString("[■]"))
	case Blue:
		fmt.Print(color.BlueString("[■]"))
	default:
		fmt.Println("ERROR: c is:", c)
	}
}

const (
	Empty Cell = iota
	Red
	Blue
)

func otherCell(c Cell) Cell {
	switch c {
	case Red:
		return Blue
	case Blue:
		return Red
	default:
		return Empty
	}
}
