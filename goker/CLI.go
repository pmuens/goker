package goker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	in      *bufio.Scanner
	store   PlayerStore
	alerter BlindAlerter
}

func NewCLI(in io.Reader, store PlayerStore, alerter BlindAlerter) *CLI {
	return &CLI{
		in:      bufio.NewScanner(in),
		store:   store,
		alerter: alerter,
	}
}

func (c *CLI) PlayPoker() {
	c.scheduleBlindAlerts()
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

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}
