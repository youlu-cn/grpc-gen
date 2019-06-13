## Shipping package
PROJECT_NAME = example

_cleango:
	@rm -rf proto/*.go
	@rm -rf proto/*.md

_gengo:
	@protoc -Iproto \
		-I../protoc-gen-auth \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:proto \
		--grpc-gateway_out=logtostderr=true:proto \
		--auth_out="lang=go:proto" \
		--markdown_out=":proto" \
		proto/*.proto

gen: clean _gengo

clean: _cleango
