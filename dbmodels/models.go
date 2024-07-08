package dbmodels

type Joke struct {
	Id   *int    `db:"id"`
	Text *string `db:"text"`
}

type JokeInput struct {
	Text string `db:"text"`
}
