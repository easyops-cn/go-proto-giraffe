regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogo

	protoc \
	--proto_path=../../../ \
	--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:../../../ \
	github.com/easyops-cn/go-proto-giraffe/giraffe.proto \
	github.com/easyops-cn/go-proto-giraffe/http.proto \

	protoc \
	--proto_path=../../../ \
	--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:../../../ \
	github.com/easyops-cn/go-proto-giraffe/plugin/extension.proto \
