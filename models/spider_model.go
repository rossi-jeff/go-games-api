package models

import "go-games-api/enum"

type Spider struct {
	BaseModel
	Status  enum.GameStatus
	Moves   int
	Elapsed int
	Suits   enum.Suit
	UserId  NullInt64 `json:"user_id" swaggerType:"string"`
	User    User      `json:"user,omitempty"`
}

type SpiderJson struct {
	BaseModel
	Moves   int
	Elapsed int
	Suits   enum.Suit
	UserId  NullInt64 `json:"user_id" swaggerType:"string"`
	User    User      `json:"user,omitempty"`
	Status  enum.GameStatusString
}

func (s Spider) Json() SpiderJson {
	return SpiderJson{
		BaseModel: s.BaseModel,
		Moves:     s.Moves,
		Elapsed:   s.Elapsed,
		Suits:     s.Suits,
		UserId:    s.UserId,
		User:      s.User,
		Status:    enum.GameStatusString(s.Status.String()),
	}
}

type SpiderPaginated struct {
	Paginated
	Items []Spider
}

type SpiderPaginatedJson struct {
	Paginated
	Items []SpiderJson
}

func (s SpiderPaginated) Json() SpiderPaginatedJson {
	result := SpiderPaginatedJson{
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
