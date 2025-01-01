default: protoc frontend backend

frontend:
	$(MAKE) -C frontend

backend:
	$(MAKE) -C backend

protoc:
	$(MAKE) -c backend go-tools
	$(MAKE) -C proto


.PHONY: default frontend backend protoc
