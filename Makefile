.PHONY: run
run:
	go run cmd/main.go -u https://httpbin.org/ -r 3 -c 2
docker-build:
	docker build -t stress-test .
docker-run:
	docker run -it --rm stress-test -u https://httpbin.org/ -r 3 -c 2