package models

import "go-games-api/enum"

type HangMan struct {
	BaseModel
	Correct, Wrong string
	Status         enum.GameStatus
	Score          int
	UserId         NullInt64 `json:"user_id" swaggerType:"string"`
	User           User      `json:"user,omitempty"`
	WordId         NullInt64 `json:"word_id" swaggerType:"string"`
	Word           Word      `json:"word,omitempty"`
}

type HangManPaginated struct {
	Items []HangMan
	Paginated
}

type HangManJson struct {
	BaseModel
	Correct, Wrong string
	Score          int
	UserId         NullInt64 `json:"user_id" swaggerType:"string"`
	User           User      `json:"user,omitempty"`
	WordId         NullInt64 `json:"word_id" swaggerType:"string"`
	Word           Word      `json:"word,omitempty"`
	Status         enum.GameStatusString
}

func (h HangMan) Json() HangManJson {
	return HangManJson{
		BaseModel: h.BaseModel,
		Correct:   h.Correct,
		Wrong:     h.Wrong,
		Score:     h.Score,
		UserId:    h.UserId,
		User:      h.User,
		WordId:    h.WordId,
		Word:      h.Word,
		Status:    enum.GameStatusString(h.Status.String()),
	}
}

type HangManPaginatedJson struct {
	Paginated
	Items []HangManJson
}

func (h HangManPaginated) Json() HangManPaginatedJson {
	result := HangManPaginatedJson{
		Paginated: h.Paginated,
	}
	if len(h.Items) > 0 {
		for i := 0; i < len(h.Items); i++ {
			newItem := h.Items[i].Json()
			result.Items = append(result.Items, newItem)
		}
	}
	return result
}
