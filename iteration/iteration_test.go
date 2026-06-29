package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("q")
	expected := "qqqqq"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("q")
	}
}
