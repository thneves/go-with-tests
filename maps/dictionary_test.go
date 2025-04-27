package maps

import "testing"

func TestSearch(t *testing.T) {
	Dictionary := Dictionary{"key": "this is the value"}
	t.Run("key exists", func(t *testing.T) {

		got, _ := Dictionary.Search("key")
		want := "this is the value"

		assertStrings(t, got, want)
	})

	t.Run("key doesn't exist", func(t *testing.T) {
		_, got := Dictionary.Search("noKeyHere")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
