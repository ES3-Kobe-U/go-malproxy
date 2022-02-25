run:
	@ echo "Boot main server. click http://localhost:3333/"
	go run main.go

all:
	git add .
	git commit -m "commit all changed files"
	git push origin HEAD

MS=""
git:
	git commit -m ${MS}
	git push origin HEAD

clean:
	find server/templates -name 'autogen*' -delete
	find server/service -name 'autogen*' -delete
	find server/service -name '*.png' -delete