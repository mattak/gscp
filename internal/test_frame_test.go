package internal

import "testing"

type TestFrame struct {
	Test *testing.T
}

func (t TestFrame) AssertEquals(a, b any) {
	if a != b {
		t.Test.Errorf("Failed to equals: %v == %v", a, b)
	}
}

func (t TestFrame) AssertNil(a any) {
	if a != nil {
		t.Test.Errorf("Failed to nil: %v", a)
	}
}

func (t TestFrame) Run(name string, runner func(tf TestFrame)) {
	t.Test.Run(name, func(t *testing.T) {
		tf := TestFrame{Test: t}
		runner(tf)
	})
}
