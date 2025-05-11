default: protoc frontend service

frontend:
	$(MAKE) -C frontend -w

service:
	$(MAKE) -C service -w

protoc:
	$(MAKE) -wC proto


.PHONY: default frontend service protoc
