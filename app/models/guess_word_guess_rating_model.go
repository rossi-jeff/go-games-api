package models

import "go-games-api/app/enum"

type GuessWordGuessRating struct {
	BaseModel
	Rating           enum.Rating
	GuessWordGuessId int64 `json:"guess_word_guess_id"`
}
