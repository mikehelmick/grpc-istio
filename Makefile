
protoc:
	@protoc --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pkg/counter/pb/*.proto
	@goimports -w pkg/counter/pb
.PHONY: protoc
