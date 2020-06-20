.PHONY: gen-proto gen-web-proto

gen-go-proto:
	protoc -I proto \
		vote.proto \
	    --go_out=plugins=grpc:server/gen/pb

gen-ts-proto:
	protoc -I proto \
		vote.proto \
		--js_out=import_style=commonjs,binary:web/gen/pb \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:web/gen/pb