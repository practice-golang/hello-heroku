build:
	go build -o bin/

push:
	git push heroku main

commit:
	git add .
	git commit -m "deploy"

deploy: commit push
