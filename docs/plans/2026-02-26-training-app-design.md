# Training App Design

**Date:** 2026-02-26
**Purpose:** A simple Go web app scaffold for AI training sessions. Developers extend it during training.

## Goal

Provide a clean, well-structured starting point that demonstrates real-world Go + Gin patterns without being overwhelming. The structure should guide developers on where to add new code.

## Architecture

Gin-based HTTP server with HTML template rendering. MVC-inspired directory layout.

```
dummy-training-app/
├── main.go                  # Entry point, router setup
├── config/
│   └── config.go            # App config (port, env)
├── handlers/
│   ├── home.go              # Home page handler
│   └── about.go             # About page handler
├── middleware/
│   └── logger.go            # Request logger middleware
├── templates/
│   ├── layout.html          # Base layout (nav, header, footer)
│   ├── home.html            # Home page content
│   └── about.html           # About page content
├── static/
│   └── style.css            # Basic CSS
├── go.mod
└── go.sum
```

## Components

- **main.go** — Loads config, registers middleware, mounts routes, starts server.
- **config/config.go** — Reads `PORT` env var with a default fallback.
- **handlers/** — One file per page. Each handler renders its template with the shared layout.
- **middleware/logger.go** — Wraps Gin's built-in logger as a named middleware example.
- **templates/** — `layout.html` defines the shared nav/footer; page templates are embedded in it.
- **static/style.css** — Minimal CSS, enough for a readable UI.

## Pages

| Route    | Handler       | Template     |
|----------|---------------|--------------|
| `/`      | handlers.Home | home.html    |
| `/about` | handlers.About| about.html   |

## Data Flow

Request → Gin router → Middleware (logger) → Handler → Template render → Response

## No Database

Intentionally omitted. Adding a datastore (SQLite, Postgres, in-memory) is a natural training exercise.

## Success Criteria

- `go run .` starts the server with no errors
- Home and About pages render correctly
- Code is easy to follow and extend
- Pushes cleanly to git@github.com:Novitsh/dummy-training-app.git
