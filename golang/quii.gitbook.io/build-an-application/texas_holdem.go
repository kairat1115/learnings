package httpserver

import (
	"io"
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexasHoldem(store PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

func (t *TexasHoldem) Start(numberOfPlayers int, alertDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	// blindIncrement := time.Duration(5+numberOfPlayers) * time.Second

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		t.alerter.ScheduleAlertAt(blindTime, blind, alertDestination)
		blindTime += blindIncrement
	}
}

func (t *TexasHoldem) Finish(winner string) {
	t.store.RecordWin(winner)
}
