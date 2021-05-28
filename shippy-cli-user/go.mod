module github.com/dbielecki97/shippy/shippy-user-cli

go 1.16

replace github.com/dbielecki97/shippy/shippy-service-user => ../shippy-service-user

require (
	github.com/asim/go-micro/v3 v3.5.1
	github.com/dbielecki97/shippy/shippy-service-user v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
)
