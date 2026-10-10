default: prod

prod:
	npm ci
	npx vite build

dev:
	npx vite

lint:
	npx eslint .
	npx stylelint resources/stylesheets/style.css

.PHONY: default prod dev lint
