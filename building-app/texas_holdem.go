package poker

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
		store:   store,
		alerter: alerter,
	}
}

func (g *TexasHoldem) Start(players int, destination io.Writer) {
	g.scheduleBlindAlerts(players, destination)
}

func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}

func (g *TexasHoldem) scheduleBlindAlerts(numberOfPlayers int, at io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, at)
		blindTime = blindTime + blindIncrement
	}
}
