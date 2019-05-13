## Shipping package
PROJECT_NAME = example

_cleango:
	@rm -rf proto/*.go

_gengo:
	@protoc -Iproto \
		-I../protoc-gen-auth/options \
		--auth_out="lang=go:proto" \
		proto/*.proto

gen: clean _gengo

clean: _cleango
