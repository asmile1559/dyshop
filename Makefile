.PHONY: gen-backend-proto
# before run gen-backend-proto,
# run "rm -rf pb/backend" to clean pb/backend folders
# after run gen-backend-proto,
# if no go.mod file, run "cd pb && go mod init github.com/dyshop/pb && go mod tidy"
# if there is go.mod, run "cd pb && go mod tidy"
# the last step is run "go work use pb"
gen-backend-proto:
	@protoc --go_out=. --proto_path=proto/backend/ proto/backend/*.proto
	@protoc --go-grpc_out=. --proto_path=proto/backend/ proto/backend/*.proto
	@cp github.com/asmile1559/dyshop/pb/* ./pb -rf
	@rm -rf github.com

.PHONY: gen-frontend-proto
gen-frontend-proto:
	@protoc --go_out=. --experimental_allow_proto3_optional --proto_path=proto/frontend/ proto/frontend/*.proto
	@protoc --go-grpc_out=. --experimental_allow_proto3_optional --proto_path=proto/frontend/ proto/frontend/*.proto
	@cp github.com/asmile1559/dyshop/pb/* ./pb -rf
	@rm -rf github.com