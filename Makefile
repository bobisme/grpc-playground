GATEWAY_IMAGE := my-grpc-gateway

gen-go:
	docker run --rm -it -v $(PWD):/defs namely/protoc-all:1.15_0 \
		-f thing.proto -l go --with-gateway --with-docs

gen-gateway: thing.proto
	docker run -v $(PWD):/defs namely/gen-grpc-gateway \
		-f thing.proto -s EchoService

gateway-image: gen-gateway
	docker build -t $(GATEWAY_IMAGE) gen/grpc-gateway/

run-grpc:
	go run main.go -echo_endpoint localhost:50051

run-gateway:
	docker run --rm -it --name my-gateway \
		-p 8088:80 \
		$(GATEWAY_IMAGE) --backend=host.docker.internal:50051

run-graphql:
	node ./gql.js
