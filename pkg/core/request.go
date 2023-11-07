package limiter

import "time"

// A request contains metadata used for deciding whether a request should be rate-limited or not.
type Request struct {
	agentId   string
	timestamp time.Time
}
