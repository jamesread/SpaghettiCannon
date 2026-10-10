default: protoc frontend service

frontend:
	$(MAKE) -C frontend -w

service:
	$(MAKE) -C service -w

protoc:
	$(MAKE) -wC proto

migrate:
	$(MAKE) -C database/mysql -w

test:
	$(MAKE) -C service -w unittests

lint:
	$(MAKE) -C service -w codestyle
	$(MAKE) -C frontend -w lint

run:
	$(MAKE) -C service -w run

docs-local:
	$(MAKE) -wC docs
	./docs/node_modules/.bin/antora antora-playbook.yml --to-dir /var/www/html/sc-docs/

docs:
	$(MAKE) -wC docs
	./docs/node_modules/.bin/antora antora-playbook.yml

.PHONY: default frontend service protoc migrate test lint run docs docs-local
