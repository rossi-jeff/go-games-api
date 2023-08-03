package utilities

import (
	"go-games-api/enum"
	"go-games-api/payloads"
	"sort"
)

func MapDieFaces(dice []int) (map[int]int, []int, []int) {
	dieMap := map[int]int{}
	var values []int
	var keys []int
	for i := 0; i < len(dice); i++ {
		face := dice[i]
		_, exists := dieMap[face]
		if !exists {
			dieMap[face] = 0
		}
		dieMap[face]++
	}
	for k, v := range dieMap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return dieMap, keys, values
}

func scoreOnesTG(dice []int) int {
	score := 0
	count := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] == 1 {
			count++
		}
	}
	if count > 0 {
		score = count * 100
	}
	return score
}

func scoreFivesTG(dice []int) int {
	score := 0
	count := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] == 5 {
			count++
		}
	}
	if count > 0 {
		score = count * 50
	}
	return score
}

func scoreFullHouseTG(dice []int) int {
	score := 0
	if len(dice) < 5 {
		return score
	}
	_, _, values := MapDieFaces(dice)
	if IntSliceIndexOf(3, values) != -1 && IntSliceIndexOf(2, values) != -1 {
		score = 1500
	}
	return score
}

func scoreStraightTG(dice []int) int {
	score := 0
	if len(dice) < 6 {
		return score
	}
	_, keys, _ := MapDieFaces(dice)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	if IntSliceJoin(keys, ",") == "1,2,3,4,5,6" {
		score = 2000
	}
	return score
}

func scoreThreePairTG(dice []int) int {
	score := 0
	if len(dice) < 6 {
		return score
	}
	_, _, values := MapDieFaces(dice)
	if IntSliceJoin(values, ",") == "2,2,2" {
		score = 1500
	}
	return score
}

func scoreDoubleThreeKindTG(dice []int) int {
	score := 0
	if len(dice) < 6 {
		return score
	}
	_, keys, values := MapDieFaces(dice)
	if IntSliceJoin(values, ",") == "3,3" {
		for i := 0; i < len(keys); i++ {
			key := keys[i]
			if key == 1 {
				score = score + 1000
			} else {
				score = score + (key * 100)
			}
		}
	}
	return score
}

func scoreThreeKindTG(dice []int) int {
	score := 0
	if len(dice) < 3 {
		return score
	}
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 3 {
			if key == 1 {
				score = score + 1000
			} else {
				score = score + (key * 100)
			}
		}
	}
	return score
}

func scoreFourKindTG(dice []int) int {
	score := 0
	if len(dice) < 4 {
		return score
	}
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 4 {
			if key == 1 {
				score = score + 2000
			} else {
				score = score + (key * 200)
			}
		}
	}
	return score
}

func scoreFiveKindTG(dice []int) int {
	score := 0
	if len(dice) < 5 {
		return score
	}
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 5 {
			if key == 1 {
				score = score + 4000
			} else {
				score = score + (key * 400)
			}
		}
	}
	return score
}

func scoreSixKindTG(dice []int) int {
	score := 0
	if len(dice) < 6 {
		return score
	}
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 6 {
			if key == 1 {
				score = score + 8000
			} else {
				score = score + (key * 800)
			}
		}
	}
	return score
}

func diceOnes(dice []int) []int {
	return diceMatchFace(dice, 1)
}

func diceFives(dice []int) []int {
	return diceMatchFace(dice, 5)
}

func diceThreePairs(dice []int) []int {
	var used []int
	_, _, values := MapDieFaces(dice)
	if IntSliceJoin(values, ",") == "2,2,2" {
		used = append(used, dice...)
	}
	return used
}

func diceStraight(dice []int) []int {
	var used []int
	_, keys, _ := MapDieFaces(dice)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	if IntSliceJoin(keys, ",") == "1,2,3,4,5,6" {
		used = append(used, dice...)
	}
	return used
}

func diceFullHouse(dice []int) []int {
	var used []int
	dieMap, keys, values := MapDieFaces(dice)
	if IntSliceIndexOf(3, values) != -1 && IntSliceIndexOf(2, values) != -1 {
		for i := 0; i < len(keys); i++ {
			key := keys[i]
			if dieMap[key] == 2 || dieMap[key] == 3 {
				used = append(used, diceMatchFace(dice, key)...)
			}
		}
	}
	return used
}

func diceDoubleThreeKind(dice []int) []int {
	var used []int
	_, _, values := MapDieFaces(dice)
	if IntSliceJoin(values, ",") == "3,3" {
		used = append(used, dice...)
	}
	return used
}

func diceThreeKind(dice []int) []int {
	var used []int
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 3 {
			used = append(used, diceMatchFace(dice, key)...)
		}
	}
	return used
}

func diceFourKind(dice []int) []int {
	var used []int
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 4 {
			used = append(used, diceMatchFace(dice, key)...)
		}
	}
	return used
}

func diceFiveKind(dice []int) []int {
	var used []int
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 5 {
			used = append(used, diceMatchFace(dice, key)...)
		}
	}
	return used
}

func diceSixKind(dice []int) []int {
	var used []int
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] == 6 {
			used = append(used, diceMatchFace(dice, key)...)
		}
	}
	return used
}

func diceMatchFace(dice []int, face int) []int {
	var used []int
	for i := 0; i < len(dice); i++ {
		d := dice[i]
		if d == face {
			used = append(used, d)
		}
	}
	return used
}

func RemoveUsedDice(dice []int, used []int) []int {
	for i := 0; i < len(used); i++ {
		face := used[i]
		idx := IntSliceIndexOf(face, dice)
		if idx != -1 {
			dice = DeleteIntSliceIndex(dice, idx)
		}
	}
	return dice
}

func TenGrandDiceScoreOptions(dice []int) []payloads.TenGrandOption {
	var options []payloads.TenGrandOption
	for i := 0; i < len(enum.TenGrandCategoryArray); i++ {
		option := payloads.TenGrandOption{
			Category: enum.TenGrandCategoryString(enum.TenGrandCategoryArray[i]),
		}
		switch enum.TenGrandCategoryArray[i] {
		case "CrapOut":
			option.Score = 0
		case "Ones":
			option.Score = scoreOnesTG(dice)
		case "Fives":
			option.Score = scoreFivesTG(dice)
		case "ThreePairs":
			option.Score = scoreThreePairTG(dice)
		case "Straight":
			option.Score = scoreStraightTG(dice)
		case "FullHouse":
			option.Score = scoreFullHouseTG(dice)
		case "DoubleThreeKind":
			option.Score = scoreDoubleThreeKindTG(dice)
		case "ThreeKind":
			option.Score = scoreThreeKindTG(dice)
		case "FourKind":
			option.Score = scoreFourKindTG(dice)
		case "FiveKind":
			option.Score = scoreFiveKindTG(dice)
		case "SixKind":
			option.Score = scoreSixKindTG(dice)
		}
		if option.Score > 0 || option.Category == enum.TG0 {
			options = append(options, option)
		}
	}
	sort.Slice(options, func(i, j int) bool {
		return options[j].Score < options[i].Score
	})
	return options
}

func CategoryScoreAndDice(category enum.TenGrandCategoryString, dice []int) (int, []int) {
	score := 0
	var used []int
	switch string(category) {
	case "CrapOut":
		score = 0
		used = append(used, dice...)
	case "Ones":
		score = scoreOnesTG(dice)
		used = diceOnes(dice)
	case "Fives":
		score = scoreFivesTG(dice)
		used = diceFives(dice)
	case "ThreePairs":
		score = scoreThreePairTG(dice)
		used = diceThreePairs(dice)
	case "Straight":
		score = scoreStraightTG(dice)
		used = diceStraight(dice)
	case "FullHouse":
		score = scoreFullHouseTG(dice)
		used = diceFullHouse(dice)
	case "DoubleThreeKind":
		score = scoreDoubleThreeKindTG(dice)
		used = diceDoubleThreeKind(dice)
	case "ThreeKind":
		score = scoreThreeKindTG(dice)
		used = diceThreeKind(dice)
	case "FourKind":
		score = scoreFourKindTG(dice)
		used = diceFourKind(dice)
	case "FiveKind":
		score = scoreFiveKindTG(dice)
		used = diceFiveKind(dice)
	case "SixKind":
		score = scoreSixKindTG(dice)
		used = diceSixKind(dice)
	}
	return score, used
}
