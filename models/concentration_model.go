package models

import "go-games-api/enum"

type Concentration struct {
	BaseModel
	Status  enum.GameStatus `gorm:"column:Status"`
	Moves   int             `gorm:"column:Moves"`
	Matched int             `gorm:"column:Matched"`
	Elapsed int             `gorm:"column:Elapsed"`
	UserId  NullInt64       `json:"user_id" swaggerType:"string"`
	User    User            `json:"user,omitempty"`
}

type ConcentrationPaginated struct {
	Items []Concentration
	Paginated
}

type ConcentrationJson struct {
	BaseModel
	Moves   int       `gorm:"column:Moves"`
	Matched int       `gorm:"column:Matched"`
	Elapsed int       `gorm:"column:Elapsed"`
	UserId  NullInt64 `json:"user_id" swaggerType:"string"`
	User    User      `json:"user,omitempty"`
	Status  enum.GameStatusString
}

func (c Concentration) Json() ConcentrationJson {
	return ConcentrationJson{
		BaseModel: c.BaseModel,
		Moves:     c.Moves,
		Matched:   c.Matched,
		Elapsed:   c.Elapsed,
		UserId:    c.UserId,
		User:      c.User,
		Status:    enum.GameStatusString(c.Status.String()),
	}
}

type ConcentrationPaginatedJson struct {
	Paginated
	Items []ConcentrationJson
}

func (c ConcentrationPaginated) Json() ConcentrationPaginatedJson {
	result := ConcentrationPaginatedJson{
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
