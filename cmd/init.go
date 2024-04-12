package cmd

type Globals struct {
	FileLoc string
	Jokes   []string
}

var globals Globals

func Init(jokes []string) {
	globals.Jokes = jokes
}

func InitFileLoc(fileLoc string) {
	globals.FileLoc = fileLoc
}
