migrate:
	go run main.go mysql -m

update:
	~/go/bin/swag init
	git add .
	git commit -m "$(commit)"
	git push -u origin master