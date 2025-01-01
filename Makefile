default: protoc frontend backend

frontend:
	$(MAKE) -C frontend -w

backend:
	$(MAKE) -C backend -w

protoc:
	$(MAKE) -C frontend -w protoc
	$(MAKE) -C backend -w go-tools
	$(MAKE) -C proto -w


.PHONY: default frontend backend protoc
