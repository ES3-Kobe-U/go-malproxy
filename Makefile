run:
	@ echo "サーバを起動します"
	go run main.go

all:
	git add .
	git commit -m "commit all changed files"
	git push origin HEAD

MS=""
git:
	git commit -m ${MS}
	git push origin HEAD