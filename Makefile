build:
	go build -o ./bin

deploy:
	git add .
	git commit -m "deploy"
	git push heroku master:main