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
	dummyGame         = &SpyGame{}
	// dummyStdIn        = &bytes.Buffer{}
	dummyStdOut = &bytes.Buffer{}
)

type SpyGame struct {
	StartedWith   int
	FinishedWith  string
	StartedCalled bool
}

func (g *SpyGame) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartedCalled = true
}

func (g *SpyGame) Finish(winner string) {
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := userSends("2", "Chris wins")
		game := &SpyGame{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 2)
		assertGameFinishedWith(t, game, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := userSends("9", "Cleo wins")
		game := &SpyGame{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 9)
		assertGameFinishedWith(t, game, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := userSends("Pies")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("it prints an error when a 0 players is entered and does not start the game", func(t *testing.T) {
		in := userSends("0")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("it prints an error when a negative amount of players is entered and does not start the game", func(t *testing.T) {
		in := userSends("-1")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("it prints an error when a invalid winner is entered and does not start the game", func(t *testing.T) {
		in := userSends("2", "Lloyd is a killer")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
	t.Run("it prints an error when a correct winner is entered but forgot to type ' wins' and does not start the game", func(t *testing.T) {
		in := userSends("6", "Chris")
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func userSends(inputs ...string) io.Reader {
	return strings.NewReader(strings.Join(inputs, "\n"))
}

func assertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.StartedCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameStartedWith(t testing.TB, game *SpyGame, startedWith int) {
	t.Helper()
	if game.StartedWith != startedWith {
		t.Errorf("got %d players but expected want %d", game.StartedWith, startedWith)
	}
}

func assertGameFinishedWith(t testing.TB, game *SpyGame, finishedWith string) {
	t.Helper()
	if game.FinishedWith != finishedWith {
		t.Errorf("got %q winner but expected %q", game.FinishedWith, finishedWith)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
