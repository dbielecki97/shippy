module github.com/dbielecki97/shippy/shippy-service-consignment

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/dbielecki97/shippy/shippy-service-vessel => ../shippy-service-vessel

require (
	github.com/asim/go-micro/v3 v3.5.1
	github.com/dbielecki97/shippy/shippy-service-vessel v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.5.2
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
