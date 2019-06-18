package poker

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	/* t.Run("/league from reader", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		//read again, test the seeker part
		got = store.GetLeague()
		assertLeague(t, got, want)
	}) */

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")

		want := 33

		AssertScoreEquals(t, got, want)

	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayStore(database)
		AssertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")

		want := 34

		AssertScoreEquals(t, got, want)

	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayStore(database)
		AssertNoError(t, err)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")

		want := 1

		AssertScoreEquals(t, got, want)

	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayStore(database)
		AssertNoError(t, err)

	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)
		//read again, test the seeker part
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}
