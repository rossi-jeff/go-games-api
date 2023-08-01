package enum

// Color enum info
// @Description database contains integer but values are "Black", "Blue", "Brown", "Green", "Orange", "Purple", "Red", "White", "Yellow"
type Color int

const (
	Black Color = iota
	Blue
	Brown
	Green
	Orange
	Purple
	Red
	White
	Yellow
)

type ColorString string

const (
	C0 ColorString = "Black"
	C1 ColorString = "Blue"
	C2 ColorString = "Brown"
	C3 ColorString = "Green"
	C4 ColorString = "Orange"
	C5 ColorString = "Purple"
	C6 ColorString = "Red"
	C7 ColorString = "White"
	C8 ColorString = "Yellow"
)

var ColorArray = [...]string{"Black", "Blue", "Brown", "Green", "Orange", "Purple", "Red", "White", "Yellow"}

func (c Color) String() string {
	if c < 0 || c >= 9 {
		return ""
	}
	return ColorArray[c]
}

func ColorArrayIndex(color string) int {
	for i := 0; i < len(ColorArray); i++ {
		if color == ColorArray[i] {
			return i
		}
	}
	return -1
}
