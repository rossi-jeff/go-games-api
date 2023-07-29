package enum

type Key int

const (
	BLACK Key = iota
	WHITE
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
