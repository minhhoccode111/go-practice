# Header + Footer Design

## Direction

Dark terminal vibe using daisyUI `dark` theme with green accent (`#3fb950`).

## Components

### Header (`components/header.templ`)

- daisyUI `navbar` with `bg-base-100 shadow-sm`
- **Left**: site name "try-htmx" in green accent, monospace-style
- **Right**: nav links — Home (`/`), Hello (`/hello`)
- **Mobile**: Alpine.js toggle for hamburger menu (`x-data`, `x-show`, `@click`)
- **HTMX**: links use `hx-get` + `hx-target="#main"` + `hx-push-url="true"` for partial page loads
- Sticky top

### Footer (`components/footer.templ`)

- daisyUI `footer` with `bg-base-200 text-base-content p-10`
- Centered layout: "Built with Go & HTMX" + copyright line
- Separated from content by subtle top border

### Layout update (`layout/app.templ`)

- Add `data-theme="dark"` to `<html>` element
- Replace empty `<header>` and `<footer>` with new components
- Add `<main id="main">` for HTMX target
- Min height `min-h-screen` with flex column so footer sticks to bottom

## Nav links

| Route    | Label |
| -------- | ----- |
| `/`      | Home  |
| `/hello` | Hello |

## Tech usage

- **daisyUI**: `navbar`, `footer`, `btn`, `menu`, `dropdown` classes
- **Tailwind**: `min-h-screen`, `flex`, `flex-col`, `flex-1` for sticky footer
- **Alpine**: `x-data="{ open: false }"`, `@click="open = !open"`, `x-show="open"` for mobile menu
- **HTMX**: `hx-get`, `hx-target="#main"`, `hx-push-url="true"` on nav links
- **Templ**: new component files, composed into layout

## Files to create

- `components/header.templ`
- `components/footer.templ`

## Files to modify

- `layout/app.templ` — wire in new components, add `data-theme`, HTMX target, sticky footer layout
- `main.go` (maybe) — if routes need nav link info passed to header
