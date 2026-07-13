# TODO_longterm.md

Findings from the OpenCode reviewer (PR #1) — address as the feature work lands.

## Major

- [ ] Add `config/config.json` stub to the repo. `main.go` hardcodes `configPath = "config/config.json"` but the file doesn't exist yet, so every poll 404s.
- [ ] Add Go tests. `config.Parse` and `github.FetchRaw` are testable in isolation. Start with table-driven tests for `config.Parse` (valid JSON, invalid JSON, empty JSON). For `FetchRaw`, inject an `http.RoundTripper` so it can be tested without network.

## Minor

- [ ] `IntervalSeconds` is defined in `internal/config/config.go` but `main.go` hardcodes `pollInterval = 1 * time.Hour` and never reads it. Either wire it up or remove the field until needed.
- [ ] `fetchAndApply` returns `*config.Config` but the polling loop discards it. The next PR that adds test execution needs a place to hold the current config (package-level var or a struct wrapping the ticker + config).

## Suggestions

- [ ] Graceful shutdown: add `SIGTERM`/`SIGINT` handling that stops the ticker and drains. Not blocking for scaffolding.
- [ ] `User-Agent` header: `internal/github/github.go` sends none. GitHub's API ToS require one. Add `req.Header.Set("User-Agent", "mrtesta-agent/0.1")`.