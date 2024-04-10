package cmd

type Globals struct {
	Jokes []string
}

var globals Globals

func Init(jokes []string) {
	globals.Jokes = jokes
}
