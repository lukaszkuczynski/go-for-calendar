localrun:
	GOOGLE_CREDENTIALS_PATH="./localcredentials.json" go run main.go

deploy:
	gcloud app deploy