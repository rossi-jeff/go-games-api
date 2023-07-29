package models

import "go-games-api/enum"

type FreeCell struct {
	BaseModel
	Status         enum.GameStatus
	Moves, Elapsed int
	UserId         int64 `json:"user_id"`
	User           User  `json:"user,omitempty"`
}

type FreeCellPaginated struct {
	Items []FreeCell
	Paginated
}

type FreeCellJson struct {
	BaseModel
	Moves, Elapsed int
	UserId         int64 `json:"user_id"`
	User           User  `json:"user,omitempty"`
	Status         string
}

func (f FreeCell) Json() FreeCellJson {
	return FreeCellJson{
		BaseModel: f.BaseModel,
		Moves:     f.Moves,
		Elapsed:   f.Elapsed,
		UserId:    f.UserId,
		User:      f.User,
		Status:    f.Status.String(),
	}
}

type FreeCellPaginatedJson struct {
	Paginated
	Items []FreeCellJson
}

func (f FreeCellPaginated) Json() FreeCellPaginatedJson {
	result := FreeCellPaginatedJson{
		Paginated: f.Paginated,
	}
	if len(f.Items) > 0 {
		for i := 0; i < len(f.Items); i++ {
			newItem := f.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
