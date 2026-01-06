package sbrdata

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestCall_CmpCorrectOrdered(t *testing.T) {
	callOne := Call{
		Date: "1767677804",
	}

	callTwo := Call{
		Date: "1767677805",
	}

	calls := Calls{
		Call: []Call{callOne, callTwo},
	}

	slices.SortFunc(calls.Call, Call.Cmp)

	t.Logf("Sorted calls: %v", calls.Call)

	if calls.Call[0].Date != "1767677804" {
		t.Errorf("Expected first call to be 1767677804, got %s", calls.Call[0].Date)
	}
}

func TestCall_CmpCorrectUnordered(t *testing.T) {
	callOne := Call{
		Date: "1767677804",
	}

	callTwo := Call{
		Date: "1767677805",
	}

	calls := Calls{
		Call: []Call{callTwo, callOne},
	}

	slices.SortFunc(calls.Call, Call.Cmp)

	t.Logf("Sorted calls: %v", calls.Call)

	if calls.Call[0].Date != "1767677804" {
		t.Errorf("Expected first call to be 1767677804, got %s", calls.Call[0].Date)
	}
}
