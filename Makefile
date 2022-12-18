prune:
	docker image prune -f
gen-protos:
	protoc --go_out=. --grpc-gateway_out=. --go-grpc_out=. protos/**/*.proto 
export_go_path:
	export GO_PATH=~/go && export PATH=$PATH:/$GO_PATH/bin