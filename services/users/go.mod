module github.com/sssoultrix/event-go/services/users

go 1.25.3

require (
	github.com/sssoultrix/event-go/contracts/users v0.0.0
	golang.org/x/crypto v0.46.0
	google.golang.org/grpc v1.78.0
	google.golang.org/protobuf v1.36.11
)

require (
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
)

replace github.com/sssoultrix/event-go/contracts/users => ../../contracts/users
