.PHONY: call
call:
	curl https://functions.yandexcloud.net/d4edeacbrakvia68do4f

.PHONY: run
run:
	go run .

.PHONY: build
build:
	# GOOS=linux GOARCH=amd64 go build -o dist/georgios
	# zip dist/georgios.zip dist/georgios
	zip dist/georgios.zip main.go go.mod

.PHONY: deploy
deploy: build
	yc serverless function version create \
    	--function-name georgios \
	    --runtime golang121 \
    	--entrypoint main.Handler \
    	--memory 128m \
    	--execution-timeout 5s \
    	--environment VAR1=value1,VAR2=value2 \
    	--source-path dist/georgios.zip
	yc serverless function invoke --name georgios --data '{"name": "Alexander"}'


