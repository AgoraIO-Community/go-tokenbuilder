package tokenbuilder

import (
	"os"
	"testing"
)

func TestVersion(t *testing.T) {
	const expected = "v1.5.0"
	if Version != expected {
		t.Fatalf("Version = %q, want %q", Version, expected)
	}
}

func TestVersionMatchesReleaseTag(t *testing.T) {
	releaseTag := os.Getenv("RELEASE_TAG")
	if releaseTag == "" {
		t.Skip("RELEASE_TAG is only set for release-tag CI runs")
	}
	if Version != releaseTag {
		t.Fatalf("Version = %q, release tag = %q", Version, releaseTag)
	}
}
