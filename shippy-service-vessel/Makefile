protobuf:
	protoc --proto_path=. --go_out=. --micro_out=. \
    		proto/vessel/vessel.proto
build:
	go mod tidy
	go mod vendor
	docker build -t shippy-service-vessel .
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-vessel
