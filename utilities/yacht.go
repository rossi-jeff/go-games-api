package utilities

import (
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"sort"
)

func YachtCatagorySkip(id int64) []string {
	var skip []string
	turns := []models.YachtTurn{}
	initializers.DB.Where("yacht_id = ? and Category IS NOT NULL", id).Select("Category").Find(&turns)
	for i := 0; i < len(turns); i++ {
		category := turns[i].Category.String()
		skip = append(skip, category)
	}
	return skip
}

func scoreNumberYacht(dice []int, number int) int {
	score := 0
	count := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] == number {
			count++
		}
	}
	score = count * number
	return score
}

func scoreFullHouseYacht(dice []int) int {
	score := 0
	_, _, values := MapDieFaces(dice)
	if IntSliceIndexOf(3, values) != -1 && IntSliceIndexOf(2, values) != -1 {
		score = 25
	}
	return score
}

func scoreFourKindYacht(dice []int) int {
	score := 0
	dieMap, keys, _ := MapDieFaces(dice)
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if dieMap[key] >= 4 {
			score = 4 * key
		}
	}
	return score
}

func scoreBigStraightYacht(dice []int) int {
	score := 0
	_, keys, _ := MapDieFaces(dice)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	if IntSliceJoin(keys, ",") == "2,3,4,5,6" {
		score = 30
	}
	return score
}

func scoreLittleStraightYacht(dice []int) int {
	score := 0
	_, keys, _ := MapDieFaces(dice)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	if IntSliceJoin(keys, ",") == "1,2,3,4,5" {
		score = 30
	}
	return score
}

func scoreChoiceYacht(dice []int) int {
	score := 0
	for i := 0; i < len(dice); i++ {
		score = score + dice[i]
	}
	return score
}

func scoreYacht(dice []int) int {
	score := 0
	_, _, values := MapDieFaces(dice)
	if IntSliceIndexOf(5, values) != -1 {
		score = 50
	}
	return score
}

func ScoreYachtCategory(category string, dice []int) int {
	score := 0
	switch category {
	case "Ones":
		score = scoreNumberYacht(dice, 1)
	case "Twos":
		score = scoreNumberYacht(dice, 2)
	case "Threes":
		score = scoreNumberYacht(dice, 3)
	case "Fours":
		score = scoreNumberYacht(dice, 4)
	case "Fives":
		score = scoreNumberYacht(dice, 5)
	case "Sixes":
		score = scoreNumberYacht(dice, 6)
	case "FullHouse":
		score = scoreFullHouseYacht(dice)
	case "FourOfKind":
		score = scoreFourKindYacht(dice)
	case "BigStraight":
		score = scoreBigStraightYacht(dice)
	case "LittleStraight":
		score = scoreLittleStraightYacht(dice)
	case "Choice":
		score = scoreChoiceYacht(dice)
	case "Yacht":
		score = scoreYacht(dice)
	}
	return score
}

func YachtScoreOptions(dice []int, skip []string) []payloads.YachtScoreOption {
	var options []payloads.YachtScoreOption
	for i := 0; i < len(enum.YachtCategoryArray); i++ {
		category := enum.YachtCategoryArray[i]
		if StringSliceIndexOf(category, skip) == -1 {
			option := payloads.YachtScoreOption{
				Score:    ScoreYachtCategory(category, dice),
				Category: enum.YachtCategoryString(category),
			}
			options = append(options, option)
		}
	}
	return options
}
