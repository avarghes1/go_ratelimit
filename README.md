Rate Limit
===========

Package can be used to throttle access to resources.
An example would be to limit requests to an external api
which has a rate limit of 10req/sec.

Installation:
---

`go get github.com/avarghes1/go_ratelimit/ratelimit`

Import:
---

`import github.com/avarghes1/go_ratelimit/ratelimit`

Usage
---

```
for i := 0; i < 100; i++ {
    go func() {
        if advance {
            time.Sleep(time.Millisecond * 100)
            go r.Put()
        }
    }()
}
```
