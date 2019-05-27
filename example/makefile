## Shipping package
PROJECT_NAME = example

_cleango:
	@rm -rf proto/*.go

_gengo:
	@protoc -Iproto \
		-I../protoc-gen-auth \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:proto \
		--grpc-gateway_out=logtostderr=true:proto \
		--auth_out="lang=go:proto" \
		proto/*.proto

gen: clean _gengo

clean: _cleango
