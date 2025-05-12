default: protoc frontend service

frontend:
	$(MAKE) -C frontend -w

service:
	$(MAKE) -C service -w

protoc:
	$(MAKE) -wC proto

docs-local:
	$(MAKE) -wC docs
	./docs/node_modules/.bin/antora antora-playbook.yml --to-dir /var/www/html/sc-docs/

make docs:
	$(MAKE) -wC docs
	./docs/node_modules/.bin/antora antora-playbook.yml

.PHONY: default frontend service protoc docs
