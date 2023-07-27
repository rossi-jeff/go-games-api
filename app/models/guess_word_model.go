package models

import "go-games-api/app/enum"

type GuessWord struct {
	BaseModel
	Status enum.GameStatus
	Score  int
	UserId int64 `json:"user_id"`
	User   User  `json:"user,omitempty"`
	WordId int64 `json:"word_id"`
	Word   Word  `json:"word,omitempty"`
}
