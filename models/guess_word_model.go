package models

import "go-games-api/enum"

type GuessWord struct {
	BaseModel
	Status  enum.GameStatus
	Score   int
	UserId  int64            `json:"user_id"`
	User    User             `json:"user,omitempty"`
	WordId  int64            `json:"word_id"`
	Word    Word             `json:"word,omitempty"`
	Guesses []GuessWordGuess `json:"guesses,omitempty"`
}

type GuessWordPaginated struct {
	Items []GuessWord
	Paginated
}

type GuessWordJson struct {
	BaseModel
	Score   int
	UserId  int64 `json:"user_id"`
	User    User  `json:"user,omitempty"`
	WordId  int64 `json:"word_id"`
	Word    Word  `json:"word,omitempty"`
	Status  string
	Guesses []GuessWordGuessJson `json:"guesses,omitempty"`
}

func (g GuessWord) Json() GuessWordJson {
	result := GuessWordJson{
		BaseModel: g.BaseModel,
		Score:     g.Score,
		UserId:    g.UserId,
		User:      g.User,
		WordId:    g.WordId,
		Word:      g.Word,
		Status:    g.Status.String(),
	}
	if len(g.Guesses) > 0 {
		for i := 0; i < len(g.Guesses); i++ {
			newGuess := g.Guesses[i].Json()
			result.Guesses = append(result.Guesses, newGuess)
		}
	}
	return result
}

type GuessWordPaginatedJson struct {
	Paginated
	Items []GuessWordJson
}

func (g GuessWordPaginated) Json() GuessWordPaginatedJson {
	result := GuessWordPaginatedJson{
		Paginated: g.Paginated,
	}
	if len(g.Items) > 0 {
		for i := 0; i < len(g.Items); i++ {
			newItem := g.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
