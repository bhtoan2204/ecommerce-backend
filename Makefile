proto-user:
	@cd ./app/user && \
		go install golang.org/x/tools/cmd/goimports@v0.1.12 && \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
		protoc --experimental_allow_proto3_optional --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. ./proto/*/*.proto && \
		goimports -w proto
.PHONY: proto-user

