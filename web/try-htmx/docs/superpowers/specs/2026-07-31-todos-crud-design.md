# Todos Full CRUD Design

**Date:** 2026-07-31
**Goal:** Finish the todos feature: add, toggle, edit, delete — with good UX (loading indicators, validation, error handling). In-memory state, no DB.

**Tech Stack:** Go, chi, templ, HTMX 2.0, Alpine.js 3, daisyUI 5, Tailwind v4

---

## Current State & Bugs

- Todos stored in `todos []entity.Todo`; IDs come from the `/counter` demo counter — shared variable corrupts both features. Fix: separate `nextID` counter.
- Toggle button uses `hx-swap="none"` — server re-renders but UI never updates.
- Form errors don't display: form element carries `hx-swap-oob="true"` — htmx rips the whole form out of band and fights the `beforeend` target.
- `TodosLayout` has invalid HTML: `<li id="todos-loading">` directly inside `<div>`.
- Server validation message contradicts its own check (`len < 1 || len > 255` vs "greater than 1 and less than 255").

## State

- Keep `todos []entity.Todo` + `nextID uint` (monotonic), unsynchronized — single-user learning app.
- `ponytail:` comment: swap slice for gorm DB later.

## Routes

| Route | Purpose | Response |
|---|---|---|
| `GET /todos` | page (exists) | layout |
| `GET /htmx/todos` | full list (exists) | `Todos()` |
| `POST /htmx/todos` | create (exists) | new row + cleared errors, or errors only |
| `PATCH /htmx/{id}/toggle` | toggle (exists, fix) | re-rendered row |
| `GET /htmx/{id}/edit` | edit form (new) | `TodoEditForm()` |
| `PUT /htmx/{id}` | save edit (new) | re-rendered row |
| `GET /htmx/{id}` | single row (new) — edit cancel | re-rendered row |
| `DELETE /htmx/{id}` | delete (new) | full list |

Invalid/missing ID: 404 empty body — htmx ignores non-2xx, no swap.

## Components

**`Todo(todo)`** — `<li>` row: title + toggle/edit/delete buttons.
- toggle: `hx-patch="/htmx/{id}/toggle"` `hx-target="closest li"` `hx-swap="outerHTML"`
- edit: `hx-get="/htmx/{id}/edit"` same target/swap
- delete: `hx-delete="/htmx/{id}"` `hx-target="#todos-list"` `hx-swap="outerHTML"` — replaces whole list (deleting a `<li>` can't be replaced by list content)
- all buttons: `hx-disabled-elt="this"`, spinner `hx-indicator`, `aria-label`

**`TodoEditForm(todo)`** — `<li>` with prefilled input + Ok/Cancel:
- Ok: `hx-put="/htmx/{id}"` `hx-target="closest li"` `hx-swap="outerHTML"`
- Cancel: `hx-get="/htmx/{id}"` same target/swap
- input `maxlength="255"` `autofocus`

**`TodoForm`** (fix) — form + sibling `<div id="form-errors">`:
- remove `hx-swap-oob` from form
- `hx-target="#todos-list"` `hx-swap="beforeend"`
- reset: `hx-on::after-request="if(event.detail.successful) this.reset()"`
- generic failure: `hx-on::response-error` → "Something went wrong" into `#form-errors`
- submit button: `hx-disabled-elt="this"` + spinner

**`Todos(todos)`** — `<li>` loop; empty slice → "No todos yet" empty state.

**`TodosLayout`** (fix) — `<ul id="todos-list">` wrapper; loading indicator as sibling `<div id="todos-loading" class="htmx-indicator">loading…</div>` with `hx-indicator="#todos-loading"`.

## Server Logic

- POST create: trim title; if `len < 1 || len > 255` → render only `<div id="form-errors">` OOB with error "Title must be 1–255 characters". No list mutation. Else append, render new `<li>` + empty errors div OOB.
- PUT edit: same validation; invalid → re-render `TodoEditForm` + errors div OOB (200). NOT OOB-only: htmx `swapOuterHTML` with empty fragment removes the target — the edit `<li>` would vanish (verified in htmx 2.0.10 source).
- Toggle: flip `IsDone`, render single row.
- Delete: filter slice, render full `Todos()`.

## Error Handling

- Validation errors: `<div id="form-errors">` swapped OOB from anywhere (create and edit).
- Stale-error clearing: toggle, edit-save, and cancel responses include `FormErrors(nil)` — otherwise a previous validation error lingers until the next create submit.
- Edit save failure: 200 + OOB errors only — htmx skips ALL swaps (including OOB) on non-2xx, so no redirect/error status needed.
- Network/5xx: `hx-on::response-error` on form → generic message.
- Missing ID: 404, silent no-op.

## UX Checklist

- Spinner visible during every request (buttons + initial list load)
- Input disabled while submitting
- Errors appear in `#form-errors`, cleared on next successful submit
- Empty state message
- Autofocus on edit input
- Confirm dialog on delete

## Verification

No test framework in project — manual browser checklist:

1. `templ generate` && `go build ./...`
2. Run server, check:
   - add valid → row appends, input clears, errors empty
   - add empty / 256-char → errors show, no append
   - toggle → strike-through toggles (row re-renders)
   - edit → row becomes form, autofocus; Ok saves; Cancel reverts
   - delete → confirm, row disappears
   - loading spinners show during 1s sleep on each action
   - initial list load shows "loading..." then rows
   - empty list → "No todos yet"
   - reload page → todos persist in session, toggle still works
   - `/counter` unaffected by todo IDs
3. No commit until user says.
