package stat

import (
	"runtime"
	"testing"
)

func TestMemory(t *testing.T) {
	mem, err := Memory()
	if err != nil {
		t.Errorf("expected err to be nil, got %#q", err)
	}

	if mem.Total <= 0 {
		t.Errorf("expected .Total to be > 0, got %d", mem.Total)
	}
	if mem.Free <= 0 {
		t.Errorf("expected .Free to be > 0, got %d", mem.Free)
	}

	if runtime.GOOS != "windows" {
		if mem.Buffers <= 0 {
			t.Errorf("expected .Buffers to be > 0, got %d", mem.Buffers)
		}
		if mem.Cached <= 0 {
			t.Errorf("expected .Cached to be > 0, got %d", mem.Cached)
		}
	}
}
