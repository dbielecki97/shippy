module github.com/dbielecki97/shippy/shippy-cli-consignment

go 1.16

replace github.com/dbielecki97/shippy/shippy-service-consignment => ../shippy-service-consignment

replace github.com/dbielecki97/shippy/shippy-service-vessel => ../shippy-service-vessel

require (
	github.com/asim/go-micro/v3 v3.5.1
	github.com/dbielecki97/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
)
