package goker_test

import (
	"strings"
	"testing"

	"github.com/pmuens/goker/goker"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &goker.StubPlayerStore{}

		cli := goker.NewCLI(in, playerStore)
		cli.PlayPoker()

		goker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from uer input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &goker.StubPlayerStore{}

		cli := goker.NewCLI(in, playerStore)
		cli.PlayPoker()

		goker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
