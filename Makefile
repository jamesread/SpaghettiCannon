default: protoc frontend backend

frontend:
	$(MAKE) -C frontend

backend:
	$(MAKE) -C backend

protoc:
	$(MAKE) -C proto


.PHONY: default frontend backend protoc
