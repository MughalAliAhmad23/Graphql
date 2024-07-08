package controllers

import (
	"Gqlgen/db"
	"Gqlgen/dbmodels"
	"Gqlgen/graph/model"
	"strconv"
)

func CreateNewJoke(input model.JokeInput) (*model.JokeResp, error) {

	var joke dbmodels.Joke

	query := "INSERT INTO myjoke (joke) VALUES ($1) RETURNING id,joke"
	err := db.DbConn.QueryRow(query, input.Text).Scan(&joke.Id, &joke.Text)
	if err != nil {
		return nil, err
	}
	var jsonJoke model.JokeResp
	convertedId := strconv.Itoa(*joke.Id)
	jsonJoke.ID = &convertedId
	jsonJoke.Text = joke.Text
	return &jsonJoke, err

}

func GetAllJoke() ([]*model.JokeResp, error) {

	query := "SELECT * FROM myjoke"
	rows, err := db.DbConn.Query(query)
	if err != nil {
		return nil, err
	}

	var jokes []*model.JokeResp

	for rows.Next() {
		var joke dbmodels.Joke
		err := rows.Scan(&joke.Id, &joke.Text)
		if err != nil {
			return nil, err
		}
		var jsonJoke model.JokeResp
		convertedId := strconv.Itoa(*joke.Id)
		jsonJoke.ID = &convertedId
		jsonJoke.Text = joke.Text
		jokes = append(jokes, &jsonJoke)
	}
	return jokes, nil
}

func GetAJoke(id string) (*model.JokeResp, error) {

	var joke dbmodels.Joke

	query := ("SELECT * FROM myjoke WHERE id = $1")
	err := db.DbConn.QueryRow(query, id).Scan(&joke.Id, &joke.Text)
	if err != nil {
		return nil, err
	}
	var jsonJoke model.JokeResp
	convertedId := strconv.Itoa(*joke.Id)
	jsonJoke.ID = &convertedId
	jsonJoke.Text = joke.Text

	return &jsonJoke, nil
}
