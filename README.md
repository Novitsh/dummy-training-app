# dummy-training-app

A simple Go + Gin web app scaffold used for AI training sessions.

## Prerequisites

- Go 1.21+

## Run

```bash
go run .
```

Open http://localhost:8080

## Structure

| Path | Purpose |
|---|---|
| `main.go` | Entry point |
| `config/` | App configuration |
| `handlers/` | HTTP handlers (add new pages here) |
| `middleware/` | Gin middleware |
| `templates/` | HTML templates |
| `static/` | CSS and static assets |

## Extend

- Add a new handler in `handlers/`
- Add a route in `main.go`
- Add a template in `templates/`
