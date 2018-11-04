build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/teamtools/teampay-greet-service \
		proto/greetpb/greetpb.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t teampay-greet-service .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		teampay-greet-service