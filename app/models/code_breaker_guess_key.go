package models

import "go-games-api/app/enum"

type CodeBreakerGuessKey struct {
	BaseModel
	Key                enum.Key
	CodeBreakerGuessId int64 `json:"code_breaker_guess_id"`
}
