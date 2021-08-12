include .env
export

setup:
	@if [ ! -d "serverless" ]; then \
		echo "installing serverless ..."; \
		docker-compose run --rm serverless_go serverless install --url https://github.com/serverless/serverless; \
	fi
	
build: setup
	cd api; env GOOS=linux go build -tags aws_lambda -ldflags="-s -w" -o ../bin/api main.go
	cd message; env GOOS=linux go build -tags aws_lambda -ldflags="-s -w" -o ../bin/message main.go

deploy:
	docker-compose build serverless_go
	docker-compose run --rm serverless_go serverless deploy -v --stage=${STAGE}

cleanup: 
	docker-compose build serverless_go
	docker-compose run --rm serverless_go serverless remove -v --stage=${STAGE}

run-local:
	@cd api; go run main.go