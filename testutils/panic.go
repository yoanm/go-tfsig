package testutils

import "testing"

func ExpectPanic(t *testing.T, tcname string, fn func(), expectedError string) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Case \"%s\": expects panic but everything went weel", tcname)
		} else if err != expectedError {
			t.Errorf("Case \"%s\": expects panic '%s', got '%s'", tcname, expectedError, err)
		}
	}()

	fn()
}
