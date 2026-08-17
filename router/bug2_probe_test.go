package router

import "testing"

// TestProbe_CatchAllFromRoot verifies that a catch-all segment in the
// first position captures the entire remaining path, including the first
// segment after the leading slash.
func TestProbe_CatchAllFromRoot(t *testing.T) {
	rt := New()
	if err := rt.Register("/*catch", "root-catch"); err != nil {
		t.Fatalf("register /*catch: %v", err)
	}
	m, ok := rt.Match("/a/b/c")
	if !ok {
		t.Fatalf("Match(/a/b/c) no match, want root-catch")
	}
	if m.Label != "root-catch" {
		t.Fatalf("label=%q want root-catch", m.Label)
	}
	if got := m.Params["catch"]; got != "a/b/c" {
		t.Fatalf("catch=%q want a/b/c", got)
	}
}
