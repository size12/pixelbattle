package entity

import (
	"errors"
	"strings"
)

type Color int

const (
	ColorWhite Color = iota
	ColorRed
	ColorGreen
	ColorBlue
	ColorBlack
	ColorYellow
)

func (c *Color) UnmarshalJSON(data []byte) error {
	color := strings.Replace(strings.ToLower(string(data)), "\"", "", -1)
	var result Color

	switch color {
	case "white":
		result = ColorWhite
	case "red":
		result = ColorRed
	case "green":
		result = ColorGreen
	case "blue":
		result = ColorBlue
	case "black":
		result = ColorBlack
	case "yellow":
		result = ColorYellow
	default:
		return errors.New("wrong color")
	}

	*c = result

	return nil
}

type Field [][]Color

type Dot struct {
	Color Color `json:"color"`
	X     int   `json:"x"`
	Y     int   `json:"y"`
}
