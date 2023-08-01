package enum

// Key enum info
// @Description database contains integer but values are "Black", "White"
type Key int

const (
	BLACK Key = iota
	WHITE
)

type KeyString string

const (
	K0 KeyString = "Black"
	K1 KeyString = "White"
)

var KeyArray = [2]string{"Black", "White"}

func (k Key) String() string {
	if k < 0 || k >= 2 {
		return ""
	}
	return KeyArray[k]
}

func (k Key) EnumIndex() int {
	return int(k)
}
