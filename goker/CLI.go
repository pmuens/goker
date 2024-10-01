package goker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	in      *bufio.Scanner
	out     io.Writer
	store   PlayerStore
	alerter BlindAlerter
}

func NewCLI(in io.Reader, out io.Writer, store PlayerStore, alerter BlindAlerter) *CLI {
	return &CLI{
		in:      bufio.NewScanner(in),
		out:     out,
		store:   store,
		alerter: alerter,
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberOfPlayers, _ := strconv.Atoi(c.readLine())

	c.scheduleBlindAlerts(numberOfPlayers)
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

func (c *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}
