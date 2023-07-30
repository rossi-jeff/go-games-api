package models

type GuessWordGuess struct {
	BaseModel
	Guess       string
	GuessWordId int64                  `json:"guess_word_id"`
	Ratings     []GuessWordGuessRating `json:"ratings,omitempty"`
}

type GuessWordGuessJson struct {
	BaseModel
	Guess       string
	GuessWordId int64                      `json:"guess_word_id"`
	Ratings     []GuessWordGuessRatingJson `json:"ratings,omitempty"`
}

func (g GuessWordGuess) Json() GuessWordGuessJson {
	result := GuessWordGuessJson{
		BaseModel:   g.BaseModel,
		Guess:       g.Guess,
		GuessWordId: g.GuessWordId,
	}
	if len(g.Ratings) > 0 {
		for i := 0; i < len(g.Ratings); i++ {
			newRating := g.Ratings[i].Json()
			result.Ratings = append(result.Ratings, newRating)
		}
	}
	return result
}
