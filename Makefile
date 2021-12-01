build:
	go build -o bin/

push:
	git push heroku main

deploy:
	git add .
	git commit -m "deploy"
	push
