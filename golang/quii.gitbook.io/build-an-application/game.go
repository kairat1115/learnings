package httpserver

import "io"

type Game interface {
	Start(numberOfPlayers int, alertDestination io.Writer)
	Finish(winner string)
}
