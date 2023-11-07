package limiter

// A Rule contains a set of conditions to apply to requests.
type Rule interface {
	Allows(Request) bool
}

// A RuleFunc is a functional interface for [Rule].
type RuleFunc func(req Request) bool

// Allows reports whether the request satisfies the rule.
func (rf RuleFunc) Allows(req Request) bool {
	return rf(req)
}
