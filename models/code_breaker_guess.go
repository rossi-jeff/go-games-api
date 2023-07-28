package models

type CodeBreakerGuess struct {
	BaseModel
	CodeBreakerId int64                   `json:"code_breaker_id"`
	Colors        []CodeBreakerGuessColor `json:"colors,omitempty"`
	Keys          []CodeBreakerGuessKey   `json:"keys,omitempty"`
}
