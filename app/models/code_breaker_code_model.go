package models

import "go-games-api/app/enum"

type CodeBreakerCode struct {
	BaseModel
	Color         enum.Color
	CodeBreakerId int64 `json:"code_breaker_id"`
}
