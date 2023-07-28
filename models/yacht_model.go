package models

type Yacht struct {
	BaseModel
	Total, NumTurns int
	UserId          int64       `json:"user_id"`
	User            User        `json:"user,omitempty"`
	Turns           []YachtTurn `json:"turns,omitempty"`
}
