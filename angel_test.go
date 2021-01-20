package angel

import "testing"

func TestHello(t *testing.T) {
	want := "Hello chris"
	if got := Hello("chris"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
