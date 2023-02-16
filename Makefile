proto-gen:
	protoc -I ./proto/ \
	--go-grpc_out=./internal/controller/grpc \
	--go_out=./pkg/auth  ./proto/*.proto --proto_path=.

