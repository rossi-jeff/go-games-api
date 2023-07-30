package models

import "go-games-api/enum"

type SeaBattle struct {
	BaseModel
	Status      enum.GameStatus
	Score, Axis int
	UserId      int64           `json:"user_id"`
	User        User            `json:"user,omitempty"`
	Ships       []SeaBattleShip `json:"ships"`
	Turns       []SeaBattleTurn `json:"turns"`
}

type SeaBattleJson struct {
	BaseModel
	Score, Axis int
	UserId      int64 `json:"user_id"`
	User        User  `json:"user,omitempty"`
	Status      string
	Ships       []SeaBattleShipJson `json:"ships"`
	Turns       []SeaBattleTurnJson `json:"turns"`
}

func (s SeaBattle) Json() SeaBattleJson {
	result := SeaBattleJson{
		BaseModel: s.BaseModel,
		Score:     s.Score,
		Axis:      s.Axis,
		UserId:    s.UserId,
		User:      s.User,
		Status:    s.Status.String(),
	}
	if len(s.Ships) > 0 {
		for i := 0; i < len(s.Ships); i++ {
			newShip := s.Ships[i].Json()
			result.Ships = append(result.Ships, newShip)
		}
	}
	if len(s.Turns) > 0 {
		for i := 0; i < len(s.Turns); i++ {
			newTurn := s.Turns[i].Json()
			result.Turns = append(result.Turns, newTurn)
		}
	}
	return result
}

type SeaBattlePaginated struct {
	Paginated
	Items []SeaBattle
}

type SeaBattlePaginatedJson struct {
	Paginated
	Items []SeaBattleJson
}

func (s SeaBattlePaginated) Json() SeaBattlePaginatedJson {
	result := SeaBattlePaginatedJson{
		Paginated: s.Paginated,
	}
	if len(s.Items) > 0 {
		for i := 0; i < len(s.Items); i++ {
			newItem := s.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
