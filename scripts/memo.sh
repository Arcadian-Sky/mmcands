#go run ./cmd/server/. -tb true 
# default tcp :3200
# protoc -I proto/googleapi/ -I proto/. --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import proto/metrics.proto
# protoc -I proto/googleapi/ -I proto/. --grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=import proto/metrics.proto



# protoc -I proto/googleapi/ -I proto/. --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import --grpc-gateway_out=. --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=import --openapiv2_out=. proto/metrics.proto

# protoc -I proto/googleapi/ -I proto/. --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --openapiv2_out=./api/. proto/metrics.proto
# protoc -I proto/googleapi/ -I proto/. --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import proto/metrics.proto
# protoc -I proto/. --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import proto/agent.proto


# mockgen -destination=internal/agent/generated/mocks/mock_agent.pb.go -source=internal/agent/generated/protoagent/agent.pb.go  -package=mocks
# mockgen -destination=internal/agent/generated/mocks/mock_agent_grpc.pb.go -source=internal/agent/generated/protoagent/agent_grpc.pb.go  -package=mocks