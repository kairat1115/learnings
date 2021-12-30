package httpserver_test

import (
	"bytes"
	"io"
	poker "learn-go-with-tests/build-an-application"
	"strings"
	"testing"
)

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdOut       = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := userSends("2", "Chris wins")
		game := &poker.SpyGame{}

		poker.NewCLI(in, dummyStdOut, game).PlayPoker()

		poker.AssertGameStartedWith(t, game, 2)
		poker.AssertGameFinishedWith(t, game, "Chris")
		poker.AssertMessagesSentToUser(t, dummyStdOut, poker.PlayerPrompt, poker.WinnerPrompt)
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := userSends("9", "Cleo wins")
		out := &bytes.Buffer{}
		game := &poker.SpyGame{}

		poker.NewCLI(in, out, game).PlayPoker()

		poker.AssertGameStartedWith(t, game, 9)
		poker.AssertGameFinishedWith(t, game, "Cleo")
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.WinnerPrompt)
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := userSends("Pies")
		out := &bytes.Buffer{}
		game := &poker.SpyGame{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertGameNotStarted(t, game)
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("it prints an error when a 0 players is entered and does not start the game", func(t *testing.T) {
		in := userSends("0")
		out := &bytes.Buffer{}
		game := &poker.SpyGame{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertGameNotStarted(t, game)
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("it prints an error when a negative amount of players is entered and does not start the game", func(t *testing.T) {
		in := userSends("-1")
		out := &bytes.Buffer{}
		game := &poker.SpyGame{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertGameNotStarted(t, game)
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("it prints an error when a invalid winner is entered and does not start the game", func(t *testing.T) {
		in := userSends("2", "Lloyd is a killer")
		out := &bytes.Buffer{}
		game := &poker.SpyGame{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertGameNotFinished(t, game)
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.WinnerPrompt, poker.BadWinnerInputErrMsg)
	})
	t.Run("it prints an error when a correct winner is entered but forgot to type ' wins' and does not start the game", func(t *testing.T) {
		in := userSends("6", "Chris")
		out := &bytes.Buffer{}
		game := &poker.SpyGame{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertGameNotFinished(t, game)
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.WinnerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func userSends(inputs ...string) io.Reader {
	return strings.NewReader(strings.Join(inputs, "\n"))
}
