package types

type DadJoke struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int32  `json:"status"`
}

type JokesJsonFile struct {
	Data []string `json:"data"`
}
