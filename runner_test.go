package runner

import (
	"testing"
)

type mockRunnable struct {
	ran bool
}

func (m *mockRunnable) Run() {
	m.ran = true
}

func TestRun(t *testing.T) {
	r := New()
	m := &mockRunnable{ran: false}
	r.Add(m)

	if m.ran {
		t.Error("Runnable ran prematurely")
	}
	r.Run()
	if !m.ran {
		t.Error("Runnable did not run")
	}
}
