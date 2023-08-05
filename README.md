# Heimdall
A configurable rate-limiter.

## Rate-limiter Components
To understand the motivation for Heimdall, you first need to know what a rate limiter does

1. Recieve request
1. Pass request through rate-limiting algorithm
1. Carry out different sets of operations depending on whether the request was rate-limited or not
For example, if the request was rate limited, we might respond with HTTP error 429.

## The Vision
Heimdall aims to work out of the box for most common uses. but it'll be fully configurable with different:
* Ways to recieve requests (HTTP, gRPC, tRPC, etc..)
* Rate-limiting algorithms (token bucket, leaking token bucket, sliding log window, etc...)
* Success actions (Forward HTTP request, log result, etc...)
* Failure actions (Respond with 429, log error, send request to message queue for best-effort processing, etc...)

The first milestone is setting up a framework for the different components to interact through.
Then it'll be easy enough for users to contribute their own rate-limiting algorithms, and actions.


