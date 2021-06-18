build:
	go mod tidy
	go mod vendor
	docker build -t shippy-service-consignment .
protobuf:
	protoc --proto_path=. --go_out=. --micro_out=. \
 		proto/consignment/consignment.proto
run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment
