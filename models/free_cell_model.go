package models

import "go-games-api/enum"

type FreeCell struct {
	BaseModel
	Status  enum.GameStatus `gorm:"column:Status"`
	Moves   int             `gorm:"column:Moves"`
	Elapsed int             `gorm:"column:Elapsed"`
	UserId  NullInt64       `json:"user_id" swaggerType:"string"`
	User    User            `json:"user,omitempty"`
}

type FreeCellPaginated struct {
	Items []FreeCell
	Paginated
}

type FreeCellJson struct {
	BaseModel
	Moves   int   `gorm:"column:Moves"`
	Elapsed int   `gorm:"column:Elapsed"`
	UserId  int64 `json:"user_id" swaggerType:"string"`
	User    User  `json:"user,omitempty"`
	Status  enum.GameStatusString
}

func (f FreeCell) Json() FreeCellJson {
	return FreeCellJson{
		BaseModel: f.BaseModel,
		Moves:     f.Moves,
		Elapsed:   f.Elapsed,
		UserId:    NullInt64Value(f.UserId),
		User:      f.User,
		Status:    enum.GameStatusString(f.Status.String()),
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
