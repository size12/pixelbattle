package entity

type Color string

const (
	ColorWhite  Color = "white"
	ColorRed    Color = "red"
	ColorGreen  Color = "green"
	ColorBlue   Color = "blue"
	ColorBlack  Color = "black"
	ColorYellow Color = "yellow"
)

type Field [][]Color

type Dot struct {
	Color Color `json:"color"`
	X     int   `json:"x"`
	Y     int   `json:"y"`
}
