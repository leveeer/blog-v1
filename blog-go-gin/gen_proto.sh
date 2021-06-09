#!/bin/bash

rm ./go_proto/*.go
rm ./cpp/proto/*.h
rm ./cpp/proto/*.cc
rm ./script/grpc/*.lua
rm ./pb/*.pb

BIN_PATH="bin"
./$BIN_PATH/protoc -I ./build/protobuf/include -I ./proto --go_out=plugins=grpc:./go_proto ./proto/*.proto
./$BIN_PATH/protoc -I ./build/protobuf/include -I ./proto --cpp_out=./cpp/proto ./proto/*.proto
./$BIN_PATH/protoc -I ./build/protobuf/include -I ./proto --grpc_out=./cpp/proto --plugin=protoc-gen-grpc=./$BIN_PATH/grpc_cpp_plugin ./proto/*.proto
./$BIN_PATH/protoc -I ./build/protobuf/include -I ./proto --gotemplate_out=template_dir=templates,single-package-mode=true:. --plugin=protoc-gen-gotemplate=./$BIN_PATH/protoc-gen-gotemplate ./proto/*.proto
for file in ./proto/*.proto
do
    if test -f $file
    then
./$BIN_PATH/protoc -I ./build/protobuf/include -I ./proto -o./pb/$(basename $file .proto).pb $file
    fi
done
