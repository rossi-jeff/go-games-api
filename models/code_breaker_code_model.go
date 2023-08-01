package models

import "go-games-api/enum"

type CodeBreakerCode struct {
	BaseModel
	Color         enum.Color
	CodeBreakerId int64 `json:"code_breaker_id"`
}

type CodeBreakerCodeJson struct {
	BaseModel
	CodeBreakerId int64 `json:"code_breaker_id"`
	Color         enum.ColorString
}

func (c CodeBreakerCode) Json() CodeBreakerCodeJson {
	return CodeBreakerCodeJson{
		BaseModel:     c.BaseModel,
		CodeBreakerId: c.CodeBreakerId,
		Color:         enum.ColorString(c.Color.String()),
	}
}
