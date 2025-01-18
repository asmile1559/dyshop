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
	@mv github.com/dyshop/pb/* ./pb
	@rm -rf github.com
