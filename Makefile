run:
	@ echo "Boot main server. click http://localhost:8081/"
	go run main.go

docker-build:
	docker build -t go-malproxy .

docker-run:
	docker run -d -p 8080:8081 go-malproxy 

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
