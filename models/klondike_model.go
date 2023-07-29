package models

import "go-games-api/enum"

type Klondike struct {
	BaseModel
	Status         enum.GameStatus
	Moves, Elapsed int
	UserId         int64 `json:"user_id"`
	User           User  `json:"user,omitempty"`
}

type KlondikePaginated struct {
	Items []Klondike
	Paginated
}

type KlondikeJson struct {
	BaseModel
	Moves, Elapsed int
	UserId         int64 `json:"user_id"`
	User           User  `json:"user,omitempty"`
	Status         string
}

func (k Klondike) Json() KlondikeJson {
	return KlondikeJson{
		BaseModel: k.BaseModel,
		Moves:     k.Moves,
		Elapsed:   k.Elapsed,
		UserId:    k.UserId,
		User:      k.User,
		Status:    k.Status.String(),
	}
}

type KlondikePaginatedJson struct {
	Paginated
	Items []KlondikeJson
}

func (k KlondikePaginated) Json() KlondikePaginatedJson {
	result := KlondikePaginatedJson{
		Paginated: k.Paginated,
	}
	if len(k.Items) > 0 {
		for i := 0; i < len(k.Items); i++ {
			newItem := k.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
