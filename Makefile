run:
	@ echo "Boot Base Server."; \
	go run main.go; \

run-auth:
	@ echo "Boot Auth Server."; \
	go run authserver_amin.go; \