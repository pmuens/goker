package goker_test

import (
	"bytes"
	"testing"

	"github.com/pmuens/goker/goker"
)

var dummyStdout = &bytes.Buffer{}
var dummyPlayerStore = &goker.StubPlayerStore{}
var dummyBlindAlerter = &goker.SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &goker.GameSpy{}
		out := &bytes.Buffer{}

		in := goker.UserSends("3", "Chris wins")
		cli := goker.NewCLI(in, out, game)

		cli.PlayPoker()

		goker.AssertMessagesSentToUser(t, out, goker.PlayerPrompt)
		goker.AssertGameStartedWith(t, game, 3)
		goker.AssertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &goker.GameSpy{}

		in := goker.UserSends("8", "Cleo wins")
		cli := goker.NewCLI(in, dummyStdout, game)

		cli.PlayPoker()

		goker.AssertGameStartedWith(t, game, 8)
		goker.AssertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non-numeric value is entered and does not start the game", func(t *testing.T) {
		game := &goker.GameSpy{}
		out := &bytes.Buffer{}

		in := goker.UserSends("pies")
		cli := goker.NewCLI(in, out, game)

		cli.PlayPoker()

		goker.AssertGameNotStarted(t, game)
		goker.AssertMessagesSentToUser(t, out, goker.PlayerPrompt, goker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &goker.GameSpy{}
		out := &bytes.Buffer{}

		in := goker.UserSends("8", "Lloyd is the winner of the game")
		cli := goker.NewCLI(in, out, game)

		cli.PlayPoker()

		goker.AssertGameNotFinished(t, game)
		goker.AssertMessagesSentToUser(t, out, goker.PlayerPrompt, goker.BadWinnerInputMsg)
	})
}
