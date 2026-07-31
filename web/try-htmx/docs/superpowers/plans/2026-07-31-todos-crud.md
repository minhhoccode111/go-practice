# Todos Full CRUD Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

> **NOTE:** User instruction — do NOT `git commit` anything. Skip all commit steps until user approves.

**Goal:** Finish todos feature: add, toggle, edit, delete, with loading indicators, validation, and error handling. In-memory state.

**Architecture:** Row-level inline editing. Each `<li>` row swaps itself (`closest li` + `outerHTML`) between display and edit form. List container is `<ul id="todos-list">`; create appends rows (`beforeend`), delete rewrites list (`innerHTML`). Validation errors always ride a 200 response as an out-of-band swap of `<div id="form-errors">` — htmx skips ALL swaps on non-2xx.

**Tech Stack:** Go, chi, templ, HTMX 2.0, daisyUI 5 (uses `list-row`, `loading` spinner, `input`, `btn` classes)

**Spec:** `docs/superpowers/specs/2026-07-31-todos-crud-design.md`

---

### Task 1: Rewrite `components/todo.templ` — row with 3 buttons

**Files:**
- Modify: `components/todo.templ`

- [ ] **Step 1: Replace file content**

```templ
package components

import "try-htmx/entity"
import "fmt"

templ Todo(todo entity.Todo) {
	<li class={ "list-row flex gap-2 items-center justify-between", templ.KV("line-through", todo.IsDone) }>
		<span class="flex-1">{ todo.Title }</span>
		<div class="flex gap-1">
			<button
				hx-patch={ fmt.Sprintf("/htmx/%d/toggle", todo.ID) }
				hx-target="closest li"
				hx-swap="outerHTML"
				aria-label={ fmt.Sprintf("Toggle %s", todo.Title) }
				class="btn btn-xs btn-primary"
			>
				<span class="htmx-indicator loading loading-spinner loading-xs"></span>
				&#10003;
			</button>
			<button
				hx-get={ fmt.Sprintf("/htmx/%d/edit", todo.ID) }
				hx-target="closest li"
				hx-swap="outerHTML"
				aria-label={ fmt.Sprintf("Edit %s", todo.Title) }
				class="btn btn-xs"
			>
				<span class="htmx-indicator loading loading-spinner loading-xs"></span>
				&#9998;
			</button>
			<button
				hx-delete={ fmt.Sprintf("/htmx/%d", todo.ID) }
				hx-target="#todos-list"
				hx-swap="innerHTML"
				hx-confirm="Delete this todo?"
				aria-label={ fmt.Sprintf("Delete %s", todo.Title) }
				class="btn btn-xs btn-error"
			>
				<span class="htmx-indicator loading loading-spinner loading-xs"></span>
				&#10005;
			</button>
		</div>
	</li>
}
```

Notes:
- `hx-disabled-elt="this"` intentionally omitted on row buttons — per-button indicator already prevents confusion; keep markup lean. (Reconsider only if double-click bugs appear.)
- `templ.KV("line-through", todo.IsDone)` — conditional class, exists already.

- [ ] **Step 2: Regenerate and verify it compiles**

Run: `templ generate && go build ./...`
Expected: no errors

---

### Task 2: Create `components/todo_edit.templ` — inline edit form

**Files:**
- Create: `components/todo_edit.templ`

- [ ] **Step 1: Write file**

```templ
package components

import "try-htmx/entity"
import "fmt"

templ TodoEditForm(todo entity.Todo) {
	<li class="list-row">
		<form
			hx-put={ fmt.Sprintf("/htmx/%d", todo.ID) }
			hx-target="closest li"
			hx-swap="outerHTML"
			hx-indicator="this"
			hx-disabled-elt="button"
			class="flex flex-1 gap-2"
		>
			<input
				name="title"
				value={ todo.Title }
				type="text"
				class="input w-full"
				required
				minlength="1"
				maxlength="255"
				autofocus
			/>
			<button class="btn btn-xs btn-primary">
				<span class="htmx-indicator loading loading-spinner loading-xs"></span>
				ok
			</button>
			<button
				type="button"
				hx-get={ fmt.Sprintf("/htmx/%d", todo.ID) }
				hx-target="closest li"
				hx-swap="outerHTML"
				class="btn btn-xs"
			>
				cancel
			</button>
		</form>
	</li>
}
```

Notes:
- Cancel button `type="button"` so it never native-submits; `hx-get` replaces row back to display.
- Validation failure response contains only the OOB `#form-errors` div — the `<li>` is untouched, edit form keeps user input.

- [ ] **Step 2: Regenerate and verify it compiles**

Run: `templ generate && go build ./...`
Expected: no errors

---

### Task 3: Rewrite `components/todo_form.templ` — fix OOB bug, add `FormErrors`

**Files:**
- Modify: `components/todo_form.templ`

- [ ] **Step 1: Replace file content**

```templ
package components

templ TodoForm(errors []string) {
	<form
		id="todo-form"
		hx-post="/htmx/todos"
		hx-target="#todos-list"
		hx-swap="beforeend"
		hx-indicator="this"
		hx-disabled-elt="find button"
		hx-on::after-request="if(event.detail.successful && !document.getElementById('form-errors').textContent.trim()) this.reset()"
		hx-on::response-error="document.getElementById('form-errors').textContent = 'Something went wrong'"
		hx-on::send-error="document.getElementById('form-errors').textContent = 'Something went wrong'"
		class="flex items-center justify-between gap-2"
	>
		<input
			id="title"
			name="title"
			type="text"
			placeholder="Add todo"
			class="input w-full"
			required
			minlength="1"
			maxlength="255"
			aria-label="New todo title"
		/>
		<button class="btn btn-primary">
			<span class="htmx-indicator loading loading-spinner loading-xs"></span>
			submit
		</button>
	</form>
	@FormErrors(errors)
}

templ FormErrors(errors []string) {
	<div
		id="form-errors"
		hx-swap-oob="true"
		class="text-xs font-bold text-red-500"
		role="alert"
	>
		@Errors(errors)
	</div>
}
```

Notes:
- **Bug fix:** form no longer carries `hx-swap-oob="true"` — it stays in layout, only `#form-errors` is OOB.
- `hx-on::after-request` — HTMX 2.0 event syntax; resets input only on success. The `#form-errors` emptiness check: validation errors ride a 200 (OOB swap lands before `after-request`), so a bare `successful` check would wipe the user's invalid input. OOB swap replaces `#form-errors` content first — error text present → no reset; empty list → reset.
- `hx-on::response-error` — fires on non-2xx responses (e.g. server 500). `hx-on::send-error` — fires on network failure. Both show the same generic message.
- The `hx-swap-oob="true"` attribute in initial page HTML is inert — htmx only processes OOB during swaps.

- [ ] **Step 2: Regenerate and verify it compiles**

Run: `templ generate && go build ./...`
Expected: no errors

---

### Task 4: Add empty state to `components/todos.templ`

**Files:**
- Modify: `components/todos.templ`

- [ ] **Step 1: Replace file content**

```templ
package components

import "try-htmx/entity"

templ Todos(todos []entity.Todo) {
	if len(todos) == 0 {
		<li class="list-row opacity-50">No todos yet &mdash; add one above</li>
	} else {
		for _, v := range todos {
			@Todo(v)
		}
	}
}
```

- [ ] **Step 2: Regenerate and verify it compiles**

Run: `templ generate && go build ./...`
Expected: no errors

---

### Task 5: Fix `layout/todos_layout.templ` — proper list container + indicator

**Files:**
- Modify: `layout/todos_layout.templ`

- [ ] **Step 1: Replace file content**

```templ
package layout

import "try-htmx/components"

templ TodosLayout() {
	<div class="mx-auto max-w-md space-y-2 p-4">
		<h2 class="text-lg font-bold">Todos</h2>
		@components.TodoForm(nil)
		<div id="todos-loading" class="htmx-indicator text-sm opacity-50">loading...</div>
		<ul
			id="todos-list"
			class="list"
			hx-get="/htmx/todos"
			hx-trigger="load"
			hx-swap="innerHTML"
			hx-indicator="#todos-loading"
		></ul>
	</div>
}
```

Notes:
- **Bug fix:** was `<div>` with a stray `<li>` loading item inside — invalid HTML. Now `<ul id="todos-list">` + sibling indicator div.
- `hx-swap="innerHTML"` — full list replaces contents; delete route also targets `#todos-list` with `innerHTML`.

- [ ] **Step 2: Regenerate and verify it compiles**

Run: `templ generate && go build ./...`
Expected: no errors

---

### Task 6: Rewrite handlers + state in `main.go`

**Files:**
- Modify: `main.go:113-176` (todos section)

- [ ] **Step 1: Replace the whole todos block (lines 113-176)**

Delete `counter`-based seed/ID hack. New block:

```go
	// todos -------------------------------------------------------------------

	var todos []entity.Todo
	var nextID uint = 1

	// ponytail: in-memory slice + manual IDs; swap for gorm DB when
	// persistence is needed. Not goroutine-safe — single-user learning app.

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		content := layout.TodosLayout()
		if r.Header.Get("HX-Request") == "true" {
			renderer.Render(r.Context(), w, templ.Join(layout.PageTitle("Todos"), content))
			return
		}

		renderer.Render(r.Context(), w, layout.App(layout.AppData{PageTitle: "Todos"}, content))
	})

	r.Get("/htmx/todos", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		renderer.Render(r.Context(), w, components.Todos(todos))
	})

	r.Post("/htmx/todos", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := strings.TrimSpace(r.FormValue("title"))
		if len(title) < 1 || len(title) > 255 {
			renderer.Render(r.Context(), w, components.FormErrors(
				[]string{"Title must be 1-255 characters"},
			))
			return
		}
		todo := entity.Todo{
			Model:  gorm.Model{ID: nextID},
			Title:  title,
			IsDone: false,
		}
		nextID++
		todos = append(todos, todo)
		renderer.Render(r.Context(), w, templ.Join(components.Todo(todo), components.FormErrors(nil)))
	})

	r.Get("/htmx/{todoID}", func(w http.ResponseWriter, r *http.Request) {
		todo, ok := getTodo(todos, r)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		renderer.Render(r.Context(), w, components.Todo(*todo))
	})

	r.Get("/htmx/{todoID}/edit", func(w http.ResponseWriter, r *http.Request) {
		todo, ok := getTodo(todos, r)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		renderer.Render(r.Context(), w, components.TodoEditForm(*todo))
	})

	r.Put("/htmx/{todoID}", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		todo, ok := getTodo(todos, r)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		title := strings.TrimSpace(r.FormValue("title"))
		if len(title) < 1 || len(title) > 255 {
			renderer.Render(r.Context(), w, templ.Join(
				components.TodoEditForm(*todo),
				components.FormErrors([]string{"Title must be 1-255 characters"}),
			))
			return
		}
		todo.Title = title
		renderer.Render(r.Context(), w, components.Todo(*todo))
	})

	r.Patch("/htmx/{todoID}/toggle", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		todo, ok := getTodo(todos, r)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		todo.IsDone = !todo.IsDone
		renderer.Render(r.Context(), w, components.Todo(*todo))
	})

	r.Delete("/htmx/{todoID}", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		id, ok := parseTodoID(r)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		for i := range todos {
			if todos[i].ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				renderer.Render(r.Context(), w, components.Todos(todos))
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	})
```

- [ ] **Step 2: Add helper functions before `main()`**

Insert after `loadConfig()` (before `func main()`):

```go
func parseTodoID(r *http.Request) (uint, bool) {
	u, err := strconv.ParseUint(chi.URLParam(r, "todoID"), 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(u), true
}

func getTodo(todos []entity.Todo, r *http.Request) (*entity.Todo, bool) {
	id, ok := parseTodoID(r)
	if !ok {
		return nil, false
	}
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], true
		}
	}
	return nil, false
}
```

Notes:
- `getTodo` returns pointer into the slice — mutation (`todo.Title = ...`, `todo.IsDone = ...`) persists. Fine within one request; `append` may reallocate but no stale pointers are kept across requests.
- Chi exact-match wins over `{todoID}` — `/htmx/todos` GET still hits the list handler.
- Non-2xx responses: htmx skips swaps — 404 = silent no-op, no error UI needed.
- **PUT validation error MUST re-render `TodoEditForm`** (not OOB-only): verified in htmx 2.0.10 source (`swapOuterHTML`, htmx.js:1697-1725) — after OOB extraction the remaining fragment is empty, and outerHTML swap removes the target entirely. The edit `<li>` would vanish. Returning the form + errors keeps the row in edit mode.

- [ ] **Step 3: Verify it compiles**

Run: `templ generate && go build ./...`
Expected: no errors

---

### Task 7: Manual browser verification

**Files:** none

- [ ] **Step 1: Rebuild CSS**

Run: `./tailwindcss -i assets/css/input.css -o assets/css/output.css`
Expected: no errors

- [ ] **Step 2: Run server**

Run: `go run .`
Expected: listens on port from `.env` (default 8080)

- [ ] **Step 3: Browser checklist** — open `http://localhost:8080/todos`

| # | Action | Expected |
|---|--------|----------|
| 1 | Initial load | "loading..." shows ~1s, then list (empty state "No todos yet") |
| 2 | Add "buy milk" | spinner on submit, row appends, input clears, no errors |
| 3 | Add empty title | red "Title must be 1-255 characters" in `#form-errors`, no row |
| 4 | Add 256-char title | same error |
| 5 | Toggle ✓ | spinner, strike-through toggles |
| 6 | Edit ✎ | row becomes form, input focused with title; edit title, Enter or ok → row re-rendered |
| 7 | Edit → invalid (empty) | error shows, row stays edit form (input reverts to stored title — server re-render, not echo) |
| 8 | Edit → cancel | row reverts, no change |
| 9 | Delete ✕ | confirm dialog, row disappears; delete all → empty state |
| 10 | Reload page | todos persist in session, toggle/edit/delete still work |
| 11 | Visit `/counter` | counter unaffected by todo operations |

- [ ] **Step 4: Devtools console check**

Open DevTools → Console: no errors. Network tab: all requests 200 (except 404 never triggered).

- [ ] **Step 5: Stop server, report results**

Report to user: checklist pass/fail, anything unexpected. Do not commit — user decides.
