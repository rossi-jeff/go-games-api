package models

type TenGrandTurn struct {
	BaseModel
	Score      int
	TenGrandId int64           `json:"ten_grand_id"`
	Scores     []TenGrandScore `json:"scores"`
}

type TenGrandTurnJson struct {
	BaseModel
	Score      int
	TenGrandId int64               `json:"ten_grand_id"`
	Scores     []TenGrandScoreJson `json:"scores"`
}

func (t TenGrandTurn) Json() TenGrandTurnJson {
	result := TenGrandTurnJson{
		BaseModel:  t.BaseModel,
		Score:      t.Score,
		TenGrandId: t.TenGrandId,
	}
	if len(t.Scores) > 0 {
		for i := 0; i < len(t.Scores); i++ {
			newScore := t.Scores[i].Json()
			result.Scores = append(result.Scores, newScore)
		}
	}
	return result
}
