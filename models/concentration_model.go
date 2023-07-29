package models

import "go-games-api/enum"

type Concentration struct {
	BaseModel
	Status                  enum.GameStatus
	Moves, Matched, Elapsed int
	UserId                  int64 `json:"user_id"`
	User                    User  `json:"user,omitempty"`
}

type ConcentrationPaginated struct {
	Items []Concentration
	Paginated
}

type ConcentrationJson struct {
	BaseModel
	Moves, Matched, Elapsed int
	UserId                  int64 `json:"user_id"`
	User                    User  `json:"user,omitempty"`
	Status                  string
}

func (c Concentration) Json() ConcentrationJson {
	return ConcentrationJson{
		BaseModel: c.BaseModel,
		Moves:     c.Moves,
		Matched:   c.Matched,
		Elapsed:   c.Elapsed,
		UserId:    c.UserId,
		User:      c.User,
		Status:    c.Status.String(),
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
