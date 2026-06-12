proto:
	protoc \
		--proto_path=proto \
		--go_out=proto \
		--go-grpc_out=proto \
		proto/*.proto
	@echo "Proto files generated in the 'proto' directory."

server:
	go run collector/main.go
	
unary:
	go run agent/main.go unary

client:
	go run agent/main.go client

.PHONY: proto server unary client
