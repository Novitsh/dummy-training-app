# Training App Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build a simple Gin-based Go web app with Home and About pages as a training scaffold for developers.

**Architecture:** MVC-inspired layout with `handlers/`, `middleware/`, `templates/`, and `static/` directories. Gin renders HTML templates using a shared layout. No database.

**Tech Stack:** Go 1.21+, Gin (`github.com/gin-gonic/gin`), HTML templates, plain CSS.

---

### Task 1: Initialise the Go module and install Gin

**Files:**
- Create: `go.mod`

**Step 1: Initialise the module**

```bash
cd /Users/kevinthiels/training/dummy-training-app
go mod init github.com/Novitsh/dummy-training-app
```

Expected: `go.mod` created.

**Step 2: Add Gin**

```bash
go get github.com/gin-gonic/gin
```

Expected: `go.mod` and `go.sum` updated with Gin dependency.

**Step 3: Commit**

```bash
git init
git remote add origin git@github.com:Novitsh/dummy-training-app.git
git add go.mod go.sum
git commit -m "chore: init module and add gin dependency"
```

---

### Task 2: Config package

**Files:**
- Create: `config/config.go`

**Step 1: Write the config struct and loader**

```go
// config/config.go
package config

import "os"

type Config struct {
    Port string
}

func Load() Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    return Config{Port: port}
}
```

**Step 2: Commit**

```bash
git add config/config.go
git commit -m "feat: add config package with PORT env support"
```

---

### Task 3: Request logger middleware

**Files:**
- Create: `middleware/logger.go`

**Step 1: Write the middleware**

```go
// middleware/logger.go
package middleware

import (
    "github.com/gin-gonic/gin"
)

// Logger returns a named request-logging middleware using Gin's built-in logger.
func Logger() gin.HandlerFunc {
    return gin.Logger()
}
```

**Step 2: Commit**

```bash
git add middleware/logger.go
git commit -m "feat: add logger middleware"
```

---

### Task 4: HTML templates

**Files:**
- Create: `templates/layout.html`
- Create: `templates/home.html`
- Create: `templates/about.html`

**Step 1: Write the shared layout**

```html
<!-- templates/layout.html -->
{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Training App</title>
  <link rel="stylesheet" href="/static/style.css" />
</head>
<body>
  <nav>
    <a href="/">Home</a>
    <a href="/about">About</a>
  </nav>
  <main>
    {{ template "content" . }}
  </main>
  <footer>
    <p>Training App &mdash; extend me!</p>
  </footer>
</body>
</html>
{{ end }}
```

**Step 2: Write the home template**

```html
<!-- templates/home.html -->
{{ define "content" }}
<h1>Welcome</h1>
<p>This is the training app. Start extending it!</p>
{{ end }}
```

**Step 3: Write the about template**

```html
<!-- templates/about.html -->
{{ define "content" }}
<h1>About</h1>
<p>A simple Go + Gin web app scaffold for training sessions.</p>
{{ end }}
```

**Step 4: Commit**

```bash
git add templates/
git commit -m "feat: add layout and page templates"
```

---

### Task 5: Static CSS

**Files:**
- Create: `static/style.css`

**Step 1: Write minimal CSS**

```css
/* static/style.css */
*, *::before, *::after { box-sizing: border-box; }

body {
  font-family: system-ui, sans-serif;
  margin: 0;
  padding: 0;
  color: #222;
  background: #f9f9f9;
}

nav {
  background: #1a1a2e;
  padding: 1rem 2rem;
  display: flex;
  gap: 1.5rem;
}

nav a {
  color: #eee;
  text-decoration: none;
  font-weight: 500;
}

nav a:hover { text-decoration: underline; }

main {
  max-width: 800px;
  margin: 2rem auto;
  padding: 0 1.5rem;
}

footer {
  text-align: center;
  padding: 2rem;
  color: #888;
  font-size: 0.85rem;
}
```

**Step 2: Commit**

```bash
git add static/style.css
git commit -m "feat: add basic CSS styles"
```

---

### Task 6: Handlers

**Files:**
- Create: `handlers/home.go`
- Create: `handlers/about.go`

**Step 1: Write the home handler**

```go
// handlers/home.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
    c.HTML(http.StatusOK, "home.html", gin.H{})
}
```

**Step 2: Write the about handler**

```go
// handlers/about.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func About(c *gin.Context) {
    c.HTML(http.StatusOK, "about.html", gin.H{})
}
```

**Step 3: Commit**

```bash
git add handlers/
git commit -m "feat: add home and about handlers"
```

---

### Task 7: main.go — wire everything together

**Files:**
- Create: `main.go`

**Step 1: Write main.go**

```go
// main.go
package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/Novitsh/dummy-training-app/config"
    "github.com/Novitsh/dummy-training-app/handlers"
    "github.com/Novitsh/dummy-training-app/middleware"
)

func main() {
    cfg := config.Load()

    r := gin.New()
    r.Use(middleware.Logger())
    r.Use(gin.Recovery())

    r.LoadHTMLGlob("templates/*.html")
    r.Static("/static", "./static")

    r.GET("/", handlers.Home)
    r.GET("/about", handlers.About)

    addr := fmt.Sprintf(":%s", cfg.Port)
    r.Run(addr)
}
```

**Step 2: Verify the app compiles and runs**

```bash
go run .
```

Expected: Server starts on port 8080. No errors.

**Step 3: Smoke test in browser or curl**

```bash
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/
# Expected: 200
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/about
# Expected: 200
```

**Step 4: Commit**

```bash
git add main.go
git commit -m "feat: wire router, middleware, and handlers in main.go"
```

---

### Task 8: README and final push

**Files:**
- Create: `README.md`

**Step 1: Write the README**

```markdown
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
```

**Step 2: Commit README and design docs**

```bash
git add README.md docs/
git commit -m "docs: add README and design documents"
```

**Step 3: Push to GitHub**

```bash
git push -u origin main
```

Expected: All commits pushed to `git@github.com:Novitsh/dummy-training-app.git`.

---
