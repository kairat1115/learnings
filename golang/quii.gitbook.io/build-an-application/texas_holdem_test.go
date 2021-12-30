package httpserver_test

import (
	"fmt"
	poker "learn-go-with-tests/build-an-application"
	"testing"
	"time"
)

func TestTexasHoldemGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)

		game.Start(5, dummyStdOut)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)

		game.Start(7, dummyStdOut)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestTexasHoldemGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(store, dummyBlindAlerter)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, alerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}
			got := alerter.Alerts[i]
			poker.AssertScheduledAlert(t, got, want)
		})
	}
}
