package goker

import (
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"slices"
	"strings"
	"testing"
	"time"
)

// Stubs.
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// Spys.
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}

type GameSpy struct {
	StartCalled     bool
	StartCalledWith int

	FinishedCalled   bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishCalledWith = winner
}

// Types.
type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

// Helper functions.
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func UserSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func CheckScheduledAlerts(t *testing.T, alerts []ScheduledAlert, blindAlerter *SpyBlindAlerter) {
	for i, want := range alerts {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			AssertScheduledAlert(t, got, want)
		})
	}
}

// Assertions.
func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner, got %q, want %q", store.winCalls[0], winner)
	}
}

func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertLeague(t testing.TB, got, want []Player) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func AssertScheduledAlert(t testing.TB, got ScheduledAlert, want ScheduledAlert) {
	t.Helper()

	gotAmount := got.Amount
	wantAmount := want.Amount
	if gotAmount != wantAmount {
		t.Errorf("got amount %d, want %d", gotAmount, wantAmount)
	}

	gotAt := got.At
	wantAt := want.At
	if gotAt != wantAt {
		t.Errorf("got scheduled time of %v, want %v", gotAt, wantAt)
	}
}

func AssertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func AssertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()

	if game.StartCalledWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d, but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func AssertGameNotFinished(t testing.TB, game *GameSpy) {
	t.Helper()

	if game.FinishedCalled {
		t.Errorf("game should not have finished")
	}
}

func AssertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func AssertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()

	if game.FinishCalledWith != winner {
		t.Errorf("expected finish called with %q, but got %q", winner, game.FinishCalledWith)
	}
}
