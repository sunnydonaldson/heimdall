// Package limiter implements functions for rate-limiting requests based on rules implementing [rule/Rule].
package limiter

import "errors"

// A Limiter defines the contract to be implemented by [limiter].
//
// Do not implement yourself.
// The only reason this interface exists is to prohibit using [limiter]'s struct literal to bypass constraints.
type Limiter interface {
	Process(req Request) error
}

// A limiter processes requests according to a set of limiting rules.
type limiter struct {
	limitingRules  []Rule
	successActions []Action
	failureActions []Action
}

// New returns an instance of limiter which satisfies the [Limiter] interface.
//
// limitingRules must not be empty, and there must be at least one successAction or failureAction.
func New(limitingRules []Rule, successActions []Action, failureActions []Action) Limiter {
	if len(limitingRules) == 0 {
		panic("Must specify at least 1 limiting rule.")
	} else if len(successActions) == 0 && len(failureActions) == 0 {
		panic("Must specify at least 1 success or failure action.")
	}

	return &limiter{limitingRules, successActions, failureActions}
}

// Process executes successActions if req satisfies all limitingRules, otherwise failureActions.
//
// Returns error if any action fails to execute.
func (l limiter) Process(req Request) error {
	var actions []Action

	err := l.validate(req)
	if err == nil {
		actions = l.successActions
	} else {
		actions = l.failureActions
	}

	return executeAll(actions)
}

func (l limiter) validate(req Request) error {
	for _, rule := range l.limitingRules {
		if !rule.Allows(req) {
			return errors.New("Rule fails")
		}
	}
	return nil
}

func executeAll(actions []Action) error {
	for _, action := range actions {
		err := action.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
