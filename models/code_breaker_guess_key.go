package models

import "go-games-api/enum"

type CodeBreakerGuessKey struct {
	BaseModel
	Key                enum.Key
	CodeBreakerGuessId int64 `json:"code_breaker_guess_id"`
}

type CodeBreakerGuessKeyJson struct {
	BaseModel
	CodeBreakerGuessId int64 `json:"code_breaker_guess_id"`
	Key                enum.KeyString
}

func (c CodeBreakerGuessKey) Json() CodeBreakerGuessKeyJson {
	return CodeBreakerGuessKeyJson{
		BaseModel:          c.BaseModel,
		CodeBreakerGuessId: c.CodeBreakerGuessId,
		Key:                enum.KeyString(c.Key.String()),
	}
}
