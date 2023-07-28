package models

import "go-games-api/enum"

type CodeBreaker struct {
	BaseModel
	Status                 enum.GameStatus
	Columns, Colors, Score int
	Available              string
	UserId                 int64              `json:"user_id"`
	User                   User               `json:"user,omitempty"`
	Codes                  []CodeBreakerCode  `json:"codes,omitempty"`
	Guesses                []CodeBreakerGuess `json:"guesses,omitempty"`
}
