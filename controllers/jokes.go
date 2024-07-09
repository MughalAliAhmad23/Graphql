package controllers

import (
	"Gqlgen/db"
	"Gqlgen/dbmodels"
	"Gqlgen/graph/model"
	"fmt"
	"strconv"
)

func CreateNewJoke(input model.JokeInput) *model.JokeResp {

	joke := dbmodels.Joke{
		Text: &input.Text,
	}

	db.DbConn.Create(&joke)

	newId := strconv.Itoa(int(joke.ID))
	jsonJoke := &model.JokeResp{
		ID:   &newId,
		Text: joke.Text,
	}

	return jsonJoke

}

func GetAllJoke() []*model.JokeResp {

	jokes := []dbmodels.Joke{}

	resp := db.DbConn.Find(&jokes)
	if resp.Error != nil {
		fmt.Println("Error:", resp.Error)
	}

	jsonJokes := []*model.JokeResp{}

	for _, dbJoke := range jokes {
		newId := strconv.Itoa(int(dbJoke.ID))
		joke := &model.JokeResp{
			ID:   &newId,
			Text: dbJoke.Text,
		}
		jsonJokes = append(jsonJokes, joke)
	}

	return jsonJokes

}

func GetAJoke(id string) *model.JokeResp {

	joke := &dbmodels.Joke{}

	resp := db.DbConn.Where("id=?", id).First(joke)
	if resp.Error != nil {
		fmt.Println("Error", resp.Error)
	}

	newId := strconv.Itoa(int(joke.ID))
	jsonJoke := model.JokeResp{
		ID:   &newId,
		Text: joke.Text,
	}

	return &jsonJoke
}
