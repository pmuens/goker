package goker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	in    *bufio.Scanner
	store PlayerStore
}

func NewCLI(in io.Reader, store PlayerStore) *CLI {
	return &CLI{
		in:    bufio.NewScanner(in),
		store: store,
	}
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.store.RecordWin(extractWinner(userInput))
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
