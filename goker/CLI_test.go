package goker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pmuens/goker/goker"
)

var dummyStdout = &bytes.Buffer{}
var dummyPlayerStore = &goker.StubPlayerStore{}
var dummyBlindAlerter = &goker.SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &goker.StubPlayerStore{}
		blindAlerter := &goker.SpyBlindAlerter{}

		cli := goker.NewCLI(in, dummyStdout, playerStore, blindAlerter)
		cli.PlayPoker()

		cases := []goker.ScheduledAlert{
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

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				goker.AssertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		in := strings.NewReader("7\n")
		out := &bytes.Buffer{}
		blindAlerter := &goker.SpyBlindAlerter{}

		cli := goker.NewCLI(in, out, dummyPlayerStore, blindAlerter)
		cli.PlayPoker()

		got := out.String()
		want := goker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []goker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				goker.AssertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		playerStore := &goker.StubPlayerStore{}

		cli := goker.NewCLI(in, dummyStdout, playerStore, dummyBlindAlerter)
		cli.PlayPoker()

		goker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from uer input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		playerStore := &goker.StubPlayerStore{}

		cli := goker.NewCLI(in, dummyStdout, playerStore, dummyBlindAlerter)
		cli.PlayPoker()

		goker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
