package models

type Yacht struct {
	BaseModel
	Total    int
	NumTurns int
	UserId   NullInt64   `json:"user_id" swaggerType:"string"`
	User     User        `json:"user,omitempty"`
	Turns    []YachtTurn `json:"turns,omitempty"`
}

type YachtJson struct {
	BaseModel
	Total    int
	NumTurns int
	UserId   NullInt64       `json:"user_id" swaggerType:"string"`
	User     User            `json:"user,omitempty"`
	Turns    []YachtTurnJson `json:"turns,omitempty"`
}

func (y Yacht) Json() YachtJson {
	result := YachtJson{
		BaseModel: y.BaseModel,
		Total:     y.Total,
		NumTurns:  y.NumTurns,
		UserId:    y.UserId,
		User:      y.User,
	}
	if len(y.Turns) > 0 {
		for i := 0; i < len(y.Turns); i++ {
			newTurn := y.Turns[i].Json()
			result.Turns = append(result.Turns, newTurn)
		}
	}
	return result
}

type YachtPaginated struct {
	Paginated
	Items []Yacht
}

type YachtPaginatedJson struct {
	Paginated
	Items []YachtJson
}

func (y YachtPaginated) Json() YachtPaginatedJson {
	result := YachtPaginatedJson{
		Paginated: y.Paginated,
	}
	if len(y.Items) > 0 {
		for i := 0; i < len(y.Items); i++ {
			newItem := y.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
