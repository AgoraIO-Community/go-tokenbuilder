package tokenbuilder

import "testing"

func TestVersion(t *testing.T) {
	const expected = "v1.5.0"
	if Version != expected {
		t.Fatalf("Version = %q, want %q", Version, expected)
	}
}
