run:
	@echo "\033[35m Boot main server. Click the following URL. \033[m" 
	@echo "\033[31m localhost     => http://localhost:8081/ \033[m"
	@echo "\033[32m Digital Ocean => http://159.89.34.164:8081/ \033[m"
	go run main.go

docker-build:
	docker build -t go-malproxy .

docker-run:
	docker run -d -p 8080:8081 go-malproxy 

deploy:
	make docker-build
	make docker-run
	@echo "\033[35m Boot main server. Click the following URL. \033[m" 
	@echo "\033[33m Docker        => http://159.89.34.164:8080/ \033[m"

all:
	git add .
	git commit -m "commit all changed files"
	git push origin HEAD

MS=""
git:
	git commit -m ${MS}
	git push origin HEAD

# linux用
clean:
	find server/templates -name 'autogen*' -delete
	find server/templates/img -name '*.png' -delete

# windows用
del:
	del D:\go-malproxy\server\templates\autogen*.html
	del D:\go-malproxy\server\templates\img\*.png
