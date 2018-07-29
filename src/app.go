package w3k

type Apper interface {
	// Config
	// Pipe []Piper{}
}

type App struct {
	Name   string
	Config map[string]interface{}
	Pipe   []Piper
}

type Piper interface {
	Execute() error
}
