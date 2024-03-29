package models

import "go-games-api/enum"

type GuessWordGuessRating struct {
	BaseModel
	Rating           enum.Rating
	GuessWordGuessId int64 `json:"guess_word_guess_id"`
}

type GuessWordGuessRatingJson struct {
	BaseModel
	GuessWordGuessId int64 `json:"guess_word_guess_id"`
	Rating           enum.RatingString
}

func (g GuessWordGuessRating) Json() GuessWordGuessRatingJson {
	return GuessWordGuessRatingJson{
		BaseModel:        g.BaseModel,
		GuessWordGuessId: g.GuessWordGuessId,
		Rating:           enum.RatingString(g.Rating.String()),
	}
}
