#!/bin/bash

DIR=.
SERVER_OUT_DIR=server/echo
CLIENT_OUT_DIR=client/echo

rm $SERVER_OUT_DIR/*pb*
rm $CLIENT_OUT_DIR/*pb*

protoc -I=$DIR echo.proto \
    --go_out=$SERVER_OUT_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$SERVER_OUT_DIR --go-grpc_opt=paths=source_relative

protoc -I=$DIR echo.proto \
    --js_out=import_style=commonjs:$CLIENT_OUT_DIR \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$CLIENT_OUT_DIR