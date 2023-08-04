package models

import "go-games-api/enum"

type PokerSquare struct {
	BaseModel
	Status enum.GameStatus
	Score  int
	UserId NullInt64 `json:"user_id" swaggerType:"string"`
	User   User      `json:"user,omitempty"`
}

type PokerSquarePaginated struct {
	Items []PokerSquare
	Paginated
}

type PokerSquareJson struct {
	BaseModel
	Score  int
	UserId int64    `json:"user_id" swaggerType:"string"`
	User   UserJson `json:"user,omitempty"`
	Status enum.GameStatusString
}

func (p PokerSquare) Json() PokerSquareJson {
	return PokerSquareJson{
		BaseModel: p.BaseModel,
		Score:     p.Score,
		UserId:    NullInt64Value(p.UserId),
		User:      p.User.Json(),
		Status:    enum.GameStatusString(p.Status.String()),
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
