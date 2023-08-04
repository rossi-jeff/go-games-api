package models

import "go-games-api/enum"

type Klondike struct {
	BaseModel
	Status         enum.GameStatus
	Moves, Elapsed int
	UserId         NullInt64 `json:"user_id" swaggerType:"string"`
	User           User      `json:"user,omitempty"`
}

type KlondikePaginated struct {
	Items []Klondike
	Paginated
}

type KlondikeJson struct {
	BaseModel
	Moves, Elapsed int
	UserId         int64    `json:"user_id" swaggerType:"string"`
	User           UserJson `json:"user,omitempty"`
	Status         enum.GameStatusString
}

func (k Klondike) Json() KlondikeJson {
	return KlondikeJson{
		BaseModel: k.BaseModel,
		Moves:     k.Moves,
		Elapsed:   k.Elapsed,
		UserId:    NullInt64Value(k.UserId),
		User:      k.User.Json(),
		Status:    enum.GameStatusString(k.Status.String()),
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
