package models

import "go-games-api/enum"

type PokerSquare struct {
	BaseModel
	Status enum.GameStatus
	Score  int
	UserId int64 `json:"user_id"`
	User   User  `json:"user,omitempty"`
}

type PokerSquarePaginated struct {
	Items []PokerSquare
	Paginated
}

type PokerSquareJson struct {
	BaseModel
	Score  int
	UserId int64 `json:"user_id"`
	User   User  `json:"user,omitempty"`
	Status string
}

func (p PokerSquare) Json() PokerSquareJson {
	return PokerSquareJson{
		BaseModel: p.BaseModel,
		Score:     p.Score,
		UserId:    p.UserId,
		User:      p.User,
		Status:    p.Status.String(),
	}
}

type PokerSquarePaginatedJson struct {
	Paginated
	Items []PokerSquareJson
}

func (p PokerSquarePaginated) Json() PokerSquarePaginatedJson {
	result := PokerSquarePaginatedJson{
		Paginated: p.Paginated,
	}
	if len(p.Items) > 0 {
		for i := 0; i < len(p.Items); i++ {
			newItem := p.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
