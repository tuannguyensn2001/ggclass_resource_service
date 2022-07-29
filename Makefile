gen-proto:
	rm -f src/pb/*.go
	protoc --proto_path=proto --go_out=src/pb --go_opt=paths=source_relative \
	--go-grpc_out=src/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=src/pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto