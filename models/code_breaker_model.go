package models

import "go-games-api/enum"

type CodeBreaker struct {
	BaseModel
	Status                 enum.GameStatus
	Columns, Colors, Score int
	Available              string
	UserId                 NullInt64          `json:"user_id" swaggerType:"string"`
	User                   User               `json:"user,omitempty"`
	Codes                  []CodeBreakerCode  `json:"codes,omitempty"`
	Guesses                []CodeBreakerGuess `json:"guesses,omitempty"`
}

type CodeBreakerPaginated struct {
	Items []CodeBreaker
	Paginated
}

type CodeBreakerJson struct {
	BaseModel
	Columns, Colors, Score int
	Available              string
	UserId                 NullInt64 `json:"user_id" swaggerType:"string"`
	User                   User      `json:"user,omitempty"`
	Status                 string
	Codes                  []CodeBreakerCodeJson  `json:"codes,omitempty"`
	Guesses                []CodeBreakerGuessJson `json:"guesses,omitempty"`
}

func (c CodeBreaker) Json() CodeBreakerJson {
	result := CodeBreakerJson{
		BaseModel: c.BaseModel,
		Columns:   c.Columns,
		Colors:    c.Colors,
		Score:     c.Score,
		Available: c.Available,
		UserId:    c.UserId,
		User:      c.User,
		Status:    c.Status.String(),
	}
	if len(c.Codes) > 0 {
		for i := 0; i < len(c.Codes); i++ {
			newCode := c.Codes[i].Json()
			result.Codes = append(result.Codes, newCode)
		}
	}
	if len(c.Guesses) > 0 {
		for i := 0; i < len(c.Guesses); i++ {
			newGuess := c.Guesses[i].Json()
			result.Guesses = append(result.Guesses, newGuess)
		}
	}
	return result
}

type CodeBreakerPaginatedJson struct {
	Paginated
	Items []CodeBreakerJson
}

func (c CodeBreakerPaginated) Json() CodeBreakerPaginatedJson {
	result := CodeBreakerPaginatedJson{
		Paginated: c.Paginated,
	}
	if len(c.Items) > 0 {
		for i := 0; i < len(c.Items); i++ {
			newItem := c.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
