.PHONY: run
run:
	go run cmd/main.go -u http://localhost:8080 -r 300 -c 200 -H "API_KEY: token_de_acesso"
docker-build:
	docker build -t stress-test .
docker-run:
	docker run -it --rm stress-test -u http://localhost:8080/ -r 3 -c 2 -H "API_KEY": token_de_acesso"