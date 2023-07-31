package payloads

type GuesssWordCreatePayload struct {
	WordId int64
}

type GuessWordGuessPayload struct {
	Word  string
	Guess string
}

type GuessWordHintsPayload struct {
	Length int
	Green  []string
	Brown  [][]string
	Gray   []string
}
