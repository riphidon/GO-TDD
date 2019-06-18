package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/riphidon/firstGoApp"
)

type GameSpy struct {
	StartedWith    int
	FinishedWith   string
	StartCalled    bool
	FinishedCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers

}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishedWith = winner
}

func UserSends(messages ...string) *strings.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got '%s' wanted '%s'", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

	})

	t.Run("Starts with 3 players and finish with Chris as winner", func(t *testing.T) {

		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		in := UserSends("3", "Chris wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		AssertGameStartedWith(t, game, 3)
		AssertGameFinishedWith(t, game, "Chris ")
		AssertMessageSentToUser(t, stdout, poker.PlayerPrompt)
	})

	t.Run("Starts with 8 players and finish with Cleo as winner", func(t *testing.T) {

		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		in := UserSends("8", "Cleo wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		AssertGameStartedWith(t, game, 8)
		AssertGameFinishedWith(t, game, "Cleo ")
		AssertMessageSentToUser(t, stdout, poker.PlayerPrompt)
	})

	/* t.Run("error if non numeric value is entered and game does not start", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := UserSends("pie")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertGameNotStarted(t, game)
		AssertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("Prints an error if Player enters 'Lloyd's a killer instead of 'Ruth wins", func(t *testing.T) {

		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		in := UserSends("3", "Lloyd wins")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		AssertMessageSentToUser(t, stdout, poker.BadWinnerInputMsg)
	}) */

}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

func AssertMessageSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got '%s' sent to stdout but expected %+v", got, messages)
	}

}

func AssertGameStartedWith(t *testing.T, game *GameSpy, numberOfPlayers int) {
	t.Helper()
	if game.StartedWith != numberOfPlayers {
		t.Errorf("wanted start called with %d but got %d", numberOfPlayers, game.StartedWith)
	}
}

func AssertGameFinishedWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("wanted finish called with '%s' as the winner but got '%s'", winner, game.FinishedWith)
	}
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game should not have finished")
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}
