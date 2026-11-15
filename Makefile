.PHONY: run
run:
	go run cmd/main.go -u https://httpbin.org/ -r 300 -c 200
docker-build:
	docker build -t stress-test .
docker-run:
	docker run -it --rm stress-test -u https://httpbin.org/ -r 3 -c 2