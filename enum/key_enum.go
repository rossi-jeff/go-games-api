package enum

type Key int

const (
	BLACK Key = iota
	WHITE
)

var KeyArray = [2]string{"Black", "White"}

func (k Key) String() string {
	return KeyArray[k]
}

func (k Key) EnumIndex() int {
	return int(k)
}
