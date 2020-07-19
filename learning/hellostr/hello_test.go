package hellostr

import (
	"testing"
)

func TestHello(t *testing.T) {

	// Writing function inside a function
	assertString := func(t *testing.T, i int, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' wanted for test num %d, wanted it to be '%s'", got, i, want)
		}
	}

	t.Run("Test with a name and Language with structs", func(t *testing.T) {
		name := "Champu"

		type Test struct {
			in  string
			out string
		}

		var tests = []Test{
			{"English", "Hello, " + name},
			{"Spanish", "Hola, " + name},
			{"French", "Bonjour, " + name},
		}

		for i, test := range tests {
			got := Hello(name, test.in)
			want := test.out
			assertString(t, i, got, want)
		}
	})

	t.Run("Test without a name and language", func(t *testing.T) {
		i := -1 // Just filling this up for our test case with a false value.
		name := ""
		language := ""
		got := Hello(name, language)
		want := "Hello, World"
		assertString(t, i, got, want)
	})

}
