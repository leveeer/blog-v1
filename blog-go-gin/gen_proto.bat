@echo off
set BIN_DIR=bin

del /f /s /q go_proto\*.go
for /f %%f in ('dir /b .\proto\*.proto' ) do (
    echo %%~nf
    %BIN_DIR%\protoc.exe -I proto --go_out=plugins=grpc:go_proto proto\%%~nf.proto
    %BIN_DIR%\protoc-go-inject-tag.exe  -input=go_proto/%%~nf.pb.go
)
