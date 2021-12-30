package httpserver

import (
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}

type SpyGame struct {
	StartedWith    int
	FinishedWith   string
	BlindAlert     []byte
	StartedCalled  bool
	FinishedCalled bool
}

func (g *SpyGame) Start(numberOfPlayers int, to io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartedCalled = true
	to.Write(g.BlindAlert)
}

func (g *SpyGame) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishedCalled = true
}

func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("scores are not equal, got %d, want %d", got, want)
	}
}

func AssertStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Code != want {
		t.Errorf("did not get correct status, got %d, want %d", got.Code, want)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func AssertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	headers := response.Result().Header
	if headers.Get("content-type") != want {
		t.Errorf("response did not have content-type of %q, got %v", want, headers)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}
	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q, want %q", store.winCalls[0], winner)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}
	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}

func AssertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.StartedCalled {
		t.Errorf("game should not have started")
	}
}

func AssertGameNotFinished(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game should not have finished")
	}
}

func AssertGameStartedWith(t testing.TB, game *SpyGame, startedWith int) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == startedWith
	})
	if !passed {
		t.Errorf("expected start called with %d but got %d", startedWith, game.StartedWith)
	}
}

func AssertGameFinishedWith(t testing.TB, game *SpyGame, finishedWith string) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == finishedWith
	})
	if !passed {
		t.Errorf("expected finish called with %q but got %q", finishedWith, game.FinishedWith)
	}
}

func AssertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func AssertWebsocketGotMsg(t testing.TB, ws *websocket.Conn, want string) {
	t.Helper()
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf("got %q, want %q", string(msg), want)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
