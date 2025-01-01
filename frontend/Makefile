cicd:
	pnpm i -g @bufbuild/protoc-gen-es
	pnpm i -g @connectrpc/protoc-gen-connect-es

webui-codestyle:
	pnpm install
	npx eslint .
	npx stylelint resources/stylesheets/style.css
	parcel build index.pug

webui-dist:
	rm -rf webui webui.dev/dist
	cd webui.dev && npm install
	cd webui.dev && parcel build --public-url "." && mv dist ../webui/

