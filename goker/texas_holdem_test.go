package goker_test

import (
	"io"
	"testing"
	"time"

	"github.com/pmuens/goker/goker"
)

func TestTexasHoldemStart(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &goker.SpyBlindAlerter{}
		game := goker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5, io.Discard)

		alerts := []goker.ScheduledAlert{
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

		goker.CheckScheduledAlerts(t, alerts, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &goker.SpyBlindAlerter{}
		game := goker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7, io.Discard)

		alerts := []goker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		goker.CheckScheduledAlerts(t, alerts, blindAlerter)
	})
}

func TestTexasHoldemFinish(t *testing.T) {
	store := &goker.StubPlayerStore{}
	game := goker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	goker.AssertPlayerWin(t, store, winner)
}
