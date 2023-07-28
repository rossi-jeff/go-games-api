package models

type TenGrandTurn struct {
	BaseModel
	Score      int
	TenGrandId int64           `json:"ten_grand_id"`
	Scores     []TenGrandScore `json:"scores"`
}
