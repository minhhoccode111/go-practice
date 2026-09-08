# try chi + gorm + templ + htmx + alpine + tailwind + daisyUI

## Installation

install HTMX

```sh
curl -o assets/js/htmx.min.js https://cdn.jsdelivr.net/npm/htmx.org@4.0.0/dist/htmx.min.js
```

install alpineJS

```sh
curl -o assets/js/alpine.min.js https://cdn.jsdelivr.net/npm/alpinejs@3/dist/cdn.min.js
```

install tailwindCSS cli

```sh
curl -sL https://github.com/tailwindlabs/tailwindcss/releases/download/v4.3.3/tailwindcss-linux-x64 -o tailwindcss && chmod +x tailwindcss
```

install daisyUI

```sh
curl -sLO https://github.com/saadeghi/daisyui/releases/latest/download/daisyui.mjs
curl -sLO https://github.com/saadeghi/daisyui/releases/latest/download/daisyui-theme.mjs
mv daisyui*.mjs assets/css/
```

## Getting started

Go server

```sh
air
```

Templ generate

```sh
templ generate --watch
```

Tailwind generate (also for daisyUI)

```sh
./tailwindcss -i assets/css/input.css -o assets/css/output.css --watch
```
