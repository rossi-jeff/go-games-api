package models

type GuessWordGuess struct {
	BaseModel
	Guess       string
	GuessWordId int64                  `json:"guess_word_id"`
	Ratings     []GuessWordGuessRating `json:"ratings,omitempty"`
}
