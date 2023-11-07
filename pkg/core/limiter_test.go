package limiter

import (
	"testing"
)

func TestNew_emptyLimitingRules_panics(t *testing.T) {
	expectedErr := "Must specify at least 1 limiting rule."
	defer func() {
		if r := recover(); r != expectedErr {
			t.Errorf("Expected panic with message '%s', but got '%s'", r, expectedErr)
		}
	}()

	New(nil, nil, nil)

	t.Error("Did not panic. Expected to panic because of empty rules.")
}

func TestNew_emptyActions_panics(t *testing.T) {
	limitingRules := []Rule{
		RuleFunc(func(_ Request) bool { return true }),
	}

	expectedErr := "Must specify at least 1 success or failure action."
	defer func() {
		if r := recover(); r != expectedErr {
			t.Errorf("Expected panic with message '%s', but got '%s'", expectedErr, r)
		}
	}()

	New(limitingRules, nil, nil)

	t.Errorf("Did not panic. Expected to panic because of no actions.")
}

func TestNew_returnsLimiter(t *testing.T) {
	limitingRules := []Rule{RuleFunc(func(_ Request) bool { return true })}
	successActions := []Action{ActionFunc(func() error { return nil })}
	failureActions := []Action{ActionFunc(func() error { return nil })}

	limiter := New(limitingRules, successActions, failureActions)

	if limiter == nil {
		t.Errorf("Expected an instance of Limiter, got nil")
	}
}

func TestProcess_validRequest_executesSuccessActions(t *testing.T) {
	limitingRules := []Rule{RuleFunc(func(_ Request) bool { return true })}
	successExecuted := false
	successActions := []Action{
		ActionFunc(func() error {
			successExecuted = true
			return nil
		}),
	}
	limiter := New(limitingRules, successActions, nil)

	limiter.Process(Request{})

	if !successExecuted {
		t.Error("Expected to call success actions.")
	}
}

func TestProcess_validRequest_notCallFailureActions(t *testing.T) {
	limitingRules := []Rule{RuleFunc(func(_ Request) bool { return true })}
	failureActions := []Action{
		ActionFunc(func() error {
			t.Error("Expected not to execute failure actions for successful request.")
			return nil
		}),
	}
	limiter := New(limitingRules, nil, failureActions)

	limiter.Process(Request{})
}

func TestProcess_invalidRequest_callsFailureActions(t *testing.T) {
	limitingRules := []Rule{RuleFunc(func(_ Request) bool { return false })}
	failureExecuted := false
	failureActions := []Action{
		ActionFunc(func() error {
			failureExecuted = true
			return nil
		}),
	}
	limiter := New(limitingRules, nil, failureActions)

	limiter.Process(Request{})

	if !failureExecuted {
		t.Error("Expected to execute failure actions when passed invalid request.")
	}
}

func TestProcess_invalidRequest_notCallsSuccessActions(t *testing.T) {
	limitingRules := []Rule{RuleFunc(func(_ Request) bool { return false })}
	successActions := []Action{
		ActionFunc(func() error {
			t.Error("Expected not to execute success actions when passed invalid request.")
			return nil
		}),
	}
	limiter := New(limitingRules, successActions, nil)

	limiter.Process(Request{})
}
