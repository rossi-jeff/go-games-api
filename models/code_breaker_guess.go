package models

type CodeBreakerGuess struct {
	BaseModel
	CodeBreakerId int64                   `json:"code_breaker_id"`
	Colors        []CodeBreakerGuessColor `json:"colors,omitempty"`
	Keys          []CodeBreakerGuessKey   `json:"keys,omitempty"`
}

type CodeBreakerGuessJson struct {
	BaseModel
	CodeBreakerId int64                       `json:"code_breaker_id"`
	Colors        []CodeBreakerGuessColorJson `json:"colors,omitempty"`
	Keys          []CodeBreakerGuessKeyJson   `json:"keys,omitempty"`
}

func (c CodeBreakerGuess) Json() CodeBreakerGuessJson {
	result := CodeBreakerGuessJson{
		BaseModel:     c.BaseModel,
		CodeBreakerId: c.CodeBreakerId,
	}
	if len(c.Colors) > 0 {
		for i := 0; i < len(c.Colors); i++ {
			newColor := c.Colors[i].Json()
			result.Colors = append(result.Colors, newColor)
		}
	}
	if len(c.Keys) > 0 {
		for i := 0; i < len(c.Keys); i++ {
			newKey := c.Keys[i].Json()
			result.Keys = append(result.Keys, newKey)
		}
	}
	return result
}
