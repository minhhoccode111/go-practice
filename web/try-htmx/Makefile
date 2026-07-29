.PHONY: watch
watch:
	trap 'kill 0' EXIT; \
	templ generate --watch & \
	./tailwindcss -i assets/css/input.css -o assets/css/output.css --watch & \
	air
