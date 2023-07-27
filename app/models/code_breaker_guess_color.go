package models

import "go-games-api/app/enum"

type CodeBreakerGuessColor struct {
	BaseModel
	Color              enum.Color
	CodeBreakerGuessId int64 `json:"code_breaker_guess_id"`
}
