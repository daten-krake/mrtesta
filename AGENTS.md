# AGENTS.md

## Project

`mrtesta` — Go-based detection engineering tool. An agent runs as a long-lived process on each client and executes tests that validate detection logic. The agent polls a GitHub repo hourly for new tests plus a config file that controls what runs.

GitHub source of truth for tests + config: `daten-krake/mrtesta`

## Hard constraints

- **Go standard library only.** Do not add external dependencies. `go.mod` must stay dependency-free; `go mod tidy` must produce no `require` block.
- Tests must be runnable in **PowerShell and Python**. The Go agent orchestrates execution of these scripts; it does not reimplement test logic in Go.
- Test selection/scheduling is driven by a **config file hosted on GitHub**, fetched at runtime — not local config. The agent applies that config rather than hardcoding test selection.
- Poll GitHub for new tests **every hour**.
- Keep code simple, readable, low-overhead. Prefer obvious control flow over clever abstractions.
- **Code must be understandable by an amateur coder.** No clever idioms, no deep nesting, no terse one-liners. Prefer verbose-but-obvious.
- **No comments.** Keep code very simple and self-explanatory.

## Module

- Module path: `mrtesta` (no vanity URL — imports are `mrtesta/...`).
- Go 1.26, pinned in `go.mod`.

## Commands

```sh
go build ./...     # build all packages
go test ./...      # run all tests
go vet ./...       # static checks
go mod tidy        # must not introduce dependencies
```

Order when modifying code: `go vet -> go test -> go build`.

## Architecture notes (target design — no implementation yet)

- Single agent binary is the entrypoint; runs persistently on client machines.
- Tests are scripts fetched from GitHub (`.ps1` / `.py`), not compiled into the agent.
- Both PowerShell and Python execution paths must be supported (assume Windows clients for PowerShell).
- Config on GitHub dictates what runs; agent interprets and applies it.