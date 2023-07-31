package payloads

type HangManCreatePayload struct {
	WordId int64
}

type HangManGuessPayload struct {
	Word   string
	Letter string
}

type HangManGuessResponse struct {
	Found  bool
	Letter string
}
