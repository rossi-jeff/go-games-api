package models

import "go-games-api/enum"

type TenGrand struct {
	BaseModel
	Status enum.GameStatus
	Score  int
	UserId NullInt64      `json:"user_id" swaggerType:"string"`
	User   User           `json:"user,omitempty"`
	Turns  []TenGrandTurn `json:"turns,omitempty"`
}

type TenGrandJson struct {
	BaseModel
	Score  int
	UserId int64 `json:"user_id" swaggerType:"string"`
	User   User  `json:"user,omitempty"`
	Status enum.GameStatusString
	Turns  []TenGrandTurnJson `json:"turns,omitempty"`
}

func (t TenGrand) Json() TenGrandJson {
	result := TenGrandJson{
		BaseModel: t.BaseModel,
		Score:     t.Score,
		UserId:    NullInt64Value(t.UserId),
		User:      t.User,
		Status:    enum.GameStatusString(t.Status.String()),
	}
	if len(t.Turns) > 0 {
		for i := 0; i < len(t.Turns); i++ {
			newTurn := t.Turns[i].Json()
			result.Turns = append(result.Turns, newTurn)
		}
	}
	return result
}

type TenGrandPaginated struct {
	Paginated
	Items []TenGrand
}

type TenGrandPaginatedJson struct {
	Paginated
	Items []TenGrandJson
}

func (t TenGrandPaginated) Json() TenGrandPaginatedJson {
	result := TenGrandPaginatedJson{
		Paginated: t.Paginated,
	}
	if len(t.Items) > 0 {
		for i := 0; i < len(t.Items); i++ {
			newItem := t.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
