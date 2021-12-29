package httpserver

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
