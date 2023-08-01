package models

import "go-games-api/enum"

type CodeBreakerGuessColor struct {
	BaseModel
	Color              enum.Color
	CodeBreakerGuessId int64 `json:"code_breaker_guess_id"`
}

type CodeBreakerGuessColorJson struct {
	BaseModel
	CodeBreakerGuessId int64 `json:"code_breaker_guess_id"`
	Color              enum.ColorString
}

func (c CodeBreakerGuessColor) Json() CodeBreakerGuessColorJson {
	return CodeBreakerGuessColorJson{
		BaseModel:          c.BaseModel,
		CodeBreakerGuessId: c.CodeBreakerGuessId,
		Color:              enum.ColorString(c.Color.String()),
	}
}
