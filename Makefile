run:
	@ echo "ベースサーバー起動"; \
	go run main.go; \

run-auth:
	@ echo "認証サーバー起動"; \
	go run authserver_amin.go; \