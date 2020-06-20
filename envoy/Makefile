.PHONY: gen-proto gen-web-proto

build-image:
	docker build -t vote-server-envoy:1.0 .

run-image:
	docker run -d -p 9090:9090 --network=host vote-server-envoy:1.0