package httpserver

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt         = "Please enter the number of players: "
	WinnerPrompt         = "Please enter the winner in format '{player} wins': "
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
	BadWinnerInputErrMsg = "bad winner received for player wins, please try again with '{player} wins'"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayersInput := cli.readline()
	numberOfPlayers, err := strconv.Atoi(numberOfPlayersInput)
	if err != nil || numberOfPlayers < 1 {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers, cli.out)

	fmt.Fprint(cli.out, WinnerPrompt)
	winnerInput := cli.readline()
	winner, err := extractWinner(winnerInput)
	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputErrMsg)
		return
	}

	cli.game.Finish(winner)
}

func (cli *CLI) readline() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) (string, error) {
	if !strings.HasSuffix(userInput, " wins") {
		return "", errors.New(BadWinnerInputErrMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}
