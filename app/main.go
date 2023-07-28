package main

import (
	"encoding/json"
	"fmt"
	"go-games-api/app/enum"
	"go-games-api/app/models"
	"log"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func main() {
	user := models.User{}
	res1, err := PrettyStruct(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res1)
	word := models.Word{}
	res2, err := PrettyStruct(word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res2)
	codeBreaker := models.CodeBreaker{
		Status: enum.Playing,
	}
	res3, err := PrettyStruct(codeBreaker)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res3)
}
